```go
package main

import "fmt"

type Arith struct {
}

func (Arith) Multiply(a float32, b float32) float32 {
    return a * b
}

func main() {
    result := (Arith).Multiply(Arith{}, 15, 25)
    fmt.Println(result)
}
```
