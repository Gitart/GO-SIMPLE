package main

import (
    "fmt"
    "io"
    "strings"
)

type ProxyWrite interface {
    Write(p []byte) (n int, err error)
    SpecialWrite(p []byte) (n int, err error)
}

type Implementer struct {
    counter int
}

func (self Implementer) Write(p []byte) (n int, err error) {
    fmt.Print("Normal write: %v", p)
    return len(p),nil
}

func (self Implementer) SpecialWrite(p []byte) (n int, err error) {
    fmt.Print("Normal write: %v\n", p)
    fmt.Println("And something else")
    self.counter += 1
    return len(p),nil
}


type WriteFunc func(p []byte) (n int, err error)

func (wf WriteFunc) Write(p []byte) (n int, err error) {
    return wf(p)
}

func main() {
    Proxies := make(map[int]ProxyWrite,2)
    Proxies[1] = new(Implementer)
    Proxies[2] = new(Implementer)

    /* runs and uses the Write method normally */
    io.Copy(Proxies[1], strings.NewReader("Hello world"))
    /* gets ./main.go:45: method Proxies[1].SpecialWrite is not an expression, must be called */
    io.Copy(WriteFunc(Proxies[1].SpecialWrite), strings.NewReader("Hello world"))
}
