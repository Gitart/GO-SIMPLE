package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	data := url.Values{
		"url":       {"https://audd.tech/example.mp3"},
		"return":    {"apple_music,spotify"},
		"api_token": {"test"},
	}
	response, _ := http.PostForm("https://api.audd.io/", data)
	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
