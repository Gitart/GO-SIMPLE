## 🚡 Cortage

**Данная процедура позволяет :**
1. Заполнять в любом месте опердленные поля в структуре
2. Дополнять поля с уже существующими данными если это []
3. Управлять полями 

## Sample 
```go
package main

import (
	"fmt"
  "encoding/json"
)


type search struct {
   whereConditions  []map[string]interface{}
   orConditions     []map[string]interface{}
   notConditions    []map[string]interface{}
   havingConditions []map[string]interface{}
   joinConditions   []map[string]interface{}
   initAttrs        []interface{}
   assignAttrs      []interface{}
   selects          map[string]interface{}
   omits            []string
   orders           []interface{}
   offset           interface{}
   limit            interface{}
   group            string
   tableName        string
   raw              bool
   Unscoped         bool
   Title            string
   ignoreOrderQuery bool
}

func (s *search) Where(query interface{}, values ...interface{}) *search {
   s.whereConditions = append(s.whereConditions, 
                              map[string]interface{}{"query": query, "args": values})
   return s
}

func (s *search) Or(query interface{}, values ...interface{}) *search {
   s.orConditions = append(s.orConditions, map[string]interface{}{"query": query, "args": values})
   return s
}

func (s *search) Omit(columns ...string) *search {
   s.omits = columns
   return s
}

func (s *search) Grp (column string) *search {
      s.group = column
      s.Title = column
      return s
}

func main(){
   var D search 
   D.Where("ttt","tttt").Or("or","or1","or2")
   D.Where("Where2","Where22")
   D.Or("A-or","A-or1","A-or2")
   D.Omit("A-omit","A-Omit","A-Omit2")
   D.Grp("testGropu")

   // fmt.Printf("%T",D)
   // fmt.Printf("%+v",D)
   
   fmt.Println(D,"------------------------------------\n")
   s,_:=json.Marshal(D)
     fmt.Println(string(s))
   
   // Tr:="N"
   // if Tr=="Y"{
   //    goto GoodEnd
   // }

   bad:=false
   if bad {
      goto BadEnd
   }

   goto End
    
   // GoodEnd:
   //    fmt.Println("Хороший документ получился")  

   BadEnd:
      fmt.Println("Что-то идет вооще не так")   

   End:
      fmt.Println("Выходим нахрен блин....")   
}
```
