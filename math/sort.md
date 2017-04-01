# Первый пример сортировки
```
package main

import "fmt"

func main() {
	var numbers []int = []int{5, 4, 2, 3, 1, 0}
	fmt.Println("Unsorted:", numbers)

	bubbleSort(numbers)
	fmt.Println("Sorted:", numbers)
}

func bubbleSort(numbers []int) {
	var N int = len(numbers)
	var i int
	for i = 0; i < N; i++ {
		sweep(numbers, i)
	}
}

func sweep(numbers []int, prevPasses int) {
	var N int = len(numbers)
	var firstIndex int = 0
	var secondIndex int = 1

	for secondIndex < (N - prevPasses) {
		var firstNumber int = numbers[firstIndex]
		var secondNumber int = numbers[secondIndex]

		if firstNumber > secondNumber {
			numbers[firstIndex] = secondNumber
			numbers[secondIndex] = firstNumber
		}

		firstIndex++
		secondIndex++
	}
}
```

# Ворой пример

```golang
package main

 import (
 	"fmt"
 )

 func bubbleSort(tosort []int) {
 	size := len(tosort)
 	if size < 2 {
 		return
 	}
 	for i := 0; i < size; i++ {
 		for j := size - 1; j >= i+1; j-- {
 			if tosort[j] < tosort[j-1] {
 				tosort[j], tosort[j-1] = tosort[j-1], tosort[j]
 			}
 		}
 	}
 }

 func main() {
 	unsorted := []int{1, 199, 3, 2, 5, 80, 99, 500}
 	
 	fmt.Println("Before : ", unsorted)
 	
 	bubbleSort(unsorted)
 	
 	fmt.Println("After : ", unsorted)
 }
 
``` 
Output :

```
Before : [1 199 3 2 5 80 99 500]
After : [1 2 3 5 80 99 199 500]
```

