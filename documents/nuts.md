# NUTS
### Cервер очериди

Краткое описание процесса :    
1. Есть программа - подпсичик     
2. Есть программа - отправитель  
Программа подписчик подписывается на опредленные кналы и слушает, что пошлет отправитель.


Для использования сервера очереди необходимо.   

1. Запустить сам сервер :  
[Инсталяция здесь](http://nats.io/download/)
[Описание](http://nats.io/documentation/tutorials/nats-client-dev/)
2. Запустить подписчика на опредленный канал
3. Запустить отправителя
4. Процесс закончится когда подписчик "поймает" отправленные данные в подписанном канале.


## Запуск сервера
-m 8222   - порт для мониторинга сервреа очереди

```bat
  nats-streaming-server-v0.4.0.exe -m 8222 
  rem nats.exe -m 8222  --log log.txt -dir data 
  pause
```

## Запуск подписчика
Компелируем файл подписчика - тот который слушает свои сообщения в нашем случае "foo"
```golang

package main

import (
  "runtime"
  "log"

  "github.com/nats-io/go-nats"
)

func main() {

    // Create authentication server connection
    natsConnection, _ := nats.Connect(nats.DefaultURL)
    log.Println("Connected to " + nats.DefaultURL)

    // Subscribe to subject
    log.Printf("Subscribing to subject 'foo'\n")
    natsConnection.Subscribe("foo", func(msg *nats.Msg) {

      // Handle the message
      log.Printf("Received message '%s\n", string(msg.Data) + "'")
  })

  // Keep the connection alive
  runtime.Goexit()
}
```

## Запуск отправителя
Компелируем файл клиента кторый посылает подписчику сообщение в нашем случае "Hello NATS"
```golang
package main

import (
  "log"

  "github.com/nats-io/go-nats"
)

func main() {

    // Connect to server with auth credentials
    // natsConnectionString := "nats://foo:bar@localhost:4222"
    natsConnection, _ := nats.Connect(nats.DefaultURL)
    defer natsConnection.Close()
    log.Println("Connected to " + nats.DefaultURL)

    // Publish message on subject
    subject := "foo"
    natsConnection.Publish(subject, []byte("Hello NATS"))
    log.Println("Published message on subject " + subject)
}
```

## Контроль
Можно контролировать состояние сервера и очереди.
```
http://localhost:8222/varz
```

#### На выходе получим JSON
Где :    
"in_msgs": 1303,  -- входящих очередей
"out_msgs": 653,  -- полученых

```
{
"server_id": "w2MbXEAzDQAMBRxcj4xuN2",
"version": "0.9.6",
"go": "go1.7.5",
"host": "0.0.0.0",
"auth_required": false,
"ssl_required": false,
"tls_required": false,
"tls_verify": false,
"addr": "0.0.0.0",
"max_connections": 65536,
"ping_interval": 120000000000,
"ping_max": 2,
"http_host": "0.0.0.0",
"http_port": 8222,
"https_port": 0,
"auth_timeout": 1,
"max_control_line": 1024,
"cluster": {
"addr": "0.0.0.0",
"cluster_port": 0,
"auth_timeout": 1
},
"tls_timeout": 0.5,
"port": 4222,
"max_payload": 1048576,
"start": "2017-04-26T17:22:37.0599924+03:00",
"now": "2017-04-26T17:26:55.3098148+03:00",
"uptime": "4m18s",
"mem": 0,
"cores": 4,
"cpu": 0,
"connections": 3,
"total_connections": 1306,
"routes": 0,
"remotes": 0,
"in_msgs": 1303,
"out_msgs": 653,
"in_bytes": 13065,
"out_bytes": 6530,
"slow_consumers": 0,
"subscriptions": 7,
"http_req_stats": {
"/": 0,
"/connz": 0,
"/routez": 0,
"/subsz": 0,
"/varz": 54
}
}
```

## Мониторинг подключений
Здесь контролируется количество подключенных к серверу клиентов и подписчиков и отправителей.    
http://localhost:8222/connz

## Прочие



