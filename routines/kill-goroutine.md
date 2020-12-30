# How to kill execution of goroutine?

A channel has a close operation that closes the channel so that a send operation on the channel cannot take place. A send operation on a closed channel will result in a panic.

When a receive operation is performed on the channel, we check if the channel is closed or not, and exit from the goroutine if the channel is closed.

```jsx
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	ch := make(chan string)
	go func() {
		for {
			channel, ok := <-ch
			if !ok {
				fmt.Println("Shut Down")
				defer wg.Done()
				return
			}
			fmt.Println(channel)
		}
	}()
	ch <- "Start"
	ch <- "Processing"
	ch <- "Finishing"
	close(ch)

	wg.Wait()
}
```

You can see the following output when you run the above program −

```markup

C:\Golang\goroutines>go run main.go
Start
Processing
Finishing
Shut Down
```
