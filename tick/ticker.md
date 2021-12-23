go Ticker()

```go
func Ticker(){
	       for range time.Tick(time.Second * 2) {
              GlobalCount=0 	
           }
}
```



```go
package main

import (
	"fmt"
	"time"
)

func main() {
	go heartBeat()
	time.Sleep(time.Second * 15)
}
func heartBeat() {
	for range time.Tick(time.Second * 1) {
		fmt.Println("Foo")
	}
}
```
