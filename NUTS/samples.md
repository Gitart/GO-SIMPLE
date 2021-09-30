Основываясь на вашем обновленном вопросе, вот возможный подход. Обратите внимание, что два дополнительных процесса представлены go-подпрограммами здесь, но вы должны были бы быть отдельными процессами в реальном случае. Я также пропустил проверку ошибок.

```go
// This represent what would be the last process in your
// example.
go func() {
    nc, _ := nats.Connect(nats.DefaultURL)
    nc.Subscribe("bar", func(m *nats.Msg) {
        fmt.Printf("Received request: %s, final stop, sending back to %v\n", m.Data, m.Reply)
        nc.Publish(m.Reply, []byte("I'm here to help!"))
    })
    nc.Flush()
    runtime.Goexit()
}()


// This would be the in-between process that receives
// the message triggered by the TCP accept
go func() {
    nc, _ := nats.Connect(nats.DefaultURL)
    nc.Subscribe("foo", func(m *nats.Msg) {
        fmt.Printf("Received request: %s, forward to bar\n", m.Data)
        nc.PublishRequest("bar", m.Reply, []byte(fmt.Sprintf("got %s", m.Data)))
    })
    nc.Flush()
    runtime.Goexit()
}()



// This would be your TCP server
l, _ := net.Listen("tcp", "127.0.0.1:1234")
for {
    c, _ := l.Accept()
    go func(c net.Conn) {
        // Close socket when done
        defer c.Close()
        // Connect to NATS
        nc, _ := nats.Connect(nats.DefaultURL)
        // Close NATS connection when done
        defer nc.Close()
        // Sends the request to first process. Note that this
        // has a timeout and so if no response is received, the
        // go-routine will exit, closing the TCP connection.
        reply, err := nc.Request("foo", []byte("help"), 10*time.Second)
        if err != nil {
            fmt.Printf("Got error: %v\n", err)
        } else {
            fmt.Printf("Got reply: %s\n", reply.Data)
        }
    }(c)
}
```


Обратите внимание: 
обычно не рекомендуется создавать очень короткие соединения NATS. Возможно, вы захотите повторно использовать соединение NATS, если оно соответствует вашей модели.
