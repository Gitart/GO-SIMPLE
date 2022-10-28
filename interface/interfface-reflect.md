# Type interface{}

---

An interface{} is like "object" in Java. It can be any kind of type. Once declared, you must "assert" the type in order to access and use the value.

```
package main

import (
    "fmt"
    "reflect"
)

func doSomething(b interface{}) interface{} {
    fmt.Println("TypeOf(b):", reflect.TypeOf(b))

    c := "Hello " + b.(string)

    fmt.Println(c)

    return 1
}

func main() {
    var a interface{} = "World"

    fmt.Println("TypeOf(a):", reflect.TypeOf(a))

    d := doSomething(a)

    fmt.Println("TypeOf(d):", reflect.TypeOf(d))

    e := 1 + d.(int)
    fmt.Println(e)
}
```
