## Как указатель & и * и ** работает в Голанге?
Программа ниже - это карри из указателей. Значение intVar равно значению ** pointerToPointerVar.

```golang
package main

import "fmt"

func main() {
    var intVar int
    var pointerVar *int
    var pointerToPointerVar **int

    intVar = 100
    pointerVar = &intVar
    pointerToPointerVar = &pointerVar
    
    fmt.Println("\n")
    fmt.Println("intVar:\t\t\t", intVar)
    fmt.Println("pointerVar:\t\t", pointerVar)
    fmt.Println("pointerToPointerVar:\t", pointerToPointerVar)
    
    
    fmt.Println("\n")
    fmt.Println("&intVar:\t\t", &intVar)
    fmt.Println("&pointerVar:\t\t", &pointerVar)
    fmt.Println("&pointerToPointerVar:\t", &pointerToPointerVar)

    fmt.Println("\n")
    fmt.Println("*pointerVar:\t\t", *pointerVar) 
    fmt.Println("*pointerToPointerVar:\t", *pointerToPointerVar)
    fmt.Println("**pointerToPointerVar:\t", **pointerToPointerVar)
}
```

