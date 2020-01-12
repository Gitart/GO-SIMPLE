// https://gist.github.com/jochasinga/0af0c5064e236aafff38
// https://medium.com/code-zen/dynamically-creating-instances-from-key-value-pair-map-and-json-in-go-feef83ab9db2
// https://ru.stackoverflow.com/questions/526069/golang-%D0%B7%D0%B0%D0%BF%D0%B8%D1%81%D1%8C-%D0%B2%D0%B8%D0%B4%D0%B0-mapstringinterface
// https://ashirobokov.wordpress.com/2016/09/22/json-golang-cheat-sheet/..
// https://medium.com/code-zen/dynamically-creating-instances-from-key-value-pair-map-and-json-in-go-feef83ab9db2

package main
import "fmt"

type Mii interface{}                                          // Interface
type Mif []interface{}                                        // Cрез Interface
type Msr []string                                             // Срез String
type Mst map[string]interface{}                               // Map - string - interface
type Mss map[string]string                                    // Map - string - string
type Msi map[string]int64                                     // Map - string - int64
type Mis map[int64]string                                     // Map - int64 - string
type Msl map[int]string                                       // Map - int   - string
type Mil []int64                                              // Array int64  


type Person struct{
    Name   string
    Title  string 
}

func main() {
    // Первый вариант
	var a Mst
    a = Mst{}
    a["test"] = 1

    // Второй вариант
    b:= make(Mst)
    b["test"] = 1
    
    // 10 Array
    l:=make([]Person,10)
    l[0].Title = "New"
    l[0].Name  = "Timms"

    l[1].Title = "Moddle"
    l[1].Name  = "Oleg"

    l[2].Title = "Old"
    l[2].Name  = "Brad"

    
    // var ss map[string]interface{}{} -- Не работает
    // Работает
    var ss map[string]interface{}
    ss=map[string]interface{}{}

    ss["sss"]="map-string-interface{}{}"
    fmt.Println("map[string]inetrface{}{}....  ", ss)

    // V
    v:= map[string]interface{}{}
    v["Name"]="sss"
    fmt.Println("[]Mst append ....  ", v)

    // Добавление в массив - первый вариант
	var c []Mst
	c=append(c,b)
    c=append(c,b)

    // Добавление в массив - второй вариант
    d := []Mst{
    	        Mst{"eee":"dddd"},
    	        Mst{"eee1":"dddd"},
    	    }

    d[0]["testss"] = 1	    
    d[1]["test"] = 1

    fmt.Println("[]Mst append ....  ", c)
	fmt.Println("[]Mst ...........  ", d)
    fmt.Println("Mst{} ...........  ", a,b)
    fmt.Println("make(Mst)........  ", a,b)
    fmt.Println("make([]*Person)..  ", l)
}
