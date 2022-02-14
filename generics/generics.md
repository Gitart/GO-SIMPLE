# Начало работы с Go Generics — руководство
 
## Дженерики.

Добавление дженериков к языку Go должно быть одной из самых спорных тем для сообщества Go.

С самого начала мне нравилась ясность и простота Go, которые он предоставляет мне как разработчику. Глядя на сигнатуру функции, я точно знаю, с каким типом я буду работать в теле этой функции, и обычно знаю, на что обращать внимание.

С добавлением дженериков наши кодовые базы становятся немного сложнее. У нас больше нет этой упрощенной эксплицитности, и нам нужно сделать небольшой вывод и покопаться, чтобы действительно узнать, что передается в наши новые функции.

## Обзор

Теперь цель этой статьи не в том, чтобы обсуждать тонкости новейшего дополнения к языку, а в том, чтобы попытаться предоставить вам все необходимое для того, чтобы приступить к работе с дженериками в ваших собственных приложениях Go. .

## Начиная

Прежде чем мы начнем, вам необходимо установить его `go1.18beta1` на свой локальный компьютер. Если вы уже `go` установили, вы можете добиться этого, запустив:

```bash
$ go install golang.org/dl/go1.18beta1@latest
$ go1.18beta1 download

```

После того, как вы успешно запустите эти две команды, вы сможете запустить `go1.18beta1` в своем терминале:

```bash
$ go1.18beta1
go version go1.18beta1 darwin/amd64

```

Отлично, теперь вы можете компилировать и запускать общий код Go!

## Написание универсальных функций

Давайте начнем с рассмотрения того, как мы можем создавать наши собственные универсальные функции в Go. Традиционно вы должны начать с сигнатуры функции, в которой явно указаны типы, которые эта функция ожидает в качестве параметров:

```go
func oldNonGenericFunc(myAge int64) {
    fmt.Println(myAge)
}

```

В новом мире, если мы хотим создать функцию, которая будет принимать, скажем, типы int64 или float64, мы можем изменить сигнатуру нашей функции следующим образом:

main.go

```go
package main

import "fmt"

func newGenericFunc[age int64 | float64](myAge age) {
	fmt.Println(myAge)
}

func main() {
	fmt.Println("Go Generics Tutorial")
	var testAge int64 = 23
	var testAge2 float64 = 24.5

	newGenericFunc(testAge)
	newGenericFunc(testAge2)
}

```

Давайте попробуем запустить это сейчас и посмотрим, что произойдет:

```bash
$ go1.18beta1 run main.go
Go Generics Tutorial
23
24.5

```

Итак, давайте разберем, что мы здесь сделали. Мы фактически создали универсальную функцию под названием `newGenericFunc` .

Затем после имени нашей функции мы открыли квадратные скобки `[]` и указали типы, с которыми мы можем разумно ожидать, что наша функция будет вызываться:

```go
[age int64 | float64]

```

Когда мы определили наши параметры для функции в скобках, мы сказали, что переменная `myAge` может иметь тип `age` , который впоследствии может быть любым `int64` или `float64` типом.

> **Примечание** . Если мы хотим добавить больше типов, мы можем перечислить дополнительные типы, используя `|` разделитель между различными типами.

### Использование любого типа

В приведенном выше примере мы указали универсальную функцию, которая может принимать число типа `int64` или `float64` , однако что, если мы хотим определить функцию, которая может принимать буквально любой тип?

Ну, чтобы добиться этого, мы можем использовать новый встроенный `any` тип следующим образом:

любой.идти

```go
package main

import "fmt"

func newGenericFunc[age any](myAge age) {
	fmt.Println(myAge)
}

func main() {
	fmt.Println("Go Generics Tutorial")
	var testAge int64 = 23
	var testAge2 float64 = 24.5

    var testString string = "Elliot"

	newGenericFunc(testAge)
	newGenericFunc(testAge2)
	newGenericFunc(testString)
}

```

После внесения этих изменений давайте попробуем запустить это сейчас:

```bash
$ go1.18beta1 run any.go
Go Generics Tutorial
23
24.5
Elliot

```

Как видите, компилятор go успешно компилирует и запускает наш новый код. Мы смогли передать любой тип без каких-либо ошибок компилятора.

Теперь вы можете заметить проблему с этим кодом. В приведенном выше примере мы создали новую универсальную функцию, которая принимает `Age` значение и распечатывает его. Затем мы проявили некоторую дерзость и передали `string` значение этой универсальной функции, и, к счастью, на этот раз наша функция обработала этот ввод без каких-либо проблем.

Однако давайте рассмотрим случай, когда это может вызвать проблему. Давайте обновим наш код, чтобы выполнить некоторые дополнительные вычисления для `myAge` параметра. Мы попробуем привести его к , `int` а затем добавим 1 к значению:

generic\_issue.go

```go
package main

import "fmt"

func newGenericFunc[age any](myAge age) {
	val := int(myAge) + 1
	fmt.Println(val)
}

func main() {
	fmt.Println("Go Generics Tutorial")
	var testAge int64 = 23
	var testAge2 float64 = 24.5

    var testString string = "Elliot"

	newGenericFunc(testAge)
	newGenericFunc(testAge2)
    newGenericFunc(testString)
}

```

Теперь, когда мы попытаемся собрать или запустить этот код, мы должны увидеть, что он не компилируется:

```bash
$ go1.18beta1 run generic_issue.go
# command-line-arguments
./generic_issue.go:6:13: cannot convert myAge (variable of type age constrained by any) to type int

```

В этом случае мы не можем пытаться привести тип `any` к `int` . Единственный способ обойти это - более явно указать типы, которые передаются, например:

```go
package main

import "fmt"

func newGenericFunc[age int64 | float64](myAge age) {
	val := int(myAge) + 1
	fmt.Println(val)
}

func main() {
	fmt.Println("Go Generics Tutorial")
	var testAge int64 = 23
	var testAge2 float64 = 24.5

	newGenericFunc(testAge)
	newGenericFunc(testAge2)
}

```

Более подробное описание типов, с которыми мы можем разумно работать, является преимуществом в большинстве сценариев, поскольку это позволяет вам быть более преднамеренным и вдумчивым в отношении того, как вы обрабатываете каждый отдельный тип.

```bash
$ go1.18beta1 run generic_issue.go
Go Generics Tutorial
24
25

```

## Явная передача аргументов типа

В большинстве случаев Go сможет определить тип параметра, который вы передаете в свои универсальные функции. Однако в некоторых сценариях вы можете проявить осмотрительность и указать тип параметра, который вы передаете этим универсальным функциям.

Чтобы быть более явным, мы можем указать тип передаваемого параметра, используя тот же `[]` синтаксис квадратных скобок:

```go
newGenericFunc[int64](testAge)

```

Это будет явно указывать, что `testAge` переменная будет иметь тип `int64` при передаче в this `newGenericFunc` .

## Ограничения типа

Давайте посмотрим, как мы можем изменить наш код и объявить ограничение типа в Go.

В этом примере мы переместим типы, которые может принимать наша универсальная функция, в интерфейс, который мы пометили `Age` . Затем мы использовали это ограничение нового типа в нашем `newGenericFunc` примере:

```go
package main

import "fmt"

type Age interface {
	int64 | int32 | float32 | float64
}

func newGenericFunc[age Age](myAge age) {
	val := int(myAge) + 1
	fmt.Println(val)
}

func main() {
	fmt.Println("Go Generics Tutorial")
	var testAge int64 = 23
	var testAge2 float64 = 24.5

	newGenericFunc(testAge)
	newGenericFunc(testAge2)
}

```

Когда мы запустим это, мы должны увидеть:

```bash
$ go1.18beta1 run type_constraints.go
Go Generics Tutorial
24
25

```

Прелесть этого подхода в том, что мы можем повторно использовать эти же ограничения типов в нашем коде, как и любой другой тип в Go.

## Более сложные ограничения типов

Давайте рассмотрим несколько более сложный вариант использования. Давайте, например, представим, что мы хотим создать `getSalary` функцию, которая будет принимать все, что удовлетворяет заданному ограничению типа. Мы могли бы добиться этого, определив интерфейс, а затем используя его в качестве ограничения типа для нашей универсальной функции:

```go
package main

import "fmt"

type Employee interface {
	PrintSalary()
}

func getSalary[E Employee](e E) {
	e.PrintSalary()
}

type Engineer struct {
	Salary int32
}

func (e Engineer) PrintSalary() {
	fmt.Println(e.Salary)
}

type Manager struct {
	Salary int64
}

func (m Manager) PrintSalary() {
	fmt.Println(m.Salary)
}

func main() {
	fmt.Println("Go Generics Tutorial")
	engineer := Engineer{Salary: 10}
	manager := Manager{Salary: 100}

	getSalary(engineer)
	getSalary(manager)
}

```

В этом случае мы указали, что наша `getSalary` функция имеет ограничение типа, `E` которое должно реализовывать наш `Employee` интерфейс. Затем в нашей основной функции мы определили инженера и менеджера и передали обе эти разные структуры в `getSalary` функцию, хотя они оба разных типов.

```bash
$ go1.18beta1 run complex_type_constraints.go
Go Generics Tutorial
10
100

```

Теперь это интересный пример, он показывает, как мы можем ввести ограничение универсальной функции, чтобы принимать только те типы, которые реализуют этот `PrintSalary` интерфейс, однако того же можно добиться, используя интерфейс непосредственно в сигнатуре функции, например так:

```go
func getSalary(e Employee) {
	e.PrintSalary()
}

```

Это похоже на тот факт, что использование интерфейсов в Go является типом универсального программирования, понимание различий между подходами и преимуществами одного над другим, вероятно, лучше объяснено в официальном посте go.dev под названием « [Почему дженерики](https://go.dev/blog/why-generics) ». .

## Преимущества дженериков

До сих пор мы только что рассмотрели базовый синтаксис, с которым вы столкнетесь при написании универсального кода на Go. Давайте сделаем шаг вперед в этих новых знаниях и посмотрим, где этот код может быть полезен в наших собственных приложениях Go.

Давайте взглянем на стандартную реализацию BubbleSort:

```go
func BubbleSort(input []int) []int {
    n := len(input)
    swapped := true
    for swapped {
        // set swapped to false
        swapped = false
        // iterate through all of the elements in our list
        for i := 0; i < n-1; i++ {
            // if the current element is greater than the next
            // element, swap them
            if input[i] > input[i+1] {
                // log that we are swapping values for posterity
                fmt.Println("Swapping")
                // swap values using Go's tuple assignment
                input[i], input[i+1] = input[i+1], input[i]
                // set swapped to true - this is important
                // if the loop ends and swapped is still equal
                // to false, our algorithm will assume the list is
                // fully sorted.
                swapped = true
            }
        }
    }
}

```

Теперь, в приведенной выше реализации, мы определили, что эта функция BubbleSort должна принимать `int` срез типа. Если бы мы попытались запустить это `int32` , например, с фрагментом типа, мы бы получили ошибку компилятора:

```
cannot use list (variable of type []int32) as type []int in argument to BubbleSort

```

Давайте посмотрим, как мы могли бы написать это, используя дженерики, и открыть ввод, чтобы принимать все типы int и float:

```go
package main

import "fmt"

type Number interface {
	int16 | int32 | int64 | float32 | float64
}

func BubbleSort[N Number](input []N) []N {
	n := len(input)
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < n-1; i++ {
		  	if input[i] > input[i+1] {
				input[i], input[i+1] = input[i+1], input[i]
				swapped = true
		  	}
		}
	}
	return input
}

func main() {
	fmt.Println("Go Generics Tutorial")
	list := []int32{4,3,1,5,}
	list2 := []float64{4.3, 5.2, 10.5, 1.2, 3.2,}
	sorted := BubbleSortGeneric(list)
	fmt.Println(sorted)

	sortedFloats := BubbleSortGeneric(list2)
	fmt.Println(sortedFloats)
}

```

Внеся эти изменения в нашу `BubbleSort` функцию и приняв Type constrained `Number` , мы фактически позволили себе сократить объем кода, который нам приходится писать, если мы хотим поддерживать все типы int и float!

Давайте попробуем запустить это сейчас:

```bash
$ go1.18beta1 run bubblesort.go
Go Generics Tutorial
[1 3 4 5]
[1.2 3.2 4.3 5.2 10.5]

```

Теперь этот единственный пример должен продемонстрировать, насколько мощна эта новая концепция в ваших приложениях Go.

## Вывод

Круто, поэтому в этом уроке мы рассмотрели основы дженериков в Go! Мы рассмотрели, как мы можем определить универсальные функции и сделать классные вещи, такие как использование ограничений типа, чтобы гарантировать, что наши универсальные функции не станут слишком дикими.
