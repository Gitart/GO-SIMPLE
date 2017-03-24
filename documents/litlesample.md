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



## IX. Reflection
### Metaprogramming

```
func TypeOf(i interface{}) Type
func ValueOf(i interface{}) Value
Kind() function and kind constants;
Interface() recovers the (interface) value;
```

## Modifying a value through reflection:

CanSet() -> settability;
v := reflect.ValueOf(x) creates a copy of x;
to change value of x, must pass the address of x: v := reflect.ValueOf(&x)
to make it settable: v = v.Elem()
Reflection on structs:

NumField() gives the number of fields in the struct;
call its methods with Method(n).Call(nil);
only exported fields of a struct are settable;
Reflection with unsafe

```golang
var byteSliceType = reflect.TypeOf(([]byte)(nil))

func AsByteSlice(x interface{}) []byte {
    v := reflect.New(reflect.TypeOf(x))
    h := *(*reflect.SliceHeader)(unsafe.Pointer(v.UnsafeAddr()))
    size := int(v.Type().Elem().Size())
    h.Len *= size
    h.Cap *= size
    return unsafe.Unreflect(byteSliceType, unsafe.Pointer(&h)).([]byte)
    // but in Go1 Unreflect are gone?
}

// or, if you prefer:

func AsByteSlice(x interface{}) []byte {
    v := reflect.New(reflect.TypeOf(x))
    h0 := (*reflect.SliceHeader)(unsafe.Pointer(v.UnsafeAddr()))
    size := int(v.Type().Elem().Size())
    h := reflect.SliceHeader{h0.Data, h0.Len * size, h0.Cap * size}
    return unsafe.Unreflect(byteSliceType, unsafe.Pointer(&h)).([]byte)
}
```


### Turning C arrays into Go slices


func C.GoBytes(cArray unsafe.Pointer, length C.int) []byte
To create a Go slice backed by a C array (without copying the original data), one needs to acquire this length at runtime and use reflect.SliceHeader.

```golang
import "C"
import "unsafe"

        var theCArray *TheCType := C.getTheArray()
        length := C.getTheArrayLength()
        var theGoSlice []TheCType
        sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&theGoSlice)))
        sliceHeader.Cap = length
        sliceHeader.Len = length
        sliceHeader.Data = uintptr(unsafe.Pointer(&theCArray[0]))
```

now theGoSlice is a normal Go slice backed by the C array
It is important to keep in mind that the Go garbage collector will not interact with this data, and that if it is freed from the C side of things, the behavior of any Go code using the slice is nondeterministic.

