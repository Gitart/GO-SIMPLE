## TCP Proxy с логированием


TCP прокси-сервер c логированием всех запросов и ответы в виде hex-дампа.

### Что мы должны получить в итоге:

прокси-сервер слушает TCP порт и отправляет все запросы на определённый адрес   
получает ответ и перенаправляет его клиенту   
логирует все запросы и ответы в файл либо stdout   
Концептуально прокси-сервер реализовать очень просто. 

### Алгоритм работы:

1. принимаем подключение от клиента
2. соединяемся с удалённым сервером
3. копируем всё что шлёт клиент на сервер
4. копируем всё что шлёт сервер на клиент

Становится интереснее когда нам требуется стать между сервером и клиентом и 
как-то получить и обработать все данные, которыми они общаются. Для этого воспользуемся полиморфизмом, 
который в go реализовывается через интерфейсы. 
Мы знаем что io.Copy принимает два аргумента: io.Writer и io.Reader и мы вполне можем сделать свой io.Writer, 
который запишет и в файл и в коннекшн.

Для начала определим структуру нашего дампера. 
Она должна содержать в себе изначальный io.Writer, дополнительный io.Writer для дампа 
(в нашем случае это будет файл) и строковый label для индикации что именно мы дампим, 
чтоб потом можно было различать откуда пришли данные и куда идут.

```golang
// dumper.go
// ...
type dumper struct {
  // куда писать
  w      io.Writer
  // что писать
  label  string
  // куда дампить
  dumpTo io.Writer
}
```

//
Всё что нам осталось это реализовать метод Write для соответствия интерфейсу io.Writer:

```golang
// dumper.go
// ...
func (d *dumper) Write(b []byte) (int, error) {
  // формируем служебное сообщение со временем и лейблом
  message := fmt.Sprintf("[%s] %s\n", time.Now().Format(time.RFC3339), d.label)
  // пишем сообщение в дамп
  io.WriteString(d.dumpTo, message)
  // добавляем hex-дамп
  io.WriteString(d.dumpTo, hex.Dump(b))
  // и наконец пишем в базовый Writer
  return d.w.Write(b)
}
```

Дампер готов, переходим к реализации самого сервера:

```golang
// server.go
type proxyServer struct {
  // файлик для дампа
  dumpTo     *os.File
  // по какому адресу слушать
  localAddr  string
  // и адрес удалённого сервера
  remoteAddr string
}
```

Далее применим стандартный подход для TCP сервера. Слушаем порт, когда по нему кто-то стучится - отправляем обработку клиента в отдельную горутину и дальше слушаем порт.


```golang
func (ps *proxyServer) start() error {
  listener, err := net.Listen("tcp", ps.localAddr)
  if err != nil {
    return err
  }
  for {
    conn, err := listener.Accept()
    if err != nil {
      log.Println(err)
    }
    go ps.handleClient(conn)
  }
}
```

И основная логика прокси-сервера – обработка соединения. Нам нужно дозвониться до удалённого сервера и настроить двухстороннюю связь сервер-клиент, при этом обернуть её в наш dumper.
И на случай если удалённый сервер захочет сразу что-то прислать в свежесозданное соединение (так делают, например, некоторые ftp-сервера – присылают баннер), мы не дожидаемся реквестов от клиента, а сразу слушаем сервер.

```golang
// server.go

func (ps *proxyServer) handleClient(conn net.Conn) {
  // дозвон до удалённого сервера
  remoteConn, err := net.Dial("tcp", ps.remoteAddr)
  if err != nil {
    log.Println(err)
    return
  }
  // оборачиваем remoteConn в дампер запросов
  req := &dumper{w: remoteConn, label: requestFrom(conn), dumpTo: ps.dumpTo}
  // оборачиваем conn в дампер ответов
  resp := &dumper{w: conn, label: responseTo(conn), dumpTo: ps.dumpTo}
  // перенаправляем ответы сервера клиенту в отдельной горутине
  go io.Copy(resp, remoteConn)
  // перенаправляем запросы клиента на удалённый сервер
  _, err = io.Copy(req, conn)
  if err != nil {
    log.Println(err)
    return
  }
  // на всякий случай скидываем всё что есть в файл, чтоб не потерялось
  err = ps.dumpTo.Sync()
  if err != nil {
    log.Println(err)
    return
  }
}

// красиво выводим информацию по запросу
func requestFrom(conn net.Conn) string {
  return fmt.Sprintf("[==>>] Request from %s", conn.RemoteAddr().String())
}

// красиво выводим информацию по ответу
func responseTo(conn net.Conn) string {
  return fmt.Sprintf("[<<==] Response to %s", conn.RemoteAddr().String())
}
// ...
```

Единственное что хотелось бы ещё добавить – это healthcheck при инициализации прокси, чтоб сразу можно было понять, готов ли удалённый сервер принимать соединения по указанному адресу. Для этого дозваниваемся по указанному адресу и сразу просто закрываем соединение, если никаких ошибок не возникло – значит сервер скорее жив чем мёртв.

```golang
// server.go
// ...
func (ps *proxyServer) healthcheck() error {
  conn, err := net.Dial("tcp", ps.remoteAddr)
  if err != nil {
    return err
  }
  return conn.Close()
}
// ...
```


Ну и осталось реализовать функцию main которая разберётся с флагами, инициализирует сервер и запустит его.

```golang
// main.go
// ...

var (
  remoteHost      = flag.String("host", "", "Remote host")
  remotePort      = flag.Int("port", 0, "Remote port")
  listen          = flag.String("listen", ":4242", "Local address to listen")
  dump            = flag.String("dump", "", "Write dump to file")
  // дадим пользователю возможность пропустить healthcheck
  skipHealthcheck = flag.Bool("skip-healthcheck", false, "Skip healthcheck")
)

func main() {
  flag.Parse()
  remoteAddr := fmt.Sprintf("%s:%d", *remoteHost, *remotePort)
  proxy := &proxyServer{localAddr: *listen, remoteAddr: remoteAddr, dumpTo: dumpTo(*dump)}
  var err error
  if !*skipHealthcheck {
    err = proxy.healthcheck()
    if err != nil {
      log.Fatal(err)
    }
    log.Printf("Healthcheck to %s OK", remoteAddr)
  }
  err = proxy.start()
  if err != nil {
    log.Fatal(err)
  }
}
```

И последний штрих, функция dumpTo, которая определяет куда дампить – в файл или в stdout:

```golang
// main.go
// ...
func dumpTo(filename string) *os.File {
  // если имя файла не передано, либо не удаётся создать файл, дампим в stdout
  dumpTo := os.Stdout
  if len(filename) > 0 {
    file, err := os.Create(filename)
    if err != nil {
      log.Printf("Fail to open file %s, fallback to stdout", filename)
    } else {
      dumpTo = file
    }
  }
  return dumpTo
}
```

И наконец попробуем, что же у нас получилось:

$ go build
$ ./goproxy --host yanzay.com --port 80 --listen ":4242"

$ # в отдельном терминале делаем тестовый запрос
$ curl localhost:4242
И любуемся красивым дампом запроса и ответа:
