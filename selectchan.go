package main

import "fmt"

func sender(out chan int, exit chan bool) {
    for i := 1; i <= 10; i++ {
        out <- i
    }
    exit <- true
}

func main() {
    out := make(chan int, 10)
    exit := make(chan bool)

    go sender(out, exit)

    for {
        select {
        case i := <-out:
            fmt.Printf("Value: %d\n", i)
            continue
        default:
        }
        select {
        case i := <-out:
            fmt.Printf("Value: %d\n", i)
            continue
        case <-exit:
            fmt.Println("Exiting")
        }
        break
    }
    fmt.Println("Did we get all 10? I think so!")
}
