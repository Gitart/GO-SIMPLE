// duck typing

package main

import "fmt"

// Virtual duck
type Duck interface {
	Quack() string
}

type Person struct {
	name string
}

type ActualDuck struct{}

func (p Person) Quack() string {
	return "I am " + p.name
}

func (d ActualDuck) Quack() string {
	return "Quaaack"
}

func test(d Duck) {
	fmt.Println(d.Quack())
}

func main() {
	test(Person{"John"})
	test(ActualDuck{})
}
