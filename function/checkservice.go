package main

import (
	"fmt"
	"net/http"
	"time"
)

const url = "https://golang.org/"

func main() {
	check_loop()
}

func check_loop() {
	for {
		check(url)
		//fmt.Println("Firm....")
		time.Sleep(1000 * time.Millisecond)
	}
}

func check(url string) {
	fmt.Println("Проверяем адрес ", url)
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("Ошибка соединения. %s\n", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Printf("Ошибка. http-статус: %s\n", resp.StatusCode)
		return
	}

	fmt.Printf("Онлайн. http-статус: %d\n", resp.StatusCode)
}

func mains() {
	fmt.Println("sss")

}
