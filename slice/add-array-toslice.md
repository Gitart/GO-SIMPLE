# Добавление массива к срезу

Добавить в существующий массив к существующему срезу, 

```go
package main
import (
 "fmt"
)
func main() {
 s := []int{1, 2, 3}
 a := [3]int{4, 5, 6}
 ref := a[:]
 fmt.Println("Existing array:\t", ref)
 t := append(s, ref...)
 fmt.Println("New slice:\t", t)
 s = append(s, ref...)
 fmt.Println("Existing slice:\t", s)
 s = append(s, s...)
 fmt.Println("s+s:\t\t", s)
}
```


## Samples
```go
package main

import (
	"fmt"
	"sort"
)

type aStructure struct {
	person string
	height int
	weight int
}

func main() {
	mySlice := make([]aStructure, 0)
	mySlice = append(mySlice, aStructure{"Mihalis", 180, 90})
	mySlice = append(mySlice, aStructure{"Bill", 134, 45})
	mySlice = append(mySlice, aStructure{"Marietta", 155, 45})
	mySlice = append(mySlice, aStructure{"Epifanios", 144, 50})
	mySlice = append(mySlice, aStructure{"Athina", 134, 40})
	fmt.Println("0:", mySlice)

	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].height < mySlice[j].height
	})
	fmt.Println("<:", mySlice)

	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].height > mySlice[j].height
	})

	fmt.Println(">:", mySlice)
}

```
