// a := []int{1, 2}
// b := []int{11, 22}
// a = append(a, b...) // a == [1 2 11 22]


package main

import (
	"fmt"

	
)

func main() {
     
SetG("fff","fff")


}


func SetG(keyval ...string) {
         
for _, s:=range keyval{
   fmt.Println(s)
} 

	kv := append([]string{"gl-go", "ddd"}, keyval...)
	fmt.Println(kv)
}



