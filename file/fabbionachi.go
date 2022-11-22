package main

import "fmt"

// Very naive answer.
// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	n := 0
	a := 0
	b := 1
	c := a + b
	return func() int {
		var ret int
		switch {
		case n == 0:
			n++
			ret = 0
		case n == 1:
			n++
			ret = 1
		default:
			ret = c
			a = b
			b = c
			c = a + b
		}
		return ret
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
