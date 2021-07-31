package main

import (
	"fmt"
	"net/http"
	"github.com/go-humble/locstor"
)


func Test2(w http.ResponseWriter, r *http.Request){
	 err:=locstor.SetItem("foo", "bar")
	 if err!=nil{
	 	fmt.Println(err.Error())
	 }

	 fmt.Println("set bar")
}
