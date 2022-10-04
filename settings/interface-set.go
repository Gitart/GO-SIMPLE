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
