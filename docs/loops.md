# Golang For Loops

In this tutorial you will learn how to repeat a block of code execution using loops in Golang.

A for loop is used for iterating over a sequence (that is either a slice, an array, a map, or a string.

As a language related to the C\-family, Golang also supports for loop style control structures.

Golang has no while loop because the for loop serves the same purpose when used with a single condition.

---

## Golang \- traditional for Statement

The for loop is used when you know in advance how many times the script should run.

Consider the following example, display the numbers from 1 to 10 in three different ways.

package main

import "fmt"

func main() {

	k := 1
	for ; k <\= 10; k++ {
		fmt.Println(k)
	}

	k = 1
	for k <\= 10 {
		fmt.Println(k)
		k++
	}

	for k := 1; ; k++ {
		fmt.Println(k)
		if k == 10 {
			break
		}
	}
}

---

## Golang \- for range Statement

The for statement supports one additional form that uses the keyword range to iterate over an expression that evaluates to an array, slice, map, string, or channel

package main

import "fmt"

func main() {

	// Example 1
	strDict := map\[string\]string{"Japan": "Tokyo", "China": "Beijing", "Canada": "Ottawa"}
	for index, element := range strDict {
		fmt.Println("Index :", index, " Element :", element)
	}

	// Example 2
	for key := range strDict {
		fmt.Println(key)
	}

	// Example 3
	for \_, value := range strDict {
		fmt.Println(value)
	}
}

---

## Golang \- range loop over string

The for loop iterate over each character of string.

Consider the following example, display "Hello" five times.

package main

import "fmt"

func main() {
	for range "Hello" {
		fmt.Println("Hello")
	}
}

---

## Golang \- Infinite loop

The for loop runs infinite times unless until we can't break.

Consider the following example, display "Hello" several times.

package main

import "fmt"

func main() {
	i := 5
	for {
		fmt.Println("Hello")
		if i == 10 {
			break
		}
		i++
	}
}
