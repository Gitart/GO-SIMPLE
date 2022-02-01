## Changed data in structure 

```go
package main

import "fmt"

type Employee struct {
	name   string
	salary int
}

func (e *Employee) changeName(newName string) {
	(*e).name = newName
}

func main() {
	e := Employee{
		name:   "Ross Geller",
		salary: 1200,
	}

	// e before name change
	fmt.Println("e before name change =", e)
	// create pointer to `e`
	ep := &e
	// change name
	ep.changeName("Monica Geller")
	// e after name change
	fmt.Println("e after name change =", e)
}
```

**Output**
```
e before name change = {Ross Geller 1200}
e after name change = {Monica Geller 1200}
```
