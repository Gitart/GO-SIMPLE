package main

import "os"

func main(){
    f, err := os.OpenFile("test.txt", os.O_APPEND|os.O_WRONLY, 0600)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    if _, err = f.WriteString("My name is Mike"); err != nil {
    panic(err)
    }
}
