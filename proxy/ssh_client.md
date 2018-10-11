# SSH клиент на go


Продолжаем знакомство с сетевым программированием на go, и сегодня мы сделаем простой ssh клиент, 
подключимся к серверу, выполним команду и получим её вывод.

В статье про netcat мы уже делали выполнение команд на удалённом сервере, 
но проблема в том, что все наши команды, равно как и ответы, шли по сети открытым текстом, 
что, как мы понимаем, откровенно небезопасно.

Что ж, поправим эту ситуацию и в этот раз воспользуемся защищённым SSH соединением, 
авторизоваться будем по ключам (вы ведь всегда используете авторизацию по ключу, правда?).

Для работы с ssh в стандартной библиотеке go нет инструментария, поэтому мы будем использовать 
один из так называемых “экспериментальных” пакетов – golang.org/x/crypto/ssh. Для его установки используем go get:

```
$ go get -u golang.org/x/crypto/ssh
```

Данный пакет предоставляет большое количество вкусных плюшек для облегчения жизни при работе с ssh, 
в том числе и клиент, и сервер, и туннелирование, и форвардинг портов, но мы начнём с простого клиента.

Что ж, приступим. Так как мы решили авторизовываться на сервере правильно, по ключу, то собственно 
для начала нам нужно прочитать файл, в котором хранится наш приватный ключ. Конечно, дадим пользователю 
возможность указать путь к нему явно, но для удобства по умолчанию будем искать его в ~/.ssh/id_rda:

```golang
var pk = flag.String("pk", defaultKeyPath(), "Private key file")

func defaultKeyPath() string {
  // с помощью os.Getenv получаем домашнюю директорию пользователя
  home := os.Getenv("HOME")
  if len(home) > 0 {
    return path.Join(home, ".ssh/id_rsa")
  }
  return ""
}
```

Чтобы прочитать файл полностью, в стандартной библиотеке есть удобный метод – ioutil.ReadFile:

```golang
key, err := ioutil.ReadFile(*pk)
```

Далее нужно создать так называемый Signer – это тот парень, который сделает всю необходимую магию с ключами 
и будет генерировать сигнатуры (отсюда и название) для проверки ключей.


```golang
// передадим ранее созданный ключ key
signer, err := ssh.ParsePrivateKey(key)
```

Теперь можно создавать конфиг для ssh-клиента, только добавим имя пользователя в параметры:

```golang
// ...
var user = flag.String("u", "", "User name")
// ...

config := &ssh.ClientConfig{
  // указываем в конфиге имя пользователя
  User: *user,
  Auth: []ssh.AuthMethod{
    // а тут метод аутентификации по ключам
    ssh.PublicKeys(signer),
  },
}
```

SSH клиент создаётся по аналогии с TCP клиентом, за исключением того что этому парню нужен ещё конфиг, который мы создали выше:

```golang
// ...
var (
  host = flag.String("h", "", "Host")
  port = flag.Int("p", 22, "Port")
)
// ...

addr := fmt.Sprintf("%s:%d", *host, *port)

// звоним на сервер
client, err := ssh.Dial("tcp", addr, config)
Если дозвонились успешно, то можно наконец создать сессию и выполнить долгожданную команду:

// создаём ssh сессию
session, err := client.NewSession()
if err != nil {
  panic(err)
}
// не забываем закрыть перед выходом
defer session.Close()
```

Тип ssh.Session очень похож на Cmd из os/exec, который мы использовали в нашей go версии netcat. У сессии есть всё те же пайпы StdinPipe, StdoutPipe, StderrPipe, но для простоты примера воспользуемся запуском команды с помощью CombinedOutput. Этот метод запустит команду, дождётся её завершения и вернёт весь stdin и stdout.

```golang
// тут происходит запуск uname -a на удалённом сервере
b, err := session.CombinedOutput("uname -a")
if err != nil {
  panic(err)
}
// выводим результат
fmt.Print(string(b))
```

Осталось опробовать ssh-клиент на деле:

```
$ go build
$ ./ssht -h yanzay.com -u core
Linux production 4.7.0-coreos #1 SMP Wed Jul 27 07:30:04 UTC 2016 x86_64 Intel(R) Xeon(R) CPU E5-2630L v2 @ 2.40GHz GenuineIntel GNU/Linux
```

Как говаривал Матроскин: “Ура! Заработало!”

## Полный код ssh-клиента:

```golang
package main

import (
  "flag"
  "fmt"
  "io/ioutil"
  "os"
  "path"

  "golang.org/x/crypto/ssh"
)

var (
  user = flag.String("u", "", "User name")
  pk   = flag.String("pk", defaultKeyPath(), "Private key file")
  host = flag.String("h", "", "Host")
  port = flag.Int("p", 22, "Port")
)

func defaultKeyPath() string {
  home := os.Getenv("HOME")
  if len(home) > 0 {
    return path.Join(home, ".ssh/id_rsa")
  }
  return ""
}

func main() {
  flag.Parse()

  key, err := ioutil.ReadFile(*pk)
  if err != nil {
    panic(err)
  }

  signer, err := ssh.ParsePrivateKey(key)
  if err != nil {
    panic(err)
  }

  config := &ssh.ClientConfig{
    User: *user,
    Auth: []ssh.AuthMethod{
      ssh.PublicKeys(signer),
    },
  }

  addr := fmt.Sprintf("%s:%d", *host, *port)
  client, err := ssh.Dial("tcp", addr, config)
  if err != nil {
    panic(err)
  }

  session, err := client.NewSession()
  if err != nil {
    panic(err)
  }
  defer session.Close()

  b, err := session.CombinedOutput("uname -a")
  if err != nil {
    panic(err)
  }
  fmt.Print(string(b))
}
```
Источник
https://yanzay.com/post/go_ssh_client_example/
