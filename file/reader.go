package main

import (
    "fmt"
    "os"
    "time"
)

func main() {
    f, err := os.OpenFile("/tmp/file.txt", os.O_RDWR, 0666)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    i := 0
    for {
        n, err := fmt.Fscanln(f, &i)
        if n == 1 {
            fmt.Println(i)
        }
        if err != nil {
            fmt.Println(err)
            return
        }
        time.Sleep(500 * time.Millisecond)
    }
}
