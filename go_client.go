package main

import (
	// "fmt"
	"log"
	"net"
	"runtime"
	"strconv"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	for i := 0; i < 10000; i++ {
		go Do(i)
	}
	log.Println("succ")
	time.Sleep(time.Second * 100)
}

func Do(i int) {
	con, err := net.Dial("tcp", "127.0.0.1:20006")

	if err != nil {
		log.Fatal(err)
		return
	}
	defer con.Close()
	con.Write([]byte("this is from client" + strconv.Itoa(i)))
}
