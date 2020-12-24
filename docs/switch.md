# Golang Switch…Case Statements

In this tutorial you will learn how to use the switch\-case statement to perform different actions based on different conditions in Golang.

Golang also supports a switch statement similar to that found in other languages such as, Php or Java. Switch statements are an alternative way to express lengthy if else comparisons into more readable code based on the state of a variable.

---

## Golang \- switch Statement

The switch statement is used to select one of many blocks of code to be executed.

Consider the following example, which display a different message for particular day.

package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now()

	switch today.Day() {
	case 5:
		fmt.Println("Today is 5th. Clean your house.")
	case 10:
		fmt.Println("Today is 10th. Buy some wine.")
	case 15:
		fmt.Println("Today is 15th. Visit a doctor.")
	case 25:
		fmt.Println("Today is 25th. Buy some food.")
	case 31:
		fmt.Println("Party tonight.")
	default:
		fmt.Println("No information available for that day.")
	}
}

The default statement is used if no match is found.

---

## Golang \- switch multiple cases Statement

The switch with multiple case line statement is used to select common block of code for many similar cases.

package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now()
	var t int = today.Day()

	switch t {
	case 5, 10, 15:
		fmt.Println("Clean your house.")
	case 25, 26, 27:
		fmt.Println("Buy some food.")
	case 31:
		fmt.Println("Party tonight.")
	default:
		fmt.Println("No information available for that day.")
	}
}

---

## Golang \- switch fallthrough case Statement

The fallthrough keyword used to force the execution flow to fall through the successive case block.

package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now()

	switch today.Day() {
	case 5:
		fmt.Println("Clean your house.")
		fallthrough
	case 10:
		fmt.Println("Buy some wine.")
		fallthrough
	case 15:
		fmt.Println("Visit a doctor.")
		fallthrough
	case 25:
		fmt.Println("Buy some food.")
		fallthrough
	case 31:
		fmt.Println("Party tonight.")
	default:
		fmt.Println("No information available for that day.")
	}
}

Below would be the output on 10th day of month.

C:\\golang\\dns\>go run example.go
Buy some wine.
Visit a doctor.
Buy some food.
Party tonight.

---

## Golang \- swith conditional cases Statement

The case statement can also used with conditional operators.

package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now()

	switch {
	case today.Day() < 5:
		fmt.Println("Clean your house.")
	case today.Day() <\= 10:
		fmt.Println("Buy some wine.")
	case today.Day() \> 15:
		fmt.Println("Visit a doctor.")
	case today.Day() == 25:
		fmt.Println("Buy some food.")
	default:
		fmt.Println("No information available for that day.")
	}
}

---

## Golang \- switch initializer Statement

The switch keyword may be immediately followed by a simple initialization statement where variables, local to the switch code block, may be declared and initialized.

package main

import (
	"fmt"
	"time"
)

func main() {
	switch today := time.Now(); {
	case today.Day() < 5:
		fmt.Println("Clean your house.")
	case today.Day() <\= 10:
		fmt.Println("Buy some wine.")
	case today.Day() \> 15:
		fmt.Println("Visit a doctor.")
	case today.Day() == 25:
		fmt.Println("Buy some food.")
	default:
		fmt.Println("No information available for that day.")
	}
}
