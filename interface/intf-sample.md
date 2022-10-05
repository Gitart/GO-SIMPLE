
```go
package main

import "fmt"

type Service struct {
	description    string
	durationMonths int
	monthlyFee     float64
	features       []string
}

// var expense Expense = product

func (s Service) getName() string {
	return s.description
}
func (s Service) getCost(recur bool) float64 {
	if recur {
		return s.monthlyFee * float64(s.durationMonths)
	}
	return s.monthlyFee
}

type Expense interface {
	getName() string
	getCost(annual bool) float64
}

type Product struct {
	name, category string
	price          float64
}

func (p Product) getName() string {
	return p.name
}
func (p Product) getCost(_ bool) float64 {
	return p.price
}

func main() {
	

	expenses := []Expense{
		Service{"Boat Cover", 12, 89.50, []string{}},
		Service{"Paddle Protect", 12, 8, []string{}},
		&Product{"Kayak", "Watersports", 275},
	}
	for _, expense := range expenses {
		if s, ok := expense.(Service); ok {
			fmt.Println("Service:", s.description, "Price:",
				s.monthlyFee*float64(s.durationMonths))
		} else {
			fmt.Println("Expense:", expense.getName(),
				"Cost:", expense.getCost(true))
		}
	}
}
```
