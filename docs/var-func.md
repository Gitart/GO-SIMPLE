# Variadic Functions

A variadic function is a function that accepts a variable number of arguments. In Golang, it is possible to pass a varying number of arguments of the same type as referenced in the function signature. To declare a variadic function, the type of the final parameter is preceded by an ellipsis, "...", which shows that the function may be called with any number of arguments of this type. This type of function is useful when you don't know the number of arguments you are passing to the function, the best example is built\-in Println function of the fmt package which is a variadic function.

## Select single argument from all arguments of variadic function.

In below example we will are going to print s\[0\] the first and s\[3\] the forth, argument value passed to `variadicExample()` function.

package main

import "fmt"

func main() {
	variadicExample("red", "blue", "green", "yellow")
}

func variadicExample(s ...string) {
	fmt.Println(s\[0\])
	fmt.Println(s\[3\])
}

C:\\golang\\example>go run test1.go
red
yellow

C:\\golang\\example>

Needs to be precise when running an empty function call, if the code inside of the function expecting an argument and absence of argument will generate an error "panic: run\-time error: index out of range". In above example you have to pass at least 4 arguments.

---

## Passing multiple string arguments to a variadic function

The parameter `s` accepts an infinite number of arguments. The tree\-dotted `ellipsis` tells the compiler that this string will accept, from zero to multiple values.

package main

import "fmt"

func main() {

	variadicExample()
	variadicExample("red", "blue")
	variadicExample("red", "blue", "green")
	variadicExample("red", "blue", "green", "yellow")
}

func variadicExample(s ...string) {
	fmt.Println(s)
}

C:\\golang\\example>go run test1.go
\[\]
\[red blue\]
\[red blue green\]
\[red blue green yellow\]

C:\\golang\\example>

In the above example, we have called the function with single and multiple arguments; and without passing any arguments.

---

## Normal function parameter with variadic function parameter

package main

import "fmt"

func main() {
	fmt.Println(calculation("Rectangle", 20, 30))
	fmt.Println(calculation("Square", 20))
}

func calculation(str string, y ...int) int {

	area := 1

	for \_, val := range y {
		if str == "Rectangle" {
			area \*= val
		} else if str == "Square" {
			area = val \* val
		}
	}
	return area
}

C:\\golang\\example>go run test1.go
600
400

C:\\golang\\example>

---

## Pass different types of arguments in variadic function

In the following example, the function signature accepts an arbitrary number of arguments of type `slice`.

package main

import (
	"fmt"
	"reflect"
)

func main() {
	variadicExample(1, "red", true, 10.5, \[\]string{"foo", "bar", "baz"},
		map\[string\]int{"apple": 23, "tomato": 13})
}

func variadicExample(i ...interface{}) {
	for \_, v := range i {
		fmt.Println(v, "\-\-", reflect.ValueOf(v).Kind())
	}
}

C:\\golang\\example>go run test3.go
1 \-\- int
red \-\- string
true \-\- bool
10.5 \-\- float64
\[foo bar baz\] \-\- slice
map\[apple:23 tomato:13\] \-\- map

C:\\golang\\example>
