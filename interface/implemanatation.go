// https://go.dev/play/p/KU82WDAjuJr

package main

import "fmt"

type Cache interface {
	Get(key string) interface{}
	Set(key string) interface{}
}

type St struct {
	Name string
	Id   string
}

func (m *St) Get(key string) interface{} {
	s := interface{}(key)
	return s
}

func (m *St) Set(key string) interface{} {
	s := interface{}(key)
	return s
}

type Sr struct {
	Name string
	Id   string
}

func (m *Sr) Get(key string) interface{} {
	s := interface{}(key)
	return s
}

func (m *Sr) Set(key string) interface{} {
	s := interface{}(key)
	return s
}

func All(c Cache) {
	fmt.Println(c)
	fmt.Println(c.Get("sss"))
	fmt.Println(c.Get("sss"))

}

func main() {
	ss := St{Name: "sss", Id: "123"}
	sd := Sr{Name: "Test", Id: "2"}
	All(&sd)
	All(&ss)
}
