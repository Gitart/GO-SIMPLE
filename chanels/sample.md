
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


## Вариант № 2

```golang
package main
import "fmt"
func FibonacciProducer(ch chan int, count int) {
    n2, n1 := 0, 1
    for count >= 0 {
        ch <- n2
        count--
        n2, n1 = n1, n2+n1
    }
    close(ch)
}
func main() {
    ch := make(chan int)
    go FibonacciProducer(ch, 10)
    idx := 0
    for num := range ch {
        fmt.Printf("F(%d): \t%d\n", idx, num)
        idx++
    }
}
```

Output 

```golang
F(0): 0
F(1): 1
F(2): 1
F(3): 2
F(4): 3
F(5): 5
F(6): 8
F(7): 13
F(8): 21
F(9): 34
F(10): 55
```

### Other
```golang
package main

import "fmt"

func main() {
    ch := make(chan int)

    go func(chan int) {
        for _, v := range []int{1, 2,345,455,567,678,8999,5677} {
            ch <- v
        }
        close(ch)
    }(ch)

    for v := range ch {
        fmt.Println(v)
    }

    fmt.Println("The channel is closed.")
}
```


