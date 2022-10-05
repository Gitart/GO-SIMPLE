![image](https://user-images.githubusercontent.com/3950155/193854117-981843d2-b011-402c-bbc7-630c4407cd6f.png)


## Work with interface 

1. Sample work with interface 
2. Examples     

## Useful Links 
[Other interface configuration](https://go.dev/play/p/ChqlpvGEKi)    
[Interested realization](https://github.com/Gitart/GO-SIMPLE/blob/master/interface/interface-set-get.go#L72)  

## Basick used

```go
// https://go.dev/play/p/ChqlpvGEKi

// My samples
// https://go.dev/play/p/KU82WDAjuJr
// https://go.dev/play/p/KYSCLpzeZMg

package main

import (
	"fmt"
)

type Repository struct {
	container map[string]interface{}
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

func (this *Repository) SetConfig(config *map[string]map[string]interface{}) {
	this.Set("config", config)
}

func (this *Repository) GetConfig(section string) map[string]interface{} {
	var configPtr *map[string]map[string]interface{}

	if val := this.Get("config"); val != nil {
		configPtr = val.(*map[string]map[string]interface{})
	}

	config := *configPtr

	if val, isPresent := config[section]; isPresent {
		return val
	}

	return make(map[string]interface{})
}

func main() {
	repo := NewRepository()
	config := map[string]map[string]interface{}{
		"app": {
			"hostname": "localhost",
		},
	}

	repo.SetConfig(&config)

	fmt.Println(repo.GetConfig("app")["hostname"].(string))
}
```

## With show 
```go
package main
import "fmt"

// *******************************************************
// Cache (BASIC INTERACE)
// *******************************************************
type Cache interface {
	Set(key string, value interface{})
	Get(key string) interface{}
}

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

// *******************************************************
// (STRUCT)
// *******************************************************
func NewSt() *St {
	return &St{make(map[string]interface{})}
}

// Kv 
type St struct {
     Kv map[string]interface{}
}

// Set 
func (m *St) Set(Key string, Value interface{}){
	 m.Kv[Key] = Value
}

func (m *St) Get(key string) interface{} {
	 return m.Kv[key]
}

// Get All
func (m *St) GetAll() *St {
	 return m
}

// *****************************************
// Preview work interface
// *****************************************
func All(c Cache) {
	fmt.Println("    >:", c)

    fmt.Println("tel  :----------->" , c.Get("te"))
	fmt.Println("pr   :----------->" , c.Get("pr"))
	fmt.Println("prom :----------->" , c.Get("prom"))
	fmt.Println("\n\n")
}

// *****************************************
// Main process 
// *****************************************
func main() {

    ss := NewSt()
    ss.Set("te",     "tel  |   telephone")

    srep := NewRepository()
    srep.Set("pr",   "pr   | printer")
    srep.Set("prom", "prom | Telegraph")
	
	All(srep)
	All(ss)
}


// *****************************************
// Additional Information
// *****************************************
func Info(){
    
    // Repository
    rep := NewRepository()
    rep.Set("News",     "Новостная лента")
    rep.Set("Old",      "Old news ")
    rep.Set("Forecast", "Prognoz")
    frep := rep.Get("Old")

    fmt.Println(frep)
    fmt.Println("-----------------------")

    // New Stat
	sr := NewSt()
	sr.Set("London", "England")
	sr.Set("Paris",  "France")
	fmt.Println("Country ", sr.Get("London"))

	sr.Set("tts34", "Элемент который знается")
	fmt.Println("Получение 34 елемента    : ", sr.Get("tts34"))
	fmt.Println("Получение всех елементов : ", sr.GetAll())
	fmt.Println("Получение всех елементов : ", *sr.GetAll() )
	fmt.Println("Количество елементов     : ", len(sr.GetAll().Kv) )
    
    dd:=sr.GetAll()
    fmt.Println(dd.Kv)
    
    // Range
    for k, r := range dd.Kv {
    	fmt.Println(k, "=", r)
    }
}
```
