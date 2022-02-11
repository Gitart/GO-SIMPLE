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
