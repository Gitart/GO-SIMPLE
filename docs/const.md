# Golang Constants

---

A constant is a name or an identifier for a fixed value. The value of a variable can vary, but the value of a constant must remain constant.

---

## Declaring (Creating) Constants

The keyword const is used for declaring constants followed by the desired name and the type of value the constant will hold. You must assign a value at the time of the constant declaration, you can't assign a value later as with variables.

```jsx
package main

import "fmt"

const PRODUCT string = "Canada"
const PRICE = 500

func main() {
	fmt.Println(PRODUCT)
	fmt.Println(PRICE)
}
```

You can also omit the type at the time the constant is declared. The type of the value assigned to the constant will be used as the type of that variable.

---

## Multilple Constants Declaration Block

Constants declaration can to be grouped together into blocks for greater readability and code quality.

```jsx
package main

import "fmt"

const (
	PRODUCT  = "Mobile"
	QUANTITY = 50
	PRICE    = 50.50
	STOCK  = true
)

func main() {
	fmt.Println(QUANTITY)
	fmt.Println(PRICE)
	fmt.Println(PRODUCT)
	fmt.Println(STOCK)
}
```

---

## Naming Conventions for Golang Constants

Name of constants must follow the same rules as variable names, which means a valid constant name must starts with a letter or underscore, followed by any number of letters, numbers or underscores.
By convention, constant names are usually written in uppercase letters. This is for their easy identification and differentiation from variables in the source code.
