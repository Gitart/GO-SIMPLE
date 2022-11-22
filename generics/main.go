package main

import "fmt"

// type parameter allows us to work on multiple types 
// here we defined T as type parameter allowing int and float64 type
func Sum[T int | float64| string|int64] (a , b T) T {
	return a + b
}


func main(){
    fmt.Println(Sum("sssss","-fgghhj"))
	fmt.Println(Sum(32,45))

    fmt.Println(Sum(32.78,45.77))
    fmt.Println(Sum(32.99,45.88))
    fmt.Println(Sum(32.77,45.87))
    fmt.Println(Sum(32.99,45.88))
}
