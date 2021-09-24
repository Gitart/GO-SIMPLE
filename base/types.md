package main
import "fmt"
 
type person struct{
    name string
    age int
}
 
type Ft map[string]interface{}

func main() {


// https://play.golang.org/p/2AaaJg_wWRh

    DOC_TYPES := map[int64]map[string]string{
		 1:  {"prefix": "P", "name":"Приход", "ff":"ffff"},
		-1:  {"prefix": "R", "name":"Приход"},
		-2:  {"prefix": "V", "name":"Возврат"},
                -3:  {"prefix": "S", "name":"Списание"},
	}
	
	
    Va:= make(Ft)       

    Va["ddd"] = "ssss"
    Va["dds"] = "ssssee"


 
    tom := person {name: "Tom", age: 22}
    var tomPointer *person = &tom
    tomPointer.age = 29
    fmt.Println(tom.age)        // 29
    (*tomPointer).age = 32
    fmt.Println(tom.age)        // 32

    fmt.Println(Va["ddd"])  

    fmt.Println(DOC_TYPES[1]["name"])      
    fmt.Println(DOC_TYPES[1]["ff"])      
      
}
