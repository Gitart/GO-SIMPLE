# Передача файлов по TCP, дорабатываем netcat на go

Это третья статья из цикла “netcat на go”, контекст можно получить прочитав первые две

В данной статье мы доработаем наш упрощённый netcat и добавим в него режим передачи файлов с клиента на сервер. Так как количество таких режимов и количество кода растёт, разобьём наше приложение на 3 файла:

main.go для парсинга флагов и запуска соответственно сервера или клиента   
server.go вся серверная логика   
client.go вся клиентская логика   

Начнём с нашей входной точки, main.go. 
Здесь ничего нового, немного исправленная функция main из предыдущих статей. 
Просто парсит флаги командной строки и просто запускает сервер или клиент:

```golang
package main

import (
  "flag"
  "fmt"
  "log"
  "os"
)

var (
  // Server
  listen     = flag.Bool("l", false, "Listen")
  host       = flag.String("h", "localhost", "Host")
  port       = flag.Int("p", 0, "Port")
  command    = flag.Bool("c", false, "Command server")
  fileServer = flag.Bool("f", false, "Server for file upload")
  // Client
  execute = flag.String("e", "", "Execute command")
  upload  = flag.String("u", "", "Upload file")
)

func main() {
  flag.Parse()
  if *listen {
    addr := fmt.Sprintf("%s:%d", *host, *port)
    // запускаем сервер
    err := startServer(addr, *command, *fileServer)
    if err != nil {
      log.Fatal(err)
    }
  } else {
    if len(flag.Args()) < 2 {
      fmt.Println("Hostname and port required")
      os.Exit(1)
    }
    serverHost := flag.Arg(0)
    serverPort := flag.Arg(1)
    addr := fmt.Sprintf("%s:%s", serverHost, serverPort)
    // запускаем клиент
    err := startClient(addr, *execute, *upload)
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }
  }
}
```

По традиции сначала реализуем сервер.

## Сервер

Итак, задача: получить от клиента файл и сохранить его в локальной директории. Для того чтобы сохранить файл нам, логично, понадобится его имя и его содержимое (хотя имя мы конечно можем сгенерировать любое, но боюсь клиент такого поведения не оценит).
Мы уже встречали подобную задачу когда реализовывали выполнение команд на удалённом сервере. И предыдущее решение тоже отлично подходит для текущей задачи – первой строкой мы можем получить от клиента имя файла, а всё остальное сохранить как контент. Приступим:
Уже привычный нам запуск сервера и распределение обрабоки соединений по горутинам: 

```golang
func startServer(addr string, command bool, fileServer bool) error {
  listener, err := net.Listen("tcp", addr)
  if err != nil {
    return err
  }

  log.Printf("Listening for connections on %s", listener.Addr().String())

  for {
    conn, err := listener.Accept()
    if err != nil {
      log.Printf("Error accepting connection from client: %s", err)
    } else {
      go processClient(conn, command, fileServer)
    }
  }
}
```

Дальше собственно функция обработки клиентского соединения:

```golang
// Обратите внимание на тип входного параметра conn - интерфейсы наше всё
func processClient(conn io.ReadWriteCloser, command bool, fileServer bool) error {
  if command && fileServer {
    return fmt.Errorf("Can't launch server in command and file mode simultaneously")
  }
  var err error
  switch {
  case command:
    // сервер будет ждать команды от клиента
    err = commandProcessor(conn)
  case fileServer:
    // ждём файл
    err = fileProcessor(conn)
  default:
    // просто выводим всё что присылает клиент на stdout
    err = defaultProcessor(conn)
  }
  return err
}
```

Как вы уже поняли, основную работу на сервере выполняют функции commandProcessor, fileProcessor и defaultProcessor. Полный исходный код можно посмотреть на гитхабе, а вот нам сейчас интересно именно получение и сохранение файлика:

```golang
// чем глубже в код, тем меньше интерфейсы
func fileProcessor(conn io.ReadCloser) error {
  defer conn.Close()
  reader := bufio.NewReader(conn)
  // читаем первую строку - это будет название файла
  line, err := reader.ReadString('\n')
  if err != nil {
    return err
  }
  line = strings.TrimSpace(line)
  // создаём файл с заданным именем в текущей директории
  file, err := os.Create(line)
  if err != nil {
    return err
  }
  // и копируем в него всё что дальше приходит от клиента
  _, err = io.Copy(file, conn)
  return err
}
```

## Клиент
Запуск клиента аналогичен запуску сервера, есть несколько режимов и за каждый режим отвечает отдельная функция:

```golang
func startClient(addr string, execute string, upload string) error {
  conn, err := net.Dial("tcp", addr)
  if err != nil {
    return fmt.Errorf("Can't connect to server: %s\n", err)
  }
  if len(execute) > 0 && len(upload) > 0 {
    return fmt.Errorf("Can't execute command and upload file simultaneously")
  }
  switch {
  case len(execute) > 0:
    err = commandClient(execute, conn)
  case len(upload) > 0:
    err = fileClient(upload, conn)
  default:
    err = defaultClient(conn)
  }
  return err
}
```

Самое интересное впереди, посмотрим на fileClient:

```golang
func fileClient(filename string, conn io.WriteCloser) error {
  // os.Open открывает файл на чтение
  file, err := os.Open(filename)
  if err != nil {
    return err
  }
  // для того чтобы получить настоящее имя файла,
  // возьмём у него Stat()
  stat, err := file.Stat()
  if err != nil {
    return err
  }
  // а в этой функции и происходит аплоад файла
  return uploadFile(stat.Name(), conn, file)
}
```

Итак, у нас есть имя файла, есть файл из которого можно читать контент и есть коннекшн в который его можно писать, что может быть проще?

```golang
func uploadFile(name string, conn io.WriteCloser, file io.ReadCloser) error {
  // сообщаем серверу имя, чтоб знал
  _, err := io.WriteString(conn, fmt.Sprintf("%s\n", name))
  if err != nil {
    return err
  }
  // и копируем весь контент файла прямиком в сеть
  _, err = io.Copy(conn, file)
  if err != nil {
    return err
  }
  // и не забываем всё за собой позакрывать
  err = conn.Close()
  if err != nil {
    return err
  }
  return file.Close()
}
```

Вот и всё, полностью работоспособный мини-netcat готов! 
Полный код здесь: https://github.com/yanzay/netcat/tree/v0.1.1


Источник : https://yanzay.com/post/go_netcat_files/
