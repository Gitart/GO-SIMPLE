package main

import "fmt"

func main() {
	//fmt.Println(greet(""))
	var foo string
	fmt.Println(greet(foo))
	fmt.Printf("foo '%s'\n", foo)

}

func greet(who string) string {
	if who == "" {
		who = "there"
	}
	return fmt.Sprintf("hello, %s!", who)
}
