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
