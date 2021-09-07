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

