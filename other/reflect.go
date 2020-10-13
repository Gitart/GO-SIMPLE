// https://play.golang.org/p/zsJ8TdNhQeG
// https://play.golang.org/p/XaI6wc9vWkr

package main

import (
        "fmt"
        "reflect"
)

func main() {
        tag := reflect.StructTag(`species:"gopher" color:"bl" cor:"bue" tt:"eeer" tn:"222" om:"foo2,foo,omitempty,string"`)
        fmt.Println(tag.Get("color"), tag.Get("species"), tag.Get("om") )
}
