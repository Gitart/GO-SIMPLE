# How to append anything (element, slice or string) to a slice

yourbasic.org/golang

![](https://yourbasic.org/golang/toddler-fiat-with-trailer.jpg)

*   [Append function basics](https://yourbasic.org/golang/append-explained/#append-function-basics)
*   [Append one slice to another](https://yourbasic.org/golang/append-explained/#append-one-slice-to-another)
*   [Append string to byte slice](https://yourbasic.org/golang/append-explained/#append-string-to-byte-slice)
*   [Performance](https://yourbasic.org/golang/append-explained/#performance)

## Append function basics

With the built-in [append function](https://golang.org/ref/spec#Appending_and_copying_slices) you can use a slice as a [dynamic array](https://yourbasic.org/algorithms/time-complexity-arrays/). The function appends any number of elements to the end of a [slice](https://yourbasic.org/golang/slices-explained/):

*   if there is enough capacity, the underlying array is reused;
*   if not, a new underlying array is allocated and the data is copied over.

Append **returns the updated slice**. Therefore you need to store the result of an append, often in the variable holding the slice itself:

```
a := []int{1, 2}
a = append(a, 3, 4) // a == [1 2 3 4]
```

In particular, it’s perfectly fine to **append to an empty slice**:

```
a := []int{}
a = append(a, 3, 4) // a == [3 4]
```

> **Warning:** See [Why doesn’t append work every time?](https://yourbasic.org/golang/gotcha-append/) for an example of what can happen if you forget that `append` may reuse the underlying array.

## Append one slice to another

You can **concatenate two slices** using the [three dots notation](https://yourbasic.org/golang/variadic-function/):

```
a := []int{1, 2}
b := []int{11, 22}
a = append(a, b...) // a == [1 2 11 22]
```

The `...` unpacks `b`. Without the dots, the code would attempt to append the slice as a whole, which is invalid.

The result does not depend on whether the **arguments overlap**:

```
a := []int{1, 2}
a = append(a, a...) // a == [1 2 1 2]
```

## Append string to byte slice

As a special case, it’s legal to append a string to a byte slice:

```
slice := append([]byte("Hello "), "world!"...)
```
