## Log 

```go
package main

import (
    "log"
    "os"
    "sync"
    "time"
)

func main() {
    logFile, err := os.OpenFile("clog", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
    if err != nil {
        log.Fatalln(err)
    }
    log.SetOutput(logFile)

    var wg sync.WaitGroup

    wg.Add(3)

    go func() {
        defer wg.Done()
        for i := 0; i < 10; i++ {
            log.Println("F1; loop:", i)
            time.Sleep(time.Millisecond)
        }
    }()

    go func() {
        defer wg.Done()
        for i := 0; i < 10; i++ {
            log.Println("F2; loop:", i)
            time.Sleep(time.Millisecond * 2)
        }
    }()

    go func() {
        defer wg.Done()
        for i := 0; i < 10; i++ {
            log.Println("F3; loop:", i)
            time.Sleep(time.Millisecond * 3)
        }
    }()

    wg.Wait()

}
```
