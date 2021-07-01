## Go Chanel 

```go
package main

import (
  "fmt"
  "time"
  "math/rand"
)

type Worker struct {
  id int
}


func (w *Worker) process(c chan int) {
  for {
    data := <-c
    fmt.Printf("обработчик %d получил %d\n", w.id, data)
  }
}


func main() {
  c := make(chan int)

  for i := 0; i < 5; i++ {
      worker := &Worker{id: i}
      go worker.process(c)
  }

  for {
    c <- rand.Int()
    time.Sleep(time.Millisecond * 50)
  }
}
```



