# Declare and initialize a map


You define a map using the keyword `map`, followed by the types of its keys and its values:

```golang
    // Keys are ints, values are ints.
    var m1 map[int]int // initialized to nil
    
    // Keys are strings, values are ints.
    var m2 map[string]int // initialized to nil
```

Maps are reference types, and once defined they have a [zero value](https://www.notion.so/kjkpublic/Zero-value-of-a-map-f7fd36d62c1b443884f83b73c72fc5ac) of `nil`. Writes to nil maps will cause a [panic](https://www.notion.so/kjkpublic/Panic-and-recover-dda02e368b3048a5a69b81c56bfa6dbb) and reads will always return the zero value.


To initialize a map, use the `[make](https://golang.org/pkg/builtin/#make)` function:

```golang
    m := make(map[string]int)
```

With the two-parameter form of `make`, it’s possible to specify an initial entry capacity for the map, overriding the default capacity:

```golang
    m := make(map[string]int, 30)
```


Alternatively, you can declare a map, initializing it to its zero value, and then assign a literal value to it later, which helps if you marshal the struct into json thereby producing an empty map on return.

```golang
    m := make(map[string]int, 0)
```    

You can also make a map and set its initial value with curly brackets (`{}`).

```golang
    var m map[string]int = map[string]int{"Foo": 20, "Bar": 30}
    fmt.Println(m["Foo"]) // outputs 20
```    

All the following statements result in the variable being bound to the same value.

   
```golang
    // Declare, initializing to zero value, then assign a literal value.
    var m map[string]int
    m = map[string]int{}
    
    // Declare and initialize via literal value.
    var m = map[string]int{}
    
    // Declare via short variable declaration and initialize with a literal value.
    m := map[string]int{}
```


We can also use a *map literal* to [create a new map](https://www.notion.so/kjkpublic/Create-a-map-2b831bac5afc414493cff5e06e8e4460) with some initial key/value pairs.

The key type can be any [comparable](http://golang.org/ref/spec#Comparison_operators) type; notably, this excludes functions, maps, and slices. The value type can be any type, including custom types or `interface{}`.

```golang
    type Person struct {
        FirstName string
        LastName  string
    }
    
    // Declare via short variable declaration and initialize with make.
    m := make(map[string]Person)
    
    // Declare, initializing to zero value, then assign a literal value.
    var m map[string]Person
    m = map[string]Person{}
    
    // Declare and initialize via literal value.
    var m = map[string]Person{}
    
    // Declare via short variable declaration and initialize with a literal value.
    m := map[string]Person{}
```    
