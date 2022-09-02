
package main

import (
    "encoding/gob"
    "os"
    "fmt"
    "log"
)


// Main process 
func main() {
  cl()
}


func cl2(in int ){
    fmt.Println("Расчет :", in * 200)
}

func cl33(in int){
    fmt.Println("Возврат cl33 :", in * 200)
}

// Вызов
func cl() {
    
     // Через переменую
     cc:=cl2
     callback(122, cc)

     // Смена функции
     cc=cl33
     callback(5000, cc)

     // На прямую
     callback(122, cl2)
}

 
func callback(y int, f func(int)) {
    f(y)
}
