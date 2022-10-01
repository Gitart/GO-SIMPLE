## Sample visit structure

```go
package main

import (
	"fmt"
)

type visit struct {
	x, y int
}

func main() {
	var visited []visit
	var unique []visit

	uniqueMap := map[visit]int{}

	visited = append(visited, visit{1, 100})
	visited = append(visited, visit{2, 2})
	visited = append(visited, visit{1, 100})
	visited = append(visited, visit{1, 1})

	for _, v := range visited {
		if _, exist := uniqueMap[v]; !exist {
			uniqueMap[v] = 1
			unique = append(unique, v)
		} else {
			uniqueMap[v]++
		}
	}
	fmt.Printf("Uniques: %v\nMaps:%v\n", unique, uniqueMap)
}
```


# Other

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
