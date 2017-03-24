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

## VII. Methods

* Method is a function that acts an variable of a certain type, called receiver.
* The receiver can be almost anything: even a function type or alias type for int, bool, string or array.
* The receiver cannot be an interface type.
* The receiver cannot be a pointer type, but it can be a pointer to any of the allowed types.
An alias of a certain type *doesn’t have the methods defined on that type.
If the method does not need to use the value recv, discard it by _.
A method and the type on which it acts must be defined in the same package. Use an alias type, or struct instead.
The receiver must have an explicit name, and this name must be used in the method.
The receiver type is called the (receiver) base type; it must be declared within the same package as all of its methods
If for a type T a method Meth() exists on *T and t is a variable of type T, then t.Meth() is automatically translated to (&t).Meth().
Pointer and value methods can both be called on pointer or non-pointer values.
It should not be possible that the fields of an object can be changed by 2 or more different threads at the same time. Use methods of the package sync.
### When an anonymous type is embedded in a struct, the visible methods of that type are inherited by the outer type:

```golang
type Point struct {
    x, y float64
}

func (p *Point) Abs() float64 {
    return math.Sqrt(p.x*p.x + p.y*p.y)
}

type NamedPoint struct {
    Point
    name string
}

func main() {
    n := &NamedPoint{Point{3,4}, "Pythagoras"}
    fmt.Println(n.Abs())
}
```

A method in the embedding type with the same name as a method in an embedded type overrides this.
```golang
func (n *NamedPoint) Abs() float64 {
    return n.Point.Abs() * 100
}
```

Embedding multiple anonymous types -> multiple inheritance.
Structs embedding structs from the same package have all access to one another’s fields and methods.
Embed functionality in a type:

Aggregation/composition: include a named field of the type of the wanted functionality
Embedding: anonymously embed the type of the wanted functionality
String() method:

do not make mistake of defining String() in terms of itself -> infinite recursion.
Summarzed: in Go types are basically classes. Go does not know inheritance like class oriented OO languages.

More OO capabilities: goop provides Go with JavaScript-style objects but supports multiple inheritance and type-dependent dispatch.
Suppose special action needs to be taken right before an object obj is removed from memory (gc), like writing to a log-file, achieved by:

```golang
runtime.SetFinalizer(obj, func(obj *typeObj))
```

SetFinalizer does not execute when the program comes to an normal end or when an error occurs, before the object was chosen by the gc process to be removed.

## VIII. Interface


Polymorphism

An interface defines abstract method set, but cannot contain variables;
Interface internal:

  +------------+------------+
  |            |   method   | 
  |  receiver  |   table    |
  |            |   pointer  |
  +------------+------------+
Thus, pointers to interface values are illegal.
Table of method pointers is built through runtime reflection capability.
Multiple types can implement the same interface;

A type that implements an interface can also have other methods;
A type cam implements many interface;
An interface type can contain a reference to an instance of any of the types that implement the interface: dynamic type;
Interface embedding interface(s): enumerating the methods;
type assertions:
```golang
v := varI.(T) // unchecked type assertion
if v, ok := varI.(T); ok {...} // checked type assertion
```

varI must be an interface variable.
testing if a value implements an interface.

type switch:

```golang
switch t := var.(type) { // t can be omited
case a:
    //
case b:
    //
default:
    //
}
```

An interface value can also be assigned to another interface value, as long as the underlying value implements the necessary methods.

Summarized:

* Pointer methods can be called with pointers;
* Value methods can be called with values;
* Value-receiver methods can be called with pointer values because they can be dereferenced first.
* Pointer-receiver methods cannot be called with values, because the value stored inside an interface has no address.

### Empty Interface

The empty interface has no methods: type Any interface{}.
Any variable, any type implements it.
It can through assignment receive a variable of any type.
Each interface{} takes up 2 words in memory: one word for the type, the other word for either value or pointer to value.
Can perform templates role in container.
copying a data-slice in a slice of interface{} dose not work:

```golang
var dataslice []myType = getSlice()
var interfaceSlice []interface{} = dataslice
```

// must be done explicitly with for-range, e.g.
Overloading functions: 
```golang
func DoSomething(f int, a ...interface{}) (n int, errno error)
```


