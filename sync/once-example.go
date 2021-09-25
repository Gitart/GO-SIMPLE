package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var once sync.Once

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("http handler start")
		once.Do(oneTimeOp)
		fmt.Println("http handler end")
		w.Write([]byte("done!"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func oneTimeOp() {
	fmt.Println("one time op start")
	time.Sleep(3 * time.Second)
	fmt.Println("one time op end")
}
