package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

func main() {
	client := &http.Client{}
	postData := make([]byte, 100)
	req, err := http.NewRequest("POST", "http://example.com", bytes.NewReader(postData))
	if err != nil {
		os.Exit(1)
	}
	req.Header.Add("User-Agent", "myClient")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	fmt.Println(resp)
}
