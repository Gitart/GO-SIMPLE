# Простейший сервер на Gorilla WebSocket

[Go \*](https://habr.com/ru/hub/go/)

Из песочницы

Tutorial

В этом небольшом туториале, мы чуть подробнее разберем использование [Gorilla WebSocket](https://github.com/gorilla/websocket) для написания своего websocket сервера, на примере чуть более функциональном, чем [базовый пример](https://github.com/gorilla/websocket/blob/master/examples/echo/server.go) и более легком для понимания, чем [пример чата](https://github.com/gorilla/websocket/tree/master/examples/chat).

Что будет уметь наш сервер?

1.  Отправлять новые сообщения от клиентов в callback
2.  Хранить активные соединения и закрывать/удалять не активные
3.  Рассылать сообщения по активным соединениям

Для начала поднимем обычный http сервер при помощи [net/http](https://pkg.go.dev/net/http), для того чтобы мы могли отлавливать запросы на соединение:

```go
package main

import (
	"fmt"
	"html"
	"net/http"
)

func main() {
	http.HandleFunc("/", echo)
	http.ListenAndServe(":8080", nil)
}

func echo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
```

Теперь научим его "апгрейдить" соединение:

```go
import (
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Пропускаем любой запрос
	},
}

func echo(w http.ResponseWriter, r *http.Request) {
	connection, _ := upgrader.Upgrade(w, r, nil)
	defer connection.Close() // Закрываем соединение
}
```

Теперь у нас есть соединение с клиентом, которые мы сразу же закрываем. Мы можем циклично читать сообщения, которые нам шлет клиент и отправлять их обратно:

```go
func echo(w http.ResponseWriter, r *http.Request) {
	connection, _ := upgrader.Upgrade(w, r, nil)

	for {
		_, message, _ := connection.ReadMessage()

    connection.WriteMessage(websocket.TextMessage, message)
		go messageHandler(message)
	}
}

func messageHandler(message []byte)  {
  fmt.Println(string(message))
}
```

Научим наш сервер закрывать соединение:

```go
func echo(w http.ResponseWriter, r *http.Request) {
  connection, _ := upgrader.Upgrade(w, r, nil)
  defer connection.Close()

	for {
		mt, message, err := connection.ReadMessage()

		if err != nil || mt == websocket.CloseMessage {
			break // Выходим из цикла, если клиент пытается закрыть соединение или связь с клиентом прервана
		}

		connection.WriteMessage(websocket.TextMessage, message)

		go messageHandler(message)
	}
}
```

Чтобы иметь возможность рассылать сообщения по разным соединениям, нам нужно где то их хранить, в нашем случае подойдет простейший map:

```go
var clients map[*websocket.Conn]bool

func echo(w http.ResponseWriter, r *http.Request) {
	connection, _ := upgrader.Upgrade(w, r, nil)
  defer connection.Close()

	clients[connection] = true // Сохраняем соединение, используя его как ключ
  defer delete(clients, connection) // Удаляем соединение

	for {
		mt, message, err := connection.ReadMessage()

		if err != nil || mt == websocket.CloseMessage {
			break // Выходим из цикла, если клиент пытается закрыть соединение или связь прервана
		}

		// Теперь мы рассылаем сообщения всем клиентам
		go writeMessage(message)

		go messageHandler(message)
	}
}

func writeMessage(message []byte) {
	for conn := range clients {
		conn.WriteMessage(websocket.TextMessage, message)
	}
}
```

Теперь мы можем запаковать наш сервер в структуру, чтобы иметь возможность рассылать и принимать сообщения из вне:

```go
package ws

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Пропускаем любой запрос
	},
}

type Server struct {
	clients       map[*websocket.Conn]bool
	handleMessage func(message []byte) // хандлер новых сообщений
}

func StartServer(handleMessage func(message []byte)) *Server {
	server := Server{
		make(map[*websocket.Conn]bool),
		handleMessage,
	}

	http.HandleFunc("/", server.echo)
	go http.ListenAndServe(":8080", nil) // Уводим http сервер в горутину

	return &server
}

func (server *Server) echo(w http.ResponseWriter, r *http.Request) {
	connection, _ := upgrader.Upgrade(w, r, nil)
  defer connection.Close()

	server.clients[connection] = true // Сохраняем соединение, используя его как ключ
  defer delete(server.clients, connection) // Удаляем соединение

	for {
		mt, message, err := connection.ReadMessage()

		if err != nil || mt == websocket.CloseMessage {
			break // Выходим из цикла, если клиент пытается закрыть соединение или связь прервана
		}

		go server.handleMessage(message)
	}
}

func (server *Server) WriteMessage(message []byte) {
	for conn := range server.clients {
		conn.WriteMessage(websocket.TextMessage, message)
	}
}
```

```go
package main

import (
	"fmt"
	"simple-webcoket/ws"
)

func main() {
	server := ws.StartServer(messageHandler)

	for {
		server.WriteMessage([]byte("Hello"))
	}
}

func messageHandler(message []byte) {
	fmt.Println(string(message))
}
```

Теперь у нас есть реализация простейшего webscoket сервера, который способен принимать и рассылать сообщения по активным соединениям.
