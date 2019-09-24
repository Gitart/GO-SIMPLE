## Интерфейс Голанг
#### Интерфейс - это абстрактный тип.

**Интерфейс** описывает все методы набора методов и предоставляет подписи для каждого метода.

Для создания интерфейса используйте ключевое слово интерфейса , за которым следуют фигурные скобки, содержащие список имен методов, а также любые параметры или возвращаемые значения, которые должны иметь методы.

```golang
// Declare an Interface Type and methods does not have a body
type Employee interface {
	PrintName() string                // Method with string return type
	PrintAddress(id int)              // Method with int parameter
	PrintSalary(b int, t int) float64 // Method with parameters and return type
}
```

Интерфейсы действуют как план для наборов методов, они должны быть реализованы перед использованием. Тип, который удовлетворяет интерфейсу, как говорят, реализует это.

Определите тип, который удовлетворяет интерфейсу
Определяет тип интерфейса с именем Employee с помощью двух методов. Затем он определяет тип с именем Emp, который удовлетворяет Employee .

Мы определяем все методы на Emp, которые нужны для удовлетворения сотрудника.

```golang
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

Если у типа есть все методы, объявленные в интерфейсе, то никаких дополнительных объявлений явно не требуется, чтобы сказать, что Emp удовлетворяет Employee .

Объявляет переменную e1 с Employee в качестве ее типа, затем создает значение Emp и присваивает его e1 .

Определите тип, который удовлетворяет нескольким интерфейсам
Интерфейсы позволяют любому пользовательскому типу одновременно удовлетворять нескольким типам интерфейсов.

Используя утверждение типа, вы можете получить обратно значение конкретного типа и вызвать для него методы, которые определены в другом интерфейсе, но не являются частью интерфейса, удовлетворяющего требованиям.

```golang
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

Когда пользовательский тип реализует набор методов, объявленных типом интерфейса, значения пользовательского типа могут быть назначены значениям типа интерфейса. Это назначение сохраняет значение пользовательского типа в значение интерфейса. Когда вызов метода выполняется для значения интерфейса, будет выполнен эквивалентный метод для сохраненного пользовательского значения. Поскольку любой пользовательский тип может реализовывать любой интерфейс, вызовы методов для значения интерфейса имеют полиморфный характер. Пользовательский тип в этом отношении часто называют конкретным типом .

### Интерфейсы с общим методом
Два или более интерфейса могут иметь один или несколько общих методов в списке наборов методов. Здесь, Структура является общим методом между двумя интерфейсами Транспортное средство и Человек .

```golang
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

Когда приведенный выше код компилируется и выполняется, он дает следующий результат -

C:\Golang>go run main.go
ECU             <=====>           Brain
Engine          <=====>           Heart
Air Filters     <=====>            Nose
Wipers          <=====>       Eyelashes
Gas Task        <=====>         Stomach

### Интерфейс, принимающий адрес переменной
В печати () методы принимают указатель приемника. Следовательно, интерфейс также должен принимать указатель получателя.

Если метод принимает значение типа, то интерфейс должен получить значение типа; если у метода есть получатель указателя, то интерфейс должен получить адрес переменной соответствующего типа.

```golang
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

### Пустой тип интерфейса
Интерфейс типа {} известен как пустой интерфейс и используется для принятия значений любого типа. Пустой интерфейс не имеет методов, которые необходимы для его удовлетворения, и поэтому каждый тип удовлетворяет этому.

```golang
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

ManyType переменная объявляется быть типа интерфейса {} и может быть присвоены значения различных типов. Функция printType () принимает параметр интерфейса типа {} , поэтому эта функция может принимать значения любого допустимого типа.

Когда приведенный выше код компилируется и выполняется, он дает следующий результат -

go run main.go
100
200.5
Germany
Go programming language
[india japan canada australia russia]
map[Mark:10 Sandy:20]
[Japan Australia Germany]

### Полиморфизм
**Полиморфизм** - это способность писать код, который может иметь различное поведение при реализации типов.

У нас есть объявление структур с именами Pentagon, Hexagon, Octagon и Decagon с реализацией интерфейса Geometry .

```golang
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

Когда приведенный выше код компилируется и выполняется, он дает следующий результат -

C:\Golang>go run main.go
25
30
40
50
У нас есть наши полиморфные функции Edges, которые принимают значения, которые реализуют интерфейс Geometry . Используя полиморфный подход, метод, созданный здесь, Параметр используется каждым конкретным значением типа, которое передается.

### Встраивание интерфейса
Интерфейсы могут включать другие интерфейсы, это поведение является аспектом интерфейсного полиморфизма, который известен как специальный полиморфизм .

```golang
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

Когда один тип внедряется в другой тип, методы встроенного типа доступны для типа внедрения. Метод или методы встроенного интерфейса доступны для интерфейса встраивания.
