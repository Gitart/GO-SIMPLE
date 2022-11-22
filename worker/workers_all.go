package main

import (
    "fmt"
    "time"
)

// Определяем глобальную переменную очереди
var quit = make(chan bool)


// *****************************************
// Запуск рутин при инициализации
// *****************************************
func main() {
    fmt.Println("Start")
    go Worker1()

    Worker2()
    fmt.Println("Start end")
}


// *************************************
// Работа воркера 1
// **************************************
func Worker1() {
     fmt.Println("Start w1")

for{
      select {
        case <- quit:
             fmt.Println("Canceled")
             return
        default:
             fmt.Println("Worker 1")
        }

       time.Sleep(time.Second*5)
   }
}

// Здесь посылаем через 20 сек прерывание Worker1
func Worker2() {

    fmt.Println("Старт воркера 2")
    time.Sleep(time.Second*20)
    quit <- true
    fmt.Println("Worker2 itercept Worker ----------")
}


