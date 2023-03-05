package main

import (
	"encoding/json"
	"fmt"
	echo "github.com/labstack/echo/v4"
	"github.com/nats-io/nats.go"
	"os"
	"time"
)

type Projects struct {
	Id      int64  `json:"id"`
	Company string `json:"company"`
	Name    string `json:"name"`
}

func main() {
	Subb()

	// Start echo package
	e := echo.New()
	pg := e.Group("/")

	pg.GET("", HomePage)
	pg.GET("tt", HomePage1)

	e.Logger.Fatal(e.Start(":8081"))

}

func Subb() {
	nnc := 0

	url := os.Getenv("NATS_URL")
	if url == "" {
		url = nats.DefaultURL
	}

	nc, eer := nats.Connect(url)
	if eer != nil {
		fmt.Printf(eer.Error())
	}

	// Scrube to ID
	nc.Subscribe("greet.Id", func(msg *nats.Msg) {

		fmt.Println("Id", string(msg.Data))
	})

	// Scribe to name
	nc.Subscribe("greet.Name", func(msg *nats.Msg) {
		nnc++
		AddDb(msg.Data)
	})
}

func AddDb(b []byte) {

	dat := Projects{}
	err := json.Unmarshal(b, &dat)
	if err != nil {
		fmt.Println("ERROR JSON: ", err.Error)
	}

	res := DB.Create(&dat)
	if res.Error != nil {
		fmt.Println("ERROR DB: ", res.Error)
	}
}

// Nats
func HomePage(e echo.Context) error {
	url := os.Getenv("NATS_URL")
	if url == "" {
		url = nats.DefaultURL
	}

	nc, eer := nats.Connect(url)
	if eer != nil {
		fmt.Printf(eer.Error())
	}
	defer nc.Drain()

	nc.Publish("greet.Id", []byte("Послал сообщение"))

	return e.JSON(200, "ok")
}

// Nats
func HomePage1(e echo.Context) error {
	url := os.Getenv("NATS_URL")
	if url == "" {
		fmt.Printf("Two sub")
		url = nats.DefaultURL
	}

	nc, eer := nats.Connect(url)
	if eer != nil {
		fmt.Printf(eer.Error())
	}
	//defer nc.Drain()

	nc.Subscribe("greet", func(msg *nats.Msg) {
		//name := msg.Subject[6:]
		//msg.Respond([]byte("hello, " + name))
		fmt.Println("Two:", string(msg.Data))
	})

	//rep, _ := nc.Request("greet.joe", nil, time.Second)
	//fmt.Println(string(rep.Data))

	//sub.Unsubscribe()
	//sub.NextMsg(time.Second * 1)
	//
	//rep, _ := nc.Request("greet.joe", nil, time.Second)
	//fmt.Println(string(rep.Data))

	return e.JSON(200, "ok")
}

func main1() {

	url := os.Getenv("NATS_URL")
	if url == "" {
		url = nats.DefaultURL
	}
	nc, _ := nats.Connect(url)
	defer nc.Drain()

	nc.Publish("greet.joe", []byte("hello"))

	sub, _ := nc.SubscribeSync("greet.*")
	msg, _ := sub.NextMsg(10 * time.Millisecond)
	fmt.Println("subscribed after a publish...")
	fmt.Printf("msg is nil? %v\n", msg == nil)

	nc.Publish("greet.joe11", []byte("hello"))
	nc.Publish("greet.pam11", []byte("hello"))
	nc.Publish("greet.tamm", []byte("hello tamm"))

	msg, _ = sub.NextMsg(10 * time.Millisecond)
	fmt.Printf("msg data: %q on subject %q\n", string(msg.Data), msg.Subject)

	msg, _ = sub.NextMsg(10 * time.Millisecond)
	fmt.Printf("msg data: %q on subject %q\n", string(msg.Data), msg.Subject)

	nc.Publish("greet.bob", []byte("hello"))
	msg, _ = sub.NextMsg(10 * time.Millisecond)
	fmt.Printf("msg data: %q on subject %q\n", string(msg.Data), msg.Subject)
}
