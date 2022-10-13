## Examples work with goroutines and chanels

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan interface{}, 1)
	
	go func() {
		for a := range c {
			fmt.Println(a)
		}
	}()
	
	c <- 21
	c <- "jazz"
	c <- person{"Chet", 88}

	time.Sleep(time.Second)
}

type person struct {
	Name string
	Age int
}
```
