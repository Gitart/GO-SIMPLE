
## Slice

```go
package main
import "fmt"

func main() {
	scores := []int{10, 20, 30, 40, 50}
	slice := scores[2:5]
	fmt.Println(slice)
	slice[0] = 999
	slice[1] = 888
	slice[2] = 777
	fmt.Println(scores)
}
```

**Output**
```
[30 40 50]
[10 20 999 888 777]
```


