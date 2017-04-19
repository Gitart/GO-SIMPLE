
## Timer

```golang
package main

import "time"
import "fmt"


func timer(ch chan string, ns, count int) {

for j := 1; j <= count; j++ {
       time.Sleep(time.Duration(ns) * time.Nanosecond)
fmt.Println(j)

       if j == count {
                fmt.Printf("[timer] Отправляю последнее сообщение...\n")
                ch <- "стоп!"
        } else {
                fmt.Printf("[timer] Отправляю...\n")
                ch <- "продолжаем"
        }
                  fmt.Printf("[timer] Отправил!\n")
     }
 }


 func main() {
	 var str string
	 ch := make(chan string)
	 go timer(ch, 1000000000, 10)

	 for {
	      fmt.Printf("[main] Принимаю...\n")
	      str = <-ch

	      if str == "стоп!" {
	         fmt.Printf("[main] Принял последнее сообщение,завершаю работу.\n")
	         return
	      } else {
	         fmt.Printf("[main] Принято!\n")
	      }
	     }
 }
 ```
 
