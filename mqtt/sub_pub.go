// https://www.emqx.com/en/blog/how-to-use-mqtt-in-golang

package main

import (
    "fmt"
    mqtt "github.com/eclipse/paho.mqtt.golang"
    // "log"
    "time"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
    fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
    fmt.Println("Connected ...")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
    fmt.Printf("Connect lost : %v", err)
}

func main() {
    // var broker = "broker.emqx.io"
    var broker = "127.0.0.1"
    var port   = 1883
    opts := mqtt.NewClientOptions()
    opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
    opts.SetClientID("go_mqtt_client")
    opts.SetUsername("emqx")
    opts.SetPassword("public")

    // Show process
    opts.SetDefaultPublishHandler(messagePubHandler)
    opts.OnConnect        = connectHandler
    opts.OnConnectionLost = connectLostHandler

    client   := mqtt.NewClient(opts)
    if token := client.Connect(); token.Wait() && token.Error() != nil {
       panic(token.Error())
    }

    // Сначала подписка
    sub(client)
    sub1(client)

    // Потом публикация
    // publish(client)

    // Закрытие
    client.Disconnect(250)
}

// Публикация
func publish(client mqtt.Client) {
    num := 10
    for i := 0; i < num; i++ {
        text  := fmt.Sprintf("Message %d", i)
        token := client.Publish("topic/test", 0, false, text)
        token.Wait()
        time.Sleep(time.Second)
    }

    num=12

    for i := 0; i < num; i++ {
        text  := fmt.Sprintf("Послано в очередь %d\t", i)
        token := client.Publish("topic/weather", 0, false, text)
        token.Wait()
         time.Sleep(time.Millisecond * 100)
    }
}


// Подписка # 1
func sub(client mqtt.Client) {
    topic := "topic/test"
    token := client.Subscribe(topic, 1, nil)
    token.Wait()
    fmt.Printf("Subscribed to topic: %s\n", topic)
}


// Подписка # 2
func sub1(client mqtt.Client) {
    topic := "topic/weather"
    token := client.Subscribe(topic, 1, nil)
    token.Wait()
    fmt.Printf("Subscribed to topic: %s\n", topic)
}
