package main

import (
    "fmt"
    "math/rand"
)

func main() {
    // Loop five times.
    for i := 0; i < 5; i++ {
	// Get random positive integer.
	value := rand.Int()
	fmt.Println(value)
    }
}
