// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"reflect"
)

type Tr struct {
	Tr string
}

func (s Tr) IsEmpty() bool {
	return reflect.DeepEqual(s, Tr{})
}

func main() {

	A := Tr{}

	if A.IsEmpty() {
		fmt.Println("Is Empty")
	}

}
