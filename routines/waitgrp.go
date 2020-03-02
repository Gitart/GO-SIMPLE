package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {

    total := 3

    // Create the waitgroup and add the total number of goroutines we're going to use
    var wg sync.WaitGroup
    wg.Add(total)

    for i := 1; i <= total; i++ {
        // Spin off the goroutine and pass the waitgroup by reference (not value!)
        go longConcurrentProcess(i, &wg)
    }

    // Wait for all goroutines to be finished
    wg.Wait()
    fmt.Println("Finished")
}

func longConcurrentProcess(sleep int, wg *sync.WaitGroup) {

    // Call Done() using defer as it's be easiest way to guarantee it's called at every exit
    defer wg.Done()

    time.Sleep(time.Duration(sleep) * time.Second)
    fmt.Println("Sleeping for", sleep, "seconds")
}                   
