# Channels Blocking Behavior

---

*   Read on open empty channels block
*   Write on full buffered channels block (unbuffered channel can be thought of being buffered channel of size=1)
*   Nil channels always block
*   Closed channels never block and return 0

```
package main

import (
    "fmt"
    "time"
)

func main() {
    ch := make(chan int, 2)

    go func(ch chan int) {
        ch <- 1
        fmt.Println("<Sleeping for 10s>")
        time.Sleep(time.Second * 10)
        fmt.Println("<Awake again>")
        ch <- 2
        close(ch)
    }(ch)

    fmt.Println("Looping and reading from channel.")
    for i := range ch {
        fmt.Println("In loop:", i)
    }

    i, ok := <- ch
    fmt.Println("Out of loop:", i, ok)
}

```

Output

```
Looping and reading from channel.
In loop: 1
<Sleeping for 10s>
<Awake again>
In loop: 2
Out of loop: 0 false
```
