```jsx
package main

import "fmt"

func main() {
	sales := [][]int{
		{100, 200},
		{300},
		{400, 500},
	}

	for _, x := range sales {
		for _, y := range x {
			fmt.Println(y)
		}
	}
}
```

```markup
100
200
300
400
500
```
