func Producer() {
    producer, err := nsq.NewProducer("192.168.132.128:4150", nsq.NewConfig())
    if err != nil {
        fmt.Println("NewProducer", err)
        panic(err)
    }

    for i := 0; i < 5; i++ {
        if err := producer.Publish("test", []byte(fmt.Sprintf("Hello World "))); err != nil {
            fmt.Println("Publish", err)
            panic(err)
        }
    }
}
