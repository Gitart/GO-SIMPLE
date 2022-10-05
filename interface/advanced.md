# Golang Interface
![image](https://user-images.githubusercontent.com/3950155/194111296-b0337c30-d9a2-4c40-91c4-9c4ee4f5d7b9.png)

![image](https://user-images.githubusercontent.com/3950155/194110230-fa8a5bc7-c14f-4cb0-9a2d-e7d552831cb5.png)

---

An Interface is an abstract type.

Interface describes all the methods of a method set and provides the signatures for each method.

To create interface use interface keyword, followed by curly braces containing a list of method names, along with any parameters or return values the methods are expected to have.

```jsx
// Declare an Interface Type and methods does not have a body
type Employee interface {
	PrintName() string                // Method with string return type
	PrintAddress(id int)              // Method with int parameter
	PrintSalary(b int, t int) float64 // Method with parameters and return type
}
```

An interfaces act as a blueprint for method sets, they must be implemented before being used. Type that satisfies an interface is said to implement it.

---

## Define Type that Satisfies an Interface

Defines an interface type named **Employee** with two methods. Then it defines a type named Emp that satisfies **Employee**.

We define all the methods on Emp that it needs to satisfy **Employee**

```jsx
package main

import "fmt"

// Employee is an interface for printing employee details
type Employee interface {
	PrintName(name string)
	PrintSalary(basic int, tax int) int
}

// Emp user-defined type
type Emp int

// PrintName method to print employee name
func (e Emp) PrintName(name string) {
	fmt.Println("Employee Id:\t", e)
	fmt.Println("Employee Name:\t", name)
}

// PrintSalary method to calculate employee salary
func (e Emp) PrintSalary(basic int, tax int) int {
	var salary = (basic * tax) / 100
	return basic - salary
}

func main() {
	var e1 Employee
	e1 = Emp(1)
	e1.PrintName("John Doe")
	fmt.Println("Employee Salary:", e1.PrintSalary(25000, 5))
}
```

If a type has all the methods declared in an interface, then no further declarations needed explicitly to say that Emp satisfies **Employee**.

Declares an *e1* variable with **Employee** as its type, then creates a Emp value and assigns it to *e1*.

---

## Define Type that Satisfies Multiple Interfaces

Interfaces allows any user\-defined type to satisfy multiple interface types at once.

Using **Type Assertion** you can get a value of a concrete type back and you can call methods on it that are defined on other interface, but aren't part of the interface satisfying.

```jsx
package main

import "fmt"

type Polygons interface {
	Perimeter()
}

type Object interface {
	NumberOfSide()
}

type Pentagon int

func (p Pentagon) Perimeter() {
	fmt.Println("Perimeter of Pentagon", 5*p)
}

func (p Pentagon) NumberOfSide() {
	fmt.Println("Pentagon has 5 sides")
}

func main() {
	var p Polygons = Pentagon(50)
	p.Perimeter()
	var o Pentagon = p.(Pentagon)
	o.NumberOfSide()

	var obj Object = Pentagon(50)
	obj.NumberOfSide()
	var pent Pentagon = obj.(Pentagon)
	pent.Perimeter()
}
```

When a user\-defined type implements the set of methods declared by an interface type, values of the user\-defined type can be assigned to values of the interface type. This assignment stores the value of the user\-defined type into the interface value. When a method call is made against an interface value, the equivalent method for the stored user\-defined value will be executed. Since any user\-defined type can implement any interface, method calls against an interface value are polymorphic in nature. The user\-defined type in this relationship is often referred as **concrete type**.

---

## Interfaces with common Method

Two or more interfaces can have one or more common method in list of method sets. Here, **Structure** is a common method between two interfaces Vehicle and Human.

```jsx
package main

import "fmt"

type Vehicle interface {
	Structure() []string // Common Method
	Speed() string
}

type Human interface {
	Structure() []string // Common Method
	Performance() string
}

type Car string

func (c Car) Structure() []string {
	var parts = []string{"ECU", "Engine", "Air Filters", "Wipers", "Gas Task"}
	return parts
}

func (c Car) Speed() string {
	return "200 Km/Hrs"
}

type Man string

func (m Man) Structure() []string {
	var parts = []string{"Brain", "Heart", "Nose", "Eyelashes", "Stomach"}
	return parts
}

func (m Man) Performance() string {
	return "8 Hrs/Day"
}

func main() {
	var bmw Vehicle
	bmw = Car("World Top Brand")

	var labour Human
	labour = Man("Software Developer")

	for i, j := range bmw.Structure() {
		fmt.Printf("%-15s <=====> %15s\n", j, labour.Structure()[i])
	}
}
```

When the above code is compiled and executed, it produces the following result −

```markup
C:\Golang>go run main.go
ECU             <=====>           Brain
Engine          <=====>           Heart
Air Filters     <=====>            Nose
Wipers          <=====>       Eyelashes
Gas Task        <=====>         Stomach
```

---

## Interface Accepting Address of the Variable

The Print() methods accept a receiver pointer. Hence, the interface must also accept a receiver pointer.

*If a method accepts a type value, then the interface must receive a type value; if a method has a pointer receiver, then the interface must receive the address of the variable of the respective type.*

```jsx
package main

import "fmt"

type Book struct {
	author, title string
}

type Magazine struct {
	title string
	issue int
}

func (b *Book) Assign(n, t string) {
	b.author = n
	b.title = t
}
func (b *Book) Print() {
	fmt.Printf("Author: %s, Title: %s\n", b.author, b.title)
}

func (m *Magazine) Assign(t string, i int) {
	m.title = t
	m.issue = i
}
func (m *Magazine) Print() {
	fmt.Printf("Title: %s, Issue: %d\n", m.title, m.issue)
}

type Printer interface {
	Print()
}

func main() {
	var b Book                                 // Declare instance of Book
	var m Magazine                             // Declare instance of Magazine
	b.Assign("Jack Rabbit", "Book of Rabbits") // Assign values to b via method
	m.Assign("Rabbit Weekly", 26)              // Assign values to m via method

	var i Printer // Declare variable of interface type
	fmt.Println("Call interface")
	i = &b    // Method has pointer receiver, interface does not
	i.Print() // Show book values via the interface
	i = &m    // Magazine also satisfies shower interface
	i.Print() // Show magazine values via the interface
}
```

---

## Empty Interface Type

The type interface{} is known as the **empty interface**, and it is used to accept values of any type. The empty interface doesn't have any methods that are required to satisfy it, and so every type satisfies it.

```jsx
package main

import "fmt"

func printType(i interface{}) {
	fmt.Println(i)
}

func main() {
	var manyType interface{}
	manyType = 100
	fmt.Println(manyType)

	manyType = 200.50
	fmt.Println(manyType)

	manyType = "Germany"
	fmt.Println(manyType)

	printType("Go programming language")
	var countries = []string{"india", "japan", "canada", "australia", "russia"}
	printType(countries)

	var employee = map[string]int{"Mark": 10, "Sandy": 20}
	printType(employee)

	country := [3]string{"Japan", "Australia", "Germany"}
	printType(country)
}
```

The manyType variable is declared to be of the type **interface{}** and it is able to be assigned values of different types. The printType() function takes a parameter of the type **interface{}**, hence this function can take the values of any valid type.

When the above code is compiled and executed, it produces the following result −

```markup
go run main.go
100
200.5
Germany
Go programming language
[india japan canada australia russia]
map[Mark:10 Sandy:20]
[Japan Australia Germany]
```

---

## Polymorphism

Polymorphism is the ability to write code that can take on different behavior through the implementation of types.

We have the declaration of a structs named Pentagon, Hexagon, Octagon and Decagon with the implementation of the **Geometry** interface.

```jsx
package main

import (
	"fmt"
)

// Geometry is an interface that defines Geometrical Calculation
type Geometry interface {
	Edges() int
}

// Pentagon defines a geometrical object
type Pentagon struct{}

// Hexagon defines a geometrical object
type Hexagon struct{}

// Octagon defines a geometrical object
type Octagon struct{}

// Decagon defines a geometrical object
type Decagon struct{}

// Edges implements the Geometry interface
func (p Pentagon) Edges() int { return 5 }

// Edges implements the Geometry interface
func (h Hexagon) Edges() int { return 6 }

// Edges implements the Geometry interface
func (o Octagon) Edges() int { return 8 }

// Edges implements the Geometry interface
func (d Decagon) Edges() int { return 10 }

// Parameter calculate parameter of object
func Parameter(geo Geometry, value int) int {
	num := geo.Edges()
	calculation := num * value
	return calculation
}

// main is the entry point for the application.
func main() {
	p := new(Pentagon)
	h := new(Hexagon)
	o := new(Octagon)
	d := new(Decagon)

	g := [...]Geometry{p, h, o, d}

	for _, i := range g {
		fmt.Println(Parameter(i, 5))
	}
}

```

When the above code is compiled and executed, it produces the following result −

```markup
C:\Golang>go run main.go
25
30
40
50
```

We have our polymorphic Edges functions that accepts values that implement the **Geometry** interface. Using polymorphic approach the method created here Parameter is used by each concrete type value that's passed in.

---

## Interface Embedding

Interfaces may embed other interfaces, this behavior is an aspect of interface polymorphism which is known as **ad hoc polymorphism**.

```jsx
package main

import "fmt"

type Geometry interface {
	Edges() int
}

type Polygons interface {
	Geometry // Interface embedding another interface
}

type Pentagon int
type Hexagon int
type Octagon int
type Decagon int

func (p Pentagon) Edges() int { return 5 }
func (h Hexagon) Edges() int  { return 6 }
func (o Octagon) Edges() int  { return 8 }
func (d Decagon) Edges() int  { return 10 }

func main() {
	p := new(Pentagon)
	h := new(Hexagon)
	o := new(Octagon)
	d := new(Decagon)

	polygons := [...]Polygons{p, h, o, d}
	for i := range polygons {
		fmt.Println(polygons[i].Edges())
	}
}
```

When one type is embedded into another type, the methods of the embedded type are available to the embedding type. The method or methods of the embedded interface are accessible to the embedding interface.
