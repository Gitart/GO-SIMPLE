# Work with tag

```golang
// https://blog.golang.org/slices

package main
import "fmt"

// Structure
type TY struct{
	Name string
	Tags []string
}


//  Main
func main() {
	var Y TY
	Z := []string{"Первый","Второй","Третий","Четвертый"}
	Y.Tags=Z
    DeleteEl(Y,2)	
    fmt.Println(Y)
}


// Delete elements in tag
func DeleteEl(L TY, Num int) TY {
     RemoveElement(L.Tags,Num)
	 return L
}

// Adding
func Chng1(){
	var Y TY
	Z := []string{1:"t",2:"g","df",10:"ddd"}
	Z=append(Z,"Adding element")

	Y.Name="Name"
	Y.Tags=Z
	Y.Tags[0]="Change"

	fmt.Println(Y.Tags[4])
	fmt.Println(Y.Tags[0])
	// fmt.Println(Y)
}


// Adding
func Chng(){
	var Y TY
	Z := make([]string,5)

	Y.Name="Name"
	Y.Tags=Z
	Y.Tags[0]="Change"
	Y.Tags[4]="Changes"

	fmt.Println(Y.Tags[0])
	// fmt.Println(Y)
}



// Remove (v.1)
func RemoveElement(s []string, i int) []string {
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}

// Remove (v.2)
func RemoveIndex(s []int, index int) []int {
    return append(s[:index], s[index+1:]...)
}
```


package main

import (
    "fmt"
)

func RemoveIndex(s []int, index int) []int {
    return append(s[:index], s[index+1:]...)
}

func main() {
    all := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    fmt.Println(all) //[0 1 2 3 4 5 6 7 8 9]
    n := RemoveIndex(all, 5)
    fmt.Println(n) //[0 1 2 3 4 6 7 8 9]
}

