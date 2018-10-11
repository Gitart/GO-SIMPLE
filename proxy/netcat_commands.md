## Удалённое выполнение программ в go netcat

В предыдущей статье мы рассмотрели реализацию простейшего аналога известной утилиты netcat на go, 
сегодня мы расширим её возможности, а именно добавим возможность выполнять произвольную команду на 
сервере и отдавать управление ей на клиент. Сценарий работы будет приблизительно такой:


Сервер запускается со специальным флагом (-c), который указывает на то что он будет выполнять команды
Клиент запускается со флагом -e в котором указывается команда для выполнения (например -e sh выполнит sh)
Клиент получает контроль над удалённой программой, то есть его stdin передаётся в stdin программы, 
а stdout и stderr программы возвращается клиенту

Итак, приступим. Для начала добавим нужные флаги для сервера и клиента:

```golang
var (
  // ...
  command = flag.Bool("c", false, "Command server")
  execute = flag.String("e", "", "Execute command")
)
```

Функции main и startServer остаются без изменений, код из предыдущей статьи:

```golang
func main() {
  flag.Parse()
  if *listen {
    startServer()
    return
  }
  if len(flag.Args()) < 2 {
    fmt.Println("Hostname and port required")
    return
  }
  serverHost := flag.Arg(0)
  serverPort := flag.Arg(1)
  startClient(fmt.Sprintf("%s:%s", serverHost, serverPort))
}

func startServer() {
  addr := fmt.Sprintf("%s:%d", *host, *port)
  listener, err := net.Listen("tcp", addr)

  if err != nil {
    panic(err)
  }

  log.Printf("Listening for connections on %s", listener.Addr().String())

  for {
    conn, err := listener.Accept()
    if err != nil {
      log.Printf("Error accepting connection from client: %s", err)
    } else {
      go processClient(conn)
    }
  }
}
```

## Сервер
Обработаем в функции processClient случай, когда сервер запущен в командном режиме:

```golang
func processClient(conn net.Conn) {
  if *command {
    // вся дальшейшая обработка соединения проходит в функции launchCommand
    err := launchCommand(conn)
    if err != nil {
      log.Println(err)
      conn.Close()
      return
    }
  }
  _, err := io.Copy(os.Stdout, conn)
  if err != nil {
    log.Println(err)
  }
  conn.Close()
}
```

Будем считать что клиент присылает первой строкой команду, которую сервер должен выполнить, а дальше уже идёт взаимодействие с выполняемой программой:

```golang
func launchCommand(conn net.Conn) error {
  // создаём bufio.Reader чтоб прочитать первую строку
  reader := bufio.NewReader(conn)
  line, err := reader.ReadString('\n')
  if err != nil {
    return err
  }
  fmt.Printf("Command: %s\n", line)
  // создаём комманду для выполнения
  cmd := exec.Command(strings.TrimSpace(line))
  // захватываем stdin
  stdin, err := cmd.StdinPipe()
  if err != nil {
    return err
  }
  // захватываем stdout
  stdout, err := cmd.StdoutPipe()
  if err != nil {
    return err
  }
  // захватываем stderr, клиент тоже хочет видеть ошибки
  stderr, err := cmd.StderrPipe()
  if err != nil {
    return err
  }
  // всё что приходит от клиента, перенаправляем в stdin программы
  go io.Copy(stdin, conn)
  // всё что программа выдаёт в stdout и в stderr, возвращаем клиенту
  go io.Copy(conn, stdout)
  go io.Copy(conn, stderr)
  // и запускаем
  return cmd.Run()
}
```

На данном этапе сервер готов, перейдём к клиентской части.

## Клиент

Как мы уже договорились выше, клиент должен послать первой строкой команду, а дальше продолжать посылать свой stdin и получать от сервера stdout, посмотрим на функцию startClient:

```golang
func startClient(addr string) {
  // как обычно дозваниваемся на сервер
  conn, err := net.Dial("tcp", addr)
  if err != nil {
    fmt.Printf("Can't connect to server: %s\n", err)
    return
  }
  // если передан параметр -e то отправляем команду в первой строке
  if len(*execute) > 0 {
    cmd := fmt.Sprintf("%s\n", *execute)
    conn.Write([]byte(cmd))
  }
  // перенаправляем всё что отвечает сервер к нам в stdout
  go io.Copy(os.Stdout, conn)
  // и знакомый уже нам код отправки всего что пишем на сервер
  _, err = io.Copy(conn, os.Stdin)
  if err != nil {
    fmt.Printf("Connection error: %s\n", err)
  }
}
```

Вот и всё что требуется от клиента, давайте проверим теперь как это всё работает вместе:

```
$ go build

$ # запускаем сервер в командном режиме
$ ./netcat -l -p 1408 -c

$ # в отдельном терминале клиент, который запускает sh
$ ./netcat -e sh localhost 1408
И попробуем выполнить несколько команд:
```


Как видим, всё работает как мы и ожидали. Теперь у нас есть полноценный удалённый shell.

Полный исходный код программы можно почитать здесь: https://github.com/yanzay/netcat/tree/v0.1

Источник
https://yanzay.com/post/go_netcat_commands/
