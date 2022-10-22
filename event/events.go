package main

import (
        "fmt"
)

type A struct {
        a *int
        c  int
        d *int
}

func main() {
        var instance A
        value := 100
        val   := 100
        
        instance.a = &value
        instance.c = 200
        instance.d = &val
        fmt.Println("A:", *instance.a, "C:", instance.c, "D:", *instance.d) // prints 14
        
        mutator(instance)
        fmt.Println("A:", *instance.a, "C:", instance.c, "D:", *instance.d)

        mutator2(instance)

        // be changed 
        // *instance.a = 1005
        //  instance.c = 2053  // be changed 

        fmt.Println("A:", *instance.a, "C:", instance.c, "D:", *instance.d)
}


func mutator(instance A) {
        *instance.a = 2000
         instance.c = 3000   // Not be  changed
        *instance.d = 7000

        // fmt.Println(*instance.a)
        // fmt.Println(instance.c)

}

func mutator2(instance A) {
        *instance.a = 12000
         instance.c = 13000
        *instance.d = 9000         

        // fmt.Println(*instance.a)
        // fmt.Println(instance.c)

}
