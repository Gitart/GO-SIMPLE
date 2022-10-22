import (
        "fmt"
)

type A struct {
        a * int
}

func main() {
        var instance A
        value := 14
        instance.a = &value
        fmt.Println(*instance.a) // prints 14
        mutator(instance)
        fmt.Println(*instance.a) // prints 11 o.o
}

func mutator(instance A) {
        *instance.a = 11
        fmt.Println(*instance.a)
}
