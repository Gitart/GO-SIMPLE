package main

import "fmt"

func someFunction1(a, b int) int {
	return a + b
}

func someFunction2(a, b int) int {
	return a - b
}

func someOtherFunction(a, b int, f func(int, int) int) int {
	return f(a, b)
}

func main() {
	fmt.Println(someOtherFunction(111, 12, someFunction1))
	fmt.Println(someOtherFunction(111, 12, someFunction2))
}
