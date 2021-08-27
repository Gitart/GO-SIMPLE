package main

import "fmt"

type fn func(int) 

func myfn1(i int) {
	fmt.Printf("\ni is %v", i)
}
func myfn2(i int) {
	fmt.Printf("\ni is %v", i)
}
func test(f fn, val int) {
	f(val)
}
func main() {
	test(myfn1, 123)
	test(myfn2, 321)
}
