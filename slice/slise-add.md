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


## Append 

```golang

append(s S, x ...T) S  // T is the element type of S

s0 := []int{0, 0}
s1 := append(s0, 2)        // append a single element     s1 == []int{0, 0, 2}
s2 := append(s1, 3, 5, 7)  // append multiple elements    s2 == []int{0, 0, 2, 3, 5, 7}
s3 := append(s2, s0...)    // append a slice              s3 == []int{0, 0, 2, 3, 5, 7, 0, 0}
```

## func append
func append(slice []Type, elems ...Type) []Type The append built-in function appends elements to the end of a slice. If it has sufficient capacity, the destination is resliced to accommodate the new elements. If it does not, a new underlying array will be allocated. Append returns the updated slice. It is therefore necessary to store the result of append, often in the variable holding the slice itself:

```go
slice = append(slice, elem1, elem2)
slice = append(slice, anotherSlice...)
As a special case, it is legal to append a string to a byte slice, like this:

slice = append([]byte("hello "), "world"...)
```


