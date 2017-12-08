package main

import "sync"
func main() {
    var wg sync.WaitGroup
    wg.Add(1)

    ch := make(chan int)
    go func() {
        for {
            foo, ok := <- ch
            if !ok {
                println("done")
                wg.Done()
                return
            }
            println(foo)
        }
    }()
    ch <- 1
    ch <- 2
    ch <- 3
    close(ch)

    wg.Wait()
}
