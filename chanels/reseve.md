# Channels Send vs Receive

---

*   Channels declared as <-chan can only receive
*   Channels declared as chan<- can only send

```
package main

import (
    "fmt"
    "sync"
)

func read_ch(ch <-chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    for i := range ch {
        fmt.Println("Read", i, "from channel")
    }

    x, ok := <- ch                  // Closed channel returns 0
    fmt.Println("ch after closing:", x, ok)
}

func write_ch(ch chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 1; i < 5; i++ {
        ch <- i
    }
    close(ch)
}

func main() {
    var ch chan int
    ch = make(chan int)

    wg := new(sync.WaitGroup)

    wg.Add(2)
    go write_ch(ch, wg)
    go read_ch(ch, wg)
    wg.Wait()
}

```

Output

```
Read 1 from channel
Read 2 from channel
Read 3 from channel
Read 4 from channel
ch after closing: 0 false
```
