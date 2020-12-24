# Golang Functions

---

A function is a group of statements that exist within a program for the purpose of performing a specific task. At a high level, a function takes an input and returns an output.

Function allows you to extract commonly used block of code into a single component.

The single most popular Go function is **main()**, which is used in every independent Go program.

---

## Creating a Function

A declaration begins with the func keyword, followed by the name you want the function to have, a pair of parentheses (), and then a block containing the function's code.

The following example has a function with the name **SimpleFunction**. It takes no parameter and returns no values.

```jsx
package main

import "fmt"

// SimpleFunction prints a message
func SimpleFunction() {
	fmt.Println("Hello World")
}

func main() {
	SimpleFunction()
}
```

When the above code is compiled and executed, it produces the following result −

```markup
Hello World
```

---

## Function with Parameters

Information can be passed to functions through arguments. An argument is just like a variable.

Arguments are specified after the function name, inside the parentheses. You can add as many arguments as you want, just separate them with a comma.

The following example has a function with two arguments of int type. When the **add()** function is called, we pass two integer values (e.g. 20,30).

```jsx
package main

import "fmt"

// Function accepting arguments
func add(x int, y int) {
	total := 0
	total = x + y
	fmt.Println(total)
}

func main() {
	// Passing arguments
	add(20, 30)
}
```

When the above code is compiled and executed, it produces the following result −

```markup
50
```

If the functions with names that start with an uppercase letter will be exported to other packages. If the function name starts with a lowercase letter, it won't be exported to other packages, but you can call this function within the same package.

---

## Function with Return Type

In this example, the add() function takes input of two integer numbers and returns an integer value with a name of **total**.

Note the return statement is required when a return value is declared as part of the function's signature.

```jsx
package main

import "fmt"

// Function with int as return type
func add(x int, y int) int {
	total := 0
	total = x + y
	return total
}

func main() {
	// Accepting return value in varaible
	sum := add(20, 30)
	fmt.Println(sum)
}
```

The types of input and return value must match with function signature. If we will modify the above program and pass some string value in argument then program will throw an exception "cannot use "test" (type string) as type int in argument to add".

---

## Named Return Values

Golang allows you to name the return values of a function. We can also name the return value by defining variables, here a variable **total** of integer type is defined in the function declaration for the value that the function returns.

```jsx
package main

import "fmt"

func rectangle(l int, b int) (area int) {
	var parameter int
	parameter = 2 * (l + b)
	fmt.Println("Parameter: ", parameter)

	area = l * b
	return // Return statement without specify variable name
}

func main() {
	fmt.Println("Area: ", rectangle(20, 30))
}
```

When the above code is compiled and executed, it produces the following result −

```markup
C:\Golang>go run main.go
Parameter:  100
Area:  600
```

Since the function is declared to return a value of type int, the last logical statement in the execution flow must be a return statement that returns a value of the declared type.

---

## Returning Multiple Values

Functions in Golang can return multiple values, which is a helpful feature in many practical scenarios.

This example declares a function with two return values and calls it from a main function.

```jsx
package main

import "fmt"

func rectangle(l int, b int) (area int, parameter int) {
	parameter = 2 * (l + b)
	area = l * b
	return // Return statement without specify variable name
}

func main() {
	var a, p int
	a, p = rectangle(20, 30)
	fmt.Println("Area:", a)
	fmt.Println("Parameter:", p)
}
```

---

## Points to remember

*   A name must begin with a letter, and can have any number of additional letters and numbers.
*   A function name cannot start with a number.
*   A function name cannot contain spaces.
*   If the functions with names that start with an uppercase letter will be exported to other packages. If the function name starts with a lowercase letter, it won't be exported to other packages, but you can call this function within the same package.
*   If a name consists of multiple words, each word after the first should be capitalized like this: empName, EmpAddress, etc.
*   function names are case\-sensitive (car, Car and CAR are three different variables).

---

## Passing Address to a Function

Passing the address of variable to the function and the value of a variables modified using dereferencing inside body of function.

```jsx
package main

import "fmt"

func update(a *int, t *string) {
	*a = *a + 5      // defrencing pointer address
	*t = *t + " Doe" // defrencing pointer address
	return
}

func main() {
	var age = 20
	var text = "John"
	fmt.Println("Before:", text, age)

	update(&age, &text)

	fmt.Println("After :", text, age)
}
```

You should see the following output when you run the above program −

```markup
C:\Golang>go run main.go
Before: John 20
After : John Doe 25
```

---

## Anonymous Functions

An anonymous function is a function that was declared without any named identifier to refer to it. Anonymous functions can accept inputs and return outputs, just as standard functions do.

**Assigning function to the variable.**

```jsx
package main

import "fmt"

var (
	area = func(l int, b int) int {
		return l * b
	}
)

func main() {
	fmt.Println(area(20, 30))
}
```

**Passing arguments to anonymous functions.**

```jsx
package main

import "fmt"

func main() {
	func(l int, b int) {
		fmt.Println(l * b)
	}(20, 30)
}
```

**Function defined to accept a parameter and return value.**

```jsx
package main

import "fmt"

func main() {
	fmt.Printf(
		"100 (°F) = %.2f (°C)\n",
		func(f float64) float64 {
			return (f - 32.0) * (5.0 / 9.0)
		}(100),
	)
}
```

Anonymous functions can be used for containing functionality that need not be named and possibly for short\-term use.

---

## Closures Functions

Closures are a special case of anonymous functions. Closures are anonymous functions which access the variables defined outside the body of the function.

**Anonymous function accessing the variable defined outside body.**

```jsx
package main

import "fmt"

func main() {
	l := 20
	b := 30

	func() {
		var area int
		area = l * b
		fmt.Println(area)
	}()
}
```

**Anonymous function accessing variable on each iteration of loop inside function body.**

```jsx
package main

import "fmt"

func main() {
	for i := 10.0; i < 100; i += 10.0 {
		rad := func() float64 {
			return i * 39.370
		}()
		fmt.Printf("%.2f Meter = %.2f Inch\n", i, rad)
	}
}

```

---

## Higher Order Functions

A Higher\-Order function is a function that receives a function as an argument or returns the function as output.

Higher order functions are functions that operate on other functions, either by taking them as arguments or by returning them.

### Passing Functions as Arguments to other Functions

```jsx
package main

import "fmt"

func sum(x, y int) int {
	return x + y
}
func partialSum(x int) func(int) int {
	return func(y int) int {
		return sum(x, y)
	}
}
func main() {
	partial := partialSum(3)
	fmt.Println(partial(7))
}
```

You should see the following output when you run the above program −

```markup
C:\Golang>go run main.go
10
```

In the program above, the **partialSum** function returns a **sum** function that takes two int arguments and returns a int argument.

### Returning Functions from other Functions

```jsx
package main

import "fmt"

func squareSum(x int) func(int) func(int) int {
	return func(y int) func(int) int {
		return func(z int) int {
			return x*x + y*y + z*z
		}
	}
}
func main() {
	// 5*5 + 6*6 + 7*7
	fmt.Println(squareSum(5)(6)(7))
}
```

You should see the following output when you run the above program −

```markup
C:\Golang>go run main.go
110
```

In the program above, the **squareSum** function signature specifying that function returns two functions and one integer value.

---

## User Defined Function Types

Golang also support to define our own function types.

The modified version of above program with function types as below:

```jsx
package main

import "fmt"

type First func(int) int
type Second func(int) First

func squareSum(x int) Second {
	return func(y int) First {
		return func(z int) int {
			return x*x + y*y + z*z
		}
	}
}

func main() {
	// 5*5 + 6*6 + 7*7
	fmt.Println(squareSum(5)(6)(7))
}
```
