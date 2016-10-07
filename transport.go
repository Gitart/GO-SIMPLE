package main

import (
    "fmt"
)

// transport is an http.RoundTripper that keeps track of the in-flight
// request and implements hooks to report HTTP tracing events.
type statistics struct {
     numbers   float64
     mean      float64
     median    float64
}

// 
func (v *statistics) Add() (float64){
     z:=v.mean+v.numbers+v.numbers/2
     return z
}



func getStats(numbers float64) (stats statistics) {
	   stats.numbers = numbers
	   stats.mean   = 1+stats.numbers
	   stats.median = 2*stats.numbers

	   return stats
}


// Программа statistics предоставляет доступ к единственной веб-
// странице на локальном компьютере. Ниже приводится функция
// main() программы:
func main() {
     
     var l statistics

     v:=getStats(1234)
     fmt.Println(v)

     l.mean=345
     l.numbers=3455
     l.median=777


     
     x:= l.Add()
     fmt.Println(x) 

     u:=statistics{1,89,4}
     fmt.Println(u.Add())      

}

















