package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Student struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func main() {

	body := &Student{
		Name:    "abc",
		Address: "xyz",
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)
	req, _ := http.NewRequest("POST", "https://httpbin.org/post", buf)

	client := &http.Client{}
	res, e := client.Do(req)
	if e != nil {
		log.Fatal(e)
	}

	defer res.Body.Close()

	fmt.Println("response Status:", res.Status)

	// Print the body to the stdout
	io.Copy(os.Stdout, res.Body)
}


// response Status: 200 OK
// {
// 	"args": {},
// 	"data": "{\"name\":\"abc\",\"address\":\"xyz\"}\n",
// 	"files": {},
// 	"form": {},
// 	"headers": {
// 		"Accept-Encoding": "gzip",
// 		"Content-Length": "31",
// 		"Host": "httpbin.org",
// 		"User-Agent": "Go-http-client/1.1"
// 	},
// 	"json": {
// 		"address": "xyz",
// 		"name": "abc"
// 	},
// 	"origin": "118.127.110.2",
// 	"url": "https://httpbin.org/post"
// }
