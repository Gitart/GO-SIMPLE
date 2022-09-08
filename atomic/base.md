```go
// Golang program to illustrate the usage of
// AddInt64 function

// Including main package
package main

// importing fmt and sync/atomic
import (
	"fmt"
	"sync/atomic"
)

// Main function
func main() {

	// Assigning values
	// to the int64
	var (
		s int64 = 67656
		t int64 = 90
		u int64 = 922337203685477580
		v int64 = -9223372036854775807
	)

	// Assigning constant
	// values to int64
	const (
		w int64 = 5
		x int64 = 8
	)

	// Calling AddInt64 method
	// with its parameters
	output_1 := atomic.AddInt64(&s, w)
	output_2 := atomic.AddInt64(&t, x-w)
	output_3 := atomic.AddInt64(&u, x-6)
	output_4 := atomic.AddInt64(&v, -x)

	// Displays the output after adding
	// addr and delta automatically
	fmt.Println(output_1)
	fmt.Println(output_2)
	fmt.Println(output_3)
	fmt.Println(output_4)
}

```
