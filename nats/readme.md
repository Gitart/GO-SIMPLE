# NATS
![image](https://user-images.githubusercontent.com/3950155/222965403-17633055-4bc2-477e-990d-34c391c24ddf.png)

```
nats-server -V -m 8222 
```


## Sub

```go
u
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

  
  
  
