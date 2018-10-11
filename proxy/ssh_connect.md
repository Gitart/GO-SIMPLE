# Безопасное TCP соединение поверх SSH на go

В прошлой статье мы рассмотрели как сделать простой SSH клиент на go и выполнить команду на удалённом сервере. 
В этот раз мы воспользуемся возможностями протокола SSH для шифрования TCP трафика. 
Для этого нам понадобится написать SSH сервер и наладить между ним и клиентом двустороннюю связь.

В итоге мы оформим всё как go пакет, чтобы можно было им пользоваться в своих дальнейших экспериментах.

### Для того чтоб удобно и прозрачно было всем этим пользоваться, я выделю пару требований:

1. защищенное соединение должно реализовывать интерфейс net.Conn
2. сервер должен реализовывать интерфейс net.Listener
3. Таким образом мы сможем поверх этого шифрованного TCP соединения пускать любой трафик уровня приложения, 
будь то HTTP, передача файлов или простых текстовых сообщений. Погнали!

Как и для реализации клиента, для сервера воспользуемся пакетом golang.org/x/crypto/ssh. 
Для запуска SSH сервера необходимо для начала запустить обычный TCP Listener, 
а потом при подключении клиента, завернуть коннект в ssh.NewServerConn.

Начнём реализации сервера с функции Listen, которая должна вернуть net.Listener:

```golang
// тут нужен адрес где слушать и ключик для сервера
func Listen(addr string, privateKeyPath string) (net.Listener, error) {
  // запускаем обычный TCP Listener
  listener, err := net.Listen("tcp", addr)
  if err != nil {
    return nil, err
  }
  // тут подготовим ssh.ServerConfig со всеми настройками
  config, err := serverConfig(privateKeyPath)
  if err != nil {
    return nil, err
  }
  // вернём объект server, который реализует net.Listener
  return &server{listener: listener, config: config}, nil
}
```

Ну и раз начали, давайте сразу реализуем net.Listener на структуре server. 
У него всего три метода: Accept(), Close() и Addr(). Последние два мы просто делегируем нижестоящему TCP соединению, 
а вот Accept() должен вернуть защищенное соединение:

```golang
func (s *server) Accept() (net.Conn, error) {
  // принимаем соединение
  conn, err := s.listener.Accept()
  if err != nil {
    return nil, err
  }
  // создаём на коннекте SSH канал (не путать с go channel)
  sshChannel, err := s.channelFromConn(conn)
  if err != nil {
    return nil, err
  }
  // и конструируем защищенное соединение
  return &secureConnection{conn: conn, channel: sshChannel}, nil
}

// делегируем TCP Listener
func (s *server) Close() error {
  return s.listener.Close()
}

// делегируем TCP Listener
func (s *server) Addr() net.Addr {
  return s.listener.Addr()
}
```

И наконец рассмотрим как же создаётся защищенный канал:

```golang
func (s *server) channelFromConn(conn net.Conn) (ssh.Channel, error) {
  // тут много всего, а нам нужен только один SSH канал
  _, ch, _, err := ssh.NewServerConn(conn, s.config)
  if err != nil {
    return nil, err
  }
  // вычитываем SSH канал из go канала :)
  c := <-ch
  // и принимаем его
  sshChannel, _, err := c.Accept()
  if err != nil {
    return nil, err
  }
  return sshChannel, nil
}
```

Теперь что касается secureConnection. Как мы уже видели выше, оно состоит из SSH канала и низлежащего TCP соединения.

```golang
type secureConnection struct {
  conn    net.Conn
  channel ssh.Channel
}
```

Это определённо должен быть net.Conn, а в этом интерфейсе аж 8 методов. Но самые интересные для нас это вот эти два:

```golang
func (sc *secureConnection) Read(b []byte) (n int, err error) {
  return sc.channel.Read(b)
}

func (sc *secureConnection) Write(b []byte) (n int, err error) {
  return sc.channel.Write(b)
}

func (sc *secureConnection) Close() error {
  err := sc.channel.Close()
  if err != nil {
    return err
  }
  return sc.conn.Close()
}
```

Вот и всё, тут мы видим достаточную реализацию io.ReadWriteCloser, с которой уже можно работать как хочешь. Остальные методы смело делегируем низлежащему TCP соединению.
Более подробно изучить код можно вот здесь: https://github.com/yanzay/seccon. А мы перейдём к самому интересному - к примеру использования.

В статье Реализация аналога netcat на go мы уже рассматривали классический подход к реализации TCP сервера и клиента, применим тот же подход и здесь.

Тут всё должно быть уже знакомо, клиент:

```golang
package main

import (
  "io"
  "os"

  "github.com/yanzay/seccon"
)

func main() {
  // создаём ssh клиент для пользователя yanzay
  client := seccon.NewClient("yanzay")
  // алло, сервер?
  conn, err := client.Dial("localhost:2022")
  if err != nil {
    panic(err)
  }
  // перенаправляем весь stdin в защищенное соединение
  _, err = io.Copy(conn, os.Stdin)
  if err != nil {
    panic(err)
  }
}
```

### И сервер:

```golang
package main

import (
  "io"
  "log"
  "os"

  "github.com/yanzay/seccon"
)

func main() {
  // запускаем сервер
  listener, err := seccon.Listen(":2022", "")
  if err != nil {
    log.Fatal(err)
  }
  for {
    // принимаем защищенное соединение
    conn, err := listener.Accept()
    if err != nil {
      log.Println(err)
      return
    }
    go func() {
      // перенаправляем всё что шлёт клиент в stdout
      _, err := io.Copy(os.Stdout, conn)
      if err != nil {
        log.Println(err)
        return
      }
    }()
  }
}
```

Ну что ж, попробуем это всё запустить! Но для наглядности пропустим трафик между клиентом и сервером через TCP прокси, 
который мы реализовали в статье TCP Proxy с логированием на go. Вот так:

```
# наш прокси работает по умолчанию на порту 4242, пусть так и будет
# направим его на наш SSH сервер
$ goproxy --host 127.0.0.1 --port 2022 --skip-healthcheck
# в отдельном окне запустим сервер
$ go run server.go
# запустим клиент, предварительно поставив адрес на прокси localhost:4242 и напишем что-нибудь
$ go run client.go
```
