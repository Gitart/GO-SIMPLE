package main

import "fmt"

type Cache interface {
	Get(key string) interface{}
	Set(key string, value interface{})
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

type Storage struct {
	Key   string
	Value interface{}
}

func (m *Sr) Set(s Storage) []Storage {
	var Dat map[string]Storage
	// s := interface{}(key)
	Dat = append(Dat, s)

	return Dat
}

func All(c Cache) {
	fmt.Println(c)
	fmt.Println(c.Get("sss"))
	fmt.Println(c.Get("sss"))
}

func main() {
	ss := St{Name: "sss", Id: "123"}
	sd := Sr{Name: "Test", Id: "2"}
	All(sd)
	All(ss)
}
