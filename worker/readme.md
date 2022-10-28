## 🥧 The Master-Worker Design Pattern

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(wg *sync.WaitGroup, ch chan string, workerName string) {
    defer wg.Done()

    for v := range ch {
        fmt.Println(workerName, "is working on", v)
        // Do stuff
        time.Sleep(time.Second)
    }
}

func master(wg *sync.WaitGroup, ch chan string) {
    defer wg.Done()
    defer close(ch)
    commands := []string{   "walking", 
                            "waiting", 
                            "going forward", 
                            "going backward", 
                            "going left", 
                            "going right"
    }

    for _, v := range commands {
        // Do stuff
        ch <- v
    }
}

func main() {
    var wg sync.WaitGroup
    var ch chan string

    ch = make(chan string)

    wg.Add(3)
    go worker(&wg, ch, "John")
    go worker(&wg, ch, "Ricky")
    go master(&wg, ch)
    wg.Wait()
}
```

Output

```
Ricky is working on walking
John is working on waiting
John is working on going forward
Ricky is working on going backward
John is working on going left
Ricky is working on going right
```

