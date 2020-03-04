https://play.golang.org/p/R_SKPQRMS3w

package main

import (
	"fmt"
	"time"
)




func main() {
t:= make(chan string,2)
lm:= make(chan string,2)

go func(){lm<-"lm 1"}()

go func(){t<-"t-2"}()
go func(){t<-"t-3"}()
go func(){t<-"t-4"}()
go func(){t<-"t-5"}()
go func(){t<-"t-6"}()
go func(){t<-"t-7"}()


// for i := 0; i < 14; i++ {
for{
select {
     case msg1 := <-t: 
           fmt.Println("Первая очередь", msg1)
     case msg2 :=  <-lm:
           fmt.Println("Вторя очередь", msg2) 

     // Без дефолта будет блокировка
     default:  
        // Без этой задержки не будет результата
        time.Sleep(10 * time.Millisecond)
}
}
}


func maina() {
t:= make([]string,2)

   t=append(t,"a")
   t=append(t,"c")
   t=append(t,"l")
   t=append(t,"s")
   t=append(t,"z")

	fmt.Println(t)
}
