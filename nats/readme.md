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
		url = nats.DefaultURL
	}

	nc, eer := nats.Connect(url)
	if eer != nil {
		fmt.Printf(eer.Error())
	}
	defer nc.Drain()

	nc.Publish("greet.Id", []byte("Послал сообщение"))
  ```
  
  ## Sub & create to db
  ```go
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
```

  
  
  
