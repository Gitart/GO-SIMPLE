## Min Max 

```golang
package main

import (
    "fmt"
    "math"
)



func main() {
	Mathmax()
	Mathmin()
}


// Max math 
func Mathmax(){
    ints := []float64{10.9,1.01, 2.02, 3.78, 4.88, 5.89}
    max  := 0.00

    for _, n := range ints {
        max = math.Max(max, n)
    }

    fmt.Println(max)
}


// Min math 
func Mathmin(){
    ints := []float64{10.9,1.01, 2.02, 3.78, 4.88, 5.89}
    min  := 123.00

    for _, n := range ints {
        min = math.Min(min, n)
    }

    fmt.Println(min)
}


```
