package main

import "fmt"

func getDataFromServer(resultCh chan string, serverName string) {
	resultCh <- "Data from server: " + serverName
}

func main() {
	res := make(chan string, 3)
	go getDataFromServer(res, "Server1")
	go getDataFromServer(res, "Server2")
	go getDataFromServer(res, "Server3")

	data := <- res
	fmt.Println(data)
}
