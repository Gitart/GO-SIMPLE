package main

import (
	"fmt"
)

/*

  - Структуры использующие интерфейс должны иметь такие же методы как и интерфейс
  - Структуры могут иметь больше методов чем интерфейс - но не меньше !
  - Названия методов должны сопадать
  - Тип должен сопадать

*/

// Interface
type geometry interface {
	people() string
	system() string
	info() string
}

// My
type my struct {
	id string
}

func (m my) info() string {
	return "MY ИНФО ID " + m.id
}

func (m my) system() string {
	return "My system : " + m.id
}

func (m my) test() string {
	return "My Test : " + m.id
}

func (m my) people() string {
	return "My People : " + m.id
}

func (m my) pe() string {
	return "PE ID" + m.id
}

// REC
type Rec struct {
	id   string
	Name string
}

func (m Rec) info() string {
	return "REC Info : " + m.id + m.Name
}

func (m Rec) system() string {
	return "REC System : " + m.id
}

func (m Rec) people() string {
	return "REC People : " + m.id
}

// Gemetry
func Mm(g geometry) {
	d := g.info()
	fmt.Println("INterface Out : ", d)
}

// Test
func main() {

	m := my{id: "mytest"}
	r := Rec{id: "Reciddd"}

	Mm(m)
	Mm(r)

	var g geometry

	g = Rec{id: "Some other notification"}
	fmt.Println(g.info())
    

	g = Rec{id: "Новости для примера", Name: "Visiul"}
	fmt.Println("***", g.info())
	fmt.Println("***", g.system())
	fmt.Println("***", g.people())

	g = my{id: "MYID"}
	fmt.Println("---", g.info())
	fmt.Println("---", g.system())
	fmt.Println("---", g.people())
}
