package main

import "fmt"

type Chain struct {
}

type ChainAndError struct {
    *Chain
    error
}

func (v *Chain)funA() ChainAndError {
    fmt.Println("A")
    return ChainAndError{v, nil}
}

func (v *Chain)funB() ChainAndError {
    fmt.Println("B")
    return ChainAndError{v, nil}
}

func (v *Chain)funC() ChainAndError {
    fmt.Println("C")
    return ChainAndError{v, nil}
}

func main() {
    fmt.Println("Hello, playground")
    c := Chain{}
    result := c.funA().funB().funC() // line 24
    fmt.Println(result.error)
}
