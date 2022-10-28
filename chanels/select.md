Channels Select
Select

If at least one channel is available for operation, randomly select channel to operate on
package main

import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)

    go func(ch chan int) { <-ch }(ch1)
    go func(ch chan int) { ch <- 2 }(ch2)

    time.Sleep(time.Second) // Force blocking of main goroutine so other goroutines can run

    for {
        select {
        case ch1 <- 1:
            fmt.Println("Was able to send to ch1")
        case x := <-ch2:
            fmt.Println("was able to read", x, "from ch2")
        default:
            fmt.Println("Neither channels available. Exiting.")
            return
        }
    }
}
Output

Was able to send to ch1
was able to read 2 from ch2
Neither channels available. Exiting.
