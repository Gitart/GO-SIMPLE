package main
import (
   "fmt"
)

func main(){
	mainfu("123")
}

// Work with any type value
func mainfu(myInt interface{} ) {
   
// String
s, ok := myInt.(string)
if ok {
   fmt.Println("Yes string:", s)
}

// Init
k, ok := myInt.(int)
if ok {
   fmt.Println("Yes this INT:", k)
}

// Float 64
v, ok := myInt.(float64)
if ok {
fmt.Println("Yes float64", v)
} else {
fmt.Println("Failed without panicking!")
}

}
