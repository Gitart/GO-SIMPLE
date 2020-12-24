# Higher Order Functions in Golang

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
