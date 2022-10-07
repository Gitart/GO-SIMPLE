# How to use function from another file golang?

This example aims to demonstrate the various calls of a function in detail. You will learn to create and call a custom package function in main package. You will also be able to call functions of custom packages from the another package using alias.

The following will be the directory structure of our application.

```
├── employee
│   ├── go.mod
│   ├── main.go
│   └── basic
│       └── basic.go
│       └── gross
│       	└── gross.go

```

Go inside the **employee** directory and run the following command to create a go module named **employee**.

go mod init employee

The above command will create a file named go.mod. The following will be the contents of the file.

module employee

go 1.14

---

employee\\main.go

To use a custom package we must import it first. The import path is the name of the module appended by the subdirectory of the package and the package name. In our example the module name is **employee** and the package **basic** is in the **basic** folder directly under **employee** folder. Moreover, the package **gross** is in the **gross** folder which is under **basic** folder.

Hence, the line import "employee/basic" will import the **basic** package, and "employee/basic/gross" will import the **gross** package

package main

import (
	b "employee/basic"
	"employee/basic/gross"

	"fmt"
)

func main() {
	b.Basic = 10000
	fmt.Println(gross.GrossSalary())
}

We are aliasing the basic package as **b**. We called GrossSalary function of **gross** package and assigned value to Basic variable of basic package.

---

employee\\basic\\basic.go

Create a file basic.go inside the **basic** folder. The file inside the basic folder should start with the line package basic as it belongs to the basic package.

package basic

var hra int = 5
var tax int = 2
var Basic int

func Calculation() (allowance int, deduction int) {
	allowance = (Basic \* hra) / 100
	deduction = (Basic \* tax) / 100
	return
}

---

employee\\basic\\gross\\gross.go

Create a file gross.go inside the **gross** folder. The file inside the gross folder should start with the line package gross as it belongs to the gross package.

package gross

import (
	b "employee/basic"
)

func GrossSalary() int {
	a, t := b.Calculation()
	return ((b.Basic + a) - t)
}

The function GrossSalary call the function Calculation of basic package. We are aliasing the basic package as **b**.

---

employee>go run main.go

If you run the program, you will get the following output.

10300
