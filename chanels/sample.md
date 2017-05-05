
##    Буферезированные

```golang

/*
  ■
  ■ Title каналы
  ■ Буферезированные
  ■ 03-04-2017
  ■ 15:32
  ■
*/

package main

import (
     "fmt"
) 


func main() {

 counter := make(chan int)       // Declare a unbuffered channel
 nums    := make(chan int, 3)    // Declare a buffered channel with capacity of 3


 go func() {
	 // Send value to the unbuffered channel
	 counter <- 1
	 close(counter) // Closes the channel
 }()

 go func() {
	 // Send values to the buffered channel
	 nums <- 10
	 nums <- 30
	 nums <- 50
 }()

 // Read the value from unbuffered channel
 fmt.Println(<-counter)

 val, ok := <-counter // Trying to read from closed channel

 if ok {
    fmt.Println(val) // This won't execute
 }
 
 // Read the 3 buffered values from the buffered channel
 fmt.Println(<-nums)
 fmt.Println(<-nums)
 fmt.Println(<-nums)
 
 close(nums) // Closes the channel
} 
```

## Обычные

```golang
package main
import "fmt"


func FP(ch chan string, n2 string) {
    ch <- n2
 
    //close(ch)
}


func main() {
    ch := make(chan string)

    go FP(ch, "1-ddd")
    go FP(ch, "2-dd4d")
    go FP(ch, "3-dd456d")


    go FP(ch, "4-dd456d")
    go FP(ch, "5-dd456d")
    go FP(ch, "6-dd456d")

fmt.Println("Cap", len(ch))
   
   // red(ch)  
    red(ch)  
    close(ch)
}


func red(ch chan string){
   
   r:=<-ch
    fmt.Println(r)
    s:=<-ch
    fmt.Println(s)
    t:=<-ch
    fmt.Println(t)

 
}
```

