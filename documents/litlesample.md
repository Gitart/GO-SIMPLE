## Array

```golang
var arrAge = [5]int{1, 2, 3, 4, 5}
var arrAge = [...]int{1, 2, 3, 4, 5}
var arrKV = [5]string{3: "Andy", 4: "Ken"}
```



## V. Maps

A map is a reference type: var map1 map[keytype]valuetyep
The length does not have to be known at declaration, but can grow dynamically.
Key type can be any type for which the operations == and != are defined.
Value type can be any type, even a func() type:
```golang
mf := map[int]func() int {
    1: func() int { return 10 },
    2: func() int { return 20 },
    5: func() int { return 50 },
}
mf[2]()
```


Arrays, slices and structs cannot be used as key type, but pointers and interface type can.
One way to use structs as a key is to provide them with a Key() or Hash() method.
Map indexing is much faster than a linear search; but still 100x slower than direct array or slice indexing.
Create maps by make(), never new().
Capacity: make(map[key]value, cap)
Test if a key-value exists in a map: _, exists := map1[key1]
Remove: delete(map1, key1)
for-range to get all keys or values.
If the value type of a map is acceptable as a key type, and the map values are unique, inverting can be done easily, via for-range.
