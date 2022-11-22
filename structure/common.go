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
