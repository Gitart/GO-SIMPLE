## Sub

```go
url := os.Getenv("NATS_URL")
	if url == "" {
		fmt.Printf("Two sub")
		url = nats.DefaultURL
	}

	nc, eer := nats.Connect(url)
	if eer != nil {
		fmt.Printf(eer.Error())
	}

	nc.Subscribe("greet", func(msg *nats.Msg) {
	
		fmt.Println("Two:", string(msg.Data))
	})
  ```
  
  
  ## PUB
  
  ```go
  url := os.Getenv("NATS_URL")
	if url == "" {
		fmt.Printf("dddddddddd")
		url = nats.DefaultURL
	}

	nc, eer := nats.Connect(url)
	if eer != nil {
		fmt.Printf(eer.Error())
	}
	defer nc.Drain()

	nc.Publish("greet.Id", []byte("Послал сообщение"))
  ```
  
  
  
  
  
