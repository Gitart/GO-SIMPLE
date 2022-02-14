# GENERICS 

// https://tutorialedge.net/golang/getting-starting-with-go-generics/

```golang
package main
import "fmt"

type Number interface {
	int16 | int32 | int64 | float32 | float64 
}

func BubbleSortGeneric[N Number](input []N) []N {
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
	list   := []int32{4,3,1,5,}
	list2  := []float64{4.3, 5.2, 10.5, 1.2, 3.2,}
	sorted := BubbleSortGeneric(list)
	fmt.Println(sorted)

	sortedFloats := BubbleSortGeneric(list2)
	fmt.Println(sortedFloats)
}
```
## variant 2
```golang
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

## variant 2
```golang
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


## Varian 3
```golang
package main
import "fmt"

/* Type Constraints
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

### varian 1

```golang
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
*/


## variant 5

```golang
package main
import "fmt"

func newGenericFunc[age any](myAge age) {
	fmt.Println(myAge)
}


func mains() {
	fmt.Println("Go Generics Tutorial")
	var testAge int64 = 23
	var testAge2 float64 = 24.5
    var testString string = "Elliot"

	newGenericFunc(testAge)
	newGenericFunc(testAge2)
	newGenericFunc(testString)
}
```
