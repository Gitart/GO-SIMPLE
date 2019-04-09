// создаем канал res, куда будем получать ответ. А затем, в отдельных горутинах, запускаем запросы к серверам. 
// рация не блокирующая, поэтому после строки с оператором go программа переходит на следующую строку. 
// Далле, программа блокируется на строке data := <- res, ожидая ответа из канала. 
// Как только ответ будет получен, мы выводим его на экран и программа завершается. 
// В данном синтетическом примере будет возвращаться ответ от Server1. 
// Но в жизни, когда выполнение запроса может занимать разное время, 
// будет возвращен ответ от самого быстрого сервера.


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
