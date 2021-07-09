## Add dots after the second slice:

```go
append([]int{1,2}, []int{3,4}...)
This is just like any other variadic function.

func foo(is ...int) {
    for i := 0; i < len(is); i++ {
        fmt.Println(is[i])
    }
}

func main() {
    foo([]int{9,8,7,6,5}...)
}
```
