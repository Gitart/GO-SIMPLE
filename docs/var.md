# Golang Variables

---

In Golang, a variable holds data temporarily to work with it. A golang variable declaration, needs four things: a statement that declaring a golang variable, a name for the variable, the type of data it can hold, and an initial value for it. Fortunately, some of the parts are optional, but that also means there's more than one way of defining a variable in golang. Here're some important things to know about golang variables:

*   Golang is statically typed language, this means that when golang variables are declared, they either explicitly or implicitly assigned a type even before your program runs.
*   Golang requires that every variable you declare inside `main()` function get used somewhere in your program.
*   You can assign new value to an existing variable, but the value need to be of same type.
*   A variable declared within brace brackets `{}` may be accessed anywhere within the block. The opening curly brace `{` introduces a new scope that ends with a closing brace `}`. Inner blocks can access variables within outer blocks. Outer blocks cannot access variables within inner blocks.

---

## Declaring Golang Variables

The keyword var is used for declaring variables followed by the desired name and the type of value the variable will hold.

You can declare a variable without assigning the value, and assign the value later.

```jsx
package main

import "fmt"

func main() {
	var i int
	var s string

	i = 10
	s = "Canada"

	fmt.Println(i)
	fmt.Println(s)
}
```

After the execution of the statements above, the variable i will hold the value 10 and the variable s will hold the value Canada.

---

## Declaration and Initialization of Golang Variable

The assignment of a value occurred inline with the initialization of the variable. It is equally valid to declare a variable and assign a value to it later.

```jsx
package main

import "fmt"

func main() {
	var i int = 10
	var s string = "Canada"

	fmt.Println(i)
	fmt.Println(s)
}
```

The integer literal 10 is assigned to the variable i and the string literal Canada is assigned to the variable s.

---

## Variable Declaration Omit Types

You can omit the variable type from the declaration, when you are assigning a value to a variable at the time of declaration. The type of the value assigned to the variable will be used as the type of that variable.

```jsx
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i = 10
	var s = "Canada"

	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(s))
}
```

---

## Short Variable Declaration in Golang

The := short variable assignment operator indicates that short variable declaration is being used. There is no need to use the var keyword or declare the variable type.

```jsx
package main

import (
	"fmt"
	"reflect"
)

func main() {
	name := "John Doe"
	fmt.Println(reflect.TypeOf(name))
}
```

The John Doe string literal will be assigned to name.

---

## Declare Multiple Variables

Golang allows you to assign values to multiple variables in one line.

```jsx
package main

import (
	"fmt"
)

func main() {
	var fname, lname string = "John", "Doe"
	m, n, o := 1, 2, 3
	item, price := "Mobile", 2000

	fmt.Println(fname + lname)
	fmt.Println(m + n + o)
	fmt.Println(item, "-", price)
}
```

---

## Scope of Golang Variables Defined by Brace Brackets

Golang uses lexical scoping based on code blocks to determine the scope of variables. Inner block can access its outer block defined variables, but outer block cannot access inner block defined variables.

```jsx
package main

import (
	"fmt"
)

var s = "Japan"

func main() {
	fmt.Println(s)
	x := true

	if x {
		y := 1
		if x != false {
			fmt.Println(s)
			fmt.Println(x)
			fmt.Println(y)
		}
	}
	fmt.Println(x)
}
```

Note that short variable declaration is allowed only for declaring local variables, variables declared within the function. When you declare variables outside the function, you must do so using the var keyword.

---

## Naming Conventions for Golang Variables

These are the following rules for naming a Golang variable:

*   A name must begin with a letter, and can have any number of additional letters and numbers.
*   A variable name cannot start with a number.
*   A variable name cannot contain spaces.
*   If the name of a variable begins with a lower\-case letter, it can only be accessed within the current package this is considered as unexported variables.
*   If the name of a variable begins with a capital letter, it can be accessed from packages outside the current package one this is considered as exported variables.
*   If a name consists of multiple words, each word after the first should be capitalized like this: empName, EmpAddress, etc.
*   Variable names are case\-sensitive (car, Car and CAR are three different variables).

---

## Zero Values

If you declare a variable without assigning it a value, Golang will automatically bind a default value (or a zero\-value) to the variable.

```jsx
package main

import "fmt"

func main() {
	var quantity float32
	fmt.Println(quantity)

	var price int16
	fmt.Println(price)

	var product string
	fmt.Println(product)

	var inStock bool
	fmt.Println(inStock)
}
```

When the above code is compiled and executed, it produces the following result −

```markup
0
0

false
```

---

## Golang Variable Declaration Block

Variables declaration can be grouped together into blocks for greater readability and code quality.

```jsx
package main

import "fmt"

var (
	product  = "Mobile"
	quantity = 50
	price    = 50.50
	inStock  = true
)

func main() {
	fmt.Println(quantity)
	fmt.Println(price)
	fmt.Println(product)
	fmt.Println(inStock)
}
```
