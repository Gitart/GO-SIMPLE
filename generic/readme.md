## Generic
Generic type constraints
Assume you have a generic function that has many type parameters, sayint, int8 , int26, int32, int64, float32, float64. Our function signature will be very long and undesirable to the eye. It is possible to put all these type parameters into one type using an interface.Essentially , we are moving the union from the function declaration into a new type constraint.

```go
package main
 
import "fmt"
 
type Number interface {
   int64 | float64
}
 
func main() {
   f := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
   i := []int64{1, 2, 3, 4, 5}
 
   s1 := genericSum(f)
   s2 := genericSum(i)
 
   fmt.Println("Sum for float64 :", s1)
   fmt.Println("Sum for int64 :", s2)
 
}
 
func genericSum[N Number](nums []N) N {
   var sum N
 
   for _, num := range nums {
       sum += num
   }
 
   return sum
}
```

