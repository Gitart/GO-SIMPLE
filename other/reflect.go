package main

import (
        "fmt"
        "reflect"
)

func main() {
        tag := reflect.StructTag(`species:"gopher" color:"bl" cor:"bue"`)
        fmt.Println(tag.Get("color"), tag.Get("species"), tag.Get("cor") )
}
