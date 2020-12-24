# Pass different types of arguments in variadic function

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
