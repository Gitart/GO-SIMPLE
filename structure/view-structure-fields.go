package main

import (
	"fmt"
	"reflect"
	"time"
)

type Book struct {
	ID uint64 
	Name uint64 

}

func main() {
	book := Book{}
	e := reflect.ValueOf(&book).Elem()

	for i := 0; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name
		varType := e.Type().Field(i).Type
		varValue := e.Field(i).Interface()
		fmt.Printf("%v %v %v\n", varName, varType, varValue)
	}
}
