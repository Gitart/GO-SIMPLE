package main
 
import (
    "fmt"
)
 
func main() {
    fmt.Println("1) Generate Power Plant Report")
    fmt.Println("2) Generate Power Grid Report")
    fmt.Println("Please choose an option: ")
     
    var option string
     
    fmt.Scanln(&option)
     
    println(option)
}
