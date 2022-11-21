
# ⭐ FULL REVIEW INTERFACES OPPORTUNITY

```go 
// 👉 Exaples : 
// 🌏 https://www.golangprograms.com/go-language/interface.html
// 🌏 https://yourbasic.org/golang/find-type-of-object/

package main

import (
	"fmt"
	"time"
)

// Define types and interfaces
type (
	Note interface {
		Info() string
	}

	NoteInfo interface {
		Note
	}
	// Mix
	Mix interface {
		Info() string        // Basic information about identity
		Add() string         // Add string
		GetAll() string      // GetAll Getting all interfaces string
		Structure() []string // Common Method
	}

	// Basic
	Basic interface {
		Info() string        // Basic information about identity
		Add() string         // Add string
		Delete() string      // Add delete
		Update() string      // Add update
		GetAll() string      // Get all string
		Structure() []string // Common Method
		//GetById() string     // Add string
		//GetLast() string     // Add string
		//GetByToday() string  // Add string
		//GetByMonth() string  // Add string
		//GetByPeriod() string // Add string
	}

	Car struct {
		Title string `json:"title"`
		Name  string `json:"name"`
		Id    int    `json:"id"`
		Uid   string `json:"uid"`
	}

	Routing struct {
		Uid   string `json:"uid"`
		Title string `json:"title"`
		From  string `json:"from"`
		To    string `json:"to"`
		Id    int    `json:"id"`
	}

	Note1 string
	Note2 string
	Note3 string
)

// For note Example
func (n Note1) Info() string { return "Info for Note 1" }
func (n Note2) Info() string { return "Info for Note 2" }
func (n Note3) Info() string { return "Info for Note 3" }

func (c *Car) Assign(t, n string) {
	c.Title = t
	c.Name = n
}

func (c *Car) Info() string {
	return "INFO : " + c.Title
}

func (c *Car) Add() string {
	return c.Title
}

func (c *Car) Delete() string {
	return c.Title
}

func (c *Car) Update() string {
	return "UPDATE:" + time.Now().Format("206-001-02-3-5") + c.Title
}

func (c *Car) Structure() []string {
	return []string{"ECU", "Engine", "Air Filters", "Wipers", "Gas Task", "New"}
}

func (c *Car) GetAll() string {
	return c.Title
}

// Routing methods
func (c *Routing) Assign(t string, n int) {
	c.Title = t
	c.Id = n
}

func (c *Routing) Info() string {
	return "INFO : " + c.Title
}

func (c *Routing) GetAll() string {
	return c.Title
}

func (c *Routing) Add() string {
	return c.Title
}

func (c *Routing) Delete() string {
	return c.Title
}

func (c *Routing) Update() string {
	return c.Title
}
func (c *Routing) Structure() []string {
	return []string{"London", "Dnepropetrovsk", "Paris", "Area", "Libain", "Rome"}
}

func PrintInterface(b Basic) {
	fmt.Println(b)
}

func WorkerInterface(b Basic) {
	fmt.Println(b.Update())
}

// Main
func main() {
	// Empty types
	EmptyTypes()

	// Interface Embedding
	// NoteEmbedding()

	// Interface Accepting Address of the Variable
	// AssignType()

	// SecondUsage()

	// DefineType()
}

// Print type
func printType(i interface{}) {
	// Get Type interface
	xType := fmt.Sprintf("%T", i)
	fmt.Println(xType)

	// Get Type interface
	switch v := i.(type) {
	case string:
		fmt.Println("string:", v)
	case []string:
		fmt.Println("[]string:", v)
	case int:
		fmt.Println("int:", v)
	case map[string]int:
		fmt.Println("Map string int:", v)
	case []int:
		fmt.Println("[]int:", v)
	case float64:
		fmt.Println("float64:", v)
	default:
		fmt.Println("unknown")
	}
	// Show interface
	fmt.Println(i)
}

// Empty types
func EmptyTypes() {
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

// Polymorphism
// Interface Embedding
func NoteEmbedding() {
	p := new(Note1)
	h := new(Note2)
	o := new(Note3) // var o Note3

	// This construction also will be worked
	// This design will also work
	// notes := [...]Note{p, h, o}

	// Embedding interface
	notes := [...]NoteInfo{p, h, o}
	for i := range notes {
		fmt.Println("NOTE INFO: ", notes[i].Info())
	}
}

// Interface Accepting Address of the Variable
func AssignType() {
	c := Car{}
	p := Routing{}

	p.Assign("Assign path", 212783)
	c.Assign("Assign: Toyota", "Assign: Pererro Alex")

	fmt.Println("Call interface")
	var b Basic
	b = &c
	fmt.Println(b.Info())
	b = &p
	fmt.Println(p.Info())

}

// Define Type that Satisfies Multiple Interfaces
func DefineType() {
	var b Basic = &Car{Title: "Toyota"}
	fmt.Printf("Car Info : %s \n", b.Info())

	var p Basic = &Routing{Title: "Kiev-Lvov"}
	fmt.Printf("Rout Info : %s \n", p.Info())

	o := p.(Basic)
	fmt.Println(o)
}

// Interfaces with common Method
func SecondUsage() {
	//var c Car
	c := Car{Title: "World Top Brand"}

	//var r Routing
	r := Routing{Title: "Software Developer"}

	for i, j := range c.Structure() {
		fmt.Printf("%-15s <=====> %15s\n", j, r.Structure()[i])
	}
}

// Basic Usage
func BasicUsage() {
	var Dat Basic

	c := Car{
		Title: "Test car title",
		Name:  "Test name",
		Id:    123,
	}

	r := Routing{
		To:    "Kiyev",
		From:  "Alexandriya",
		Id:    123839,
		Title: "Basic routing to route by country",
	}

	print("#1 First example to using :")
	WorkerInterface(&c)
	PrintInterface(&c)
	PrintInterface(&r)

	print("#2 Secondary example to using : ")
	Dat = &c
	Res := Dat.Update()
	fmt.Println(Res)
}

```
