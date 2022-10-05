## Used Different Data in interface

```go
package main

import "fmt"

// *******************************************************
// (STRUCT)
// *******************************************************
type Repository struct {
     container  map[string]interface{}
}

func NewRepository() *Repository {
	return &Repository{make(map[string]interface{})}
}

func (this *Repository) Set(key string, value interface{}) {
	 this.container[key] = value
}

func (this *Repository) Get(key string) interface{} {
	 return this.container[key]
}

func (this *Repository) GetAll() *Repository {
	 return this
}

type Tp struct{
	Fm string
	Tr string
}

// Used
func main(){
    var Dat = Tp{} 
    Dat.Fm="Fm name"
    Dat.Tr="Tr name"

    // Used
    rep := NewRepository()
    rep.Set("pr",   "pr   | printer")
    rep.Set("prom", "prom | Telegraph")
    rep.Set("tr",    Dat)

    // Print
    fmt.Println("Get:",rep.Get("pr"))
    fmt.Println("Get:",rep.Get("prom"))
	fmt.Println("Get:",rep.GetAll())
}
```
