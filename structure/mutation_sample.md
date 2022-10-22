src/ptrstruct/ptrstruct.go:

package ptrstruct

type PtrStruct struct {
    pn *int
}

func New(n int) *PtrStruct {
    return &PtrStruct{pn: &n}
}

func (s *PtrStruct) N() int {
    return *s.pn
}

func (s *PtrStruct) SetN(n int) {
    *s.pn = n
}

func (s *PtrStruct) Clone() *PtrStruct {
    // make a deep clone
    t := &PtrStruct{pn: new(int)}
    *t.pn = *s.pn
    return t
}
src/ptrstruct.go:

package main

import (
    "fmt"

    "ptrstruct"
)

func main() {
    ps := ptrstruct.New(42)
    fmt.Println(ps.N())
    pc := ps.Clone()
    fmt.Println(pc.N())
    pc.SetN(7)
    fmt.Println(pc.N())
    fmt.Println(ps.N())
}
Output:

src $ go run ptrstruct.go
42
42
7
42
src $ 
