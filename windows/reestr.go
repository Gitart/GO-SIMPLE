package main

import (
    "fmt"

    "golang.org/x/sys/windows/registry"
)

func main() {
    k, err := registry.OpenKey(registry.CURRENT_USER,
        `SOFTWARE\Microsoft\Windows\CurrentVersion\Explorer\Shell Folders`,
        registry.QUERY_VALUE)
    if err != nil {
        panic(err)
    }
    defer k.Close()

    s, _, err := k.GetStringValue("{374DE290-123F-4565-9164-39C4925E467B}")
    if err != nil {
        panic(err)
    }

    fmt.Println(s)
}
