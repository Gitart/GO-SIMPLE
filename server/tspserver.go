package main

import (
	"fmt"
	"net"
	"net/http"
	"net/http/fcgi"
)

func handler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello World!")
}

func main() {
	// For local machine
	// l, _ := net.Listen("unix", "/var/run/go-fcgi.sock")

	l, err := net.Listen("tcp", "0.0.0.0:5000") // TCP 5000 listen
	if err != nil {
		return
	}
	http.HandleFunc("/", handler)
	fcgi.Serve(l, nil)
}
