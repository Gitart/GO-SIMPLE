
## Chanels

```golang
package main

import (
"fmt"
"sync"
)

// wg is used to wait for the program to finish.
var wg sync.WaitGroup

func main() {
count := make(chan int)
wg.Add(2)
fmt.Println("Start Goroutines")
go printCounts("Goroutine-1", count)
go printCounts("Goroutine-2", count)
fmt.Println("Communication of channel begins")

count <- 1
fmt.Println("Waiting To Finish")
wg.Wait()
fmt.Println("\nTerminating the Program")
}


func printCounts(label string, count chan int) {
// Schedule the call to WaitGroup's Done to tell goroutine is completed.
defer wg.Done()
for val := range count {
    fmt.Printf("Count: %d received from %s \n", val, label)

    if val == 10 {
       fmt.Printf("Channel Closed from %s \n", label)
       close(count)
       return
     }
val++
count <- val
}
}
```

