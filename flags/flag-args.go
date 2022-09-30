package main

import (
    "fmt"
    "flag"
)

func main() {
    flag.Parse()
    fmt.Println(flag.Args()...)
}
