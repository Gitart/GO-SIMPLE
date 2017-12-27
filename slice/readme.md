## The append method from § 7.5 is very versatile and can be used for all kinds of manipulations:
So to represent a resizable sequence of elements use a slice and the append-operation

1) Append a slice b to an existing slice     a: a = append(a, b...)
2) Copy a slice a to a new slice              b: b = make([]T, len(a)) copy(b, a)
3) Delete item at index                       i: a = append(a[:i], a[i+1:]...)
4) Cut from index i till j out of slice       a: a = append(a[:i], a[j:]...)
5) Extend slice a with a new slice of length  j: a = append(a, make([]T, j)...)
6) Insert item x at index                     i: a = append(a[:i], append([]T{x},a[i:]...)...)
7) Insert a new slice of length j at index    i: a = append(a[:i], append(make([]T,j), a[i:]...)...)
8) Insert an existing slice b at index        i: a = append(a[:i], append(b,a[i:]...)...)
9) Pop highest element from stack:            x, a = a[len(a)-1], a[:len(a)-1]
10) Push an element x on a stack:             a = append(a, x)
