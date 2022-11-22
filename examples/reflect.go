// Транслятор работает с адресами в памяти, для него нет "имен" в твоем представлении. 
// Есть адрес,тип,значение. На этом все. Извлечь имена полей структуры ты можешь путем перебора по индексу, по другому никак.

package main

import (
    "fmt"
    "reflect"
    "log"
)

type (
    User struct {
        ID string
        Name string
        FIO string
    }
)

func main() {
    user := &User{"IDsome","spouk","fio"}
    ShowStructure(user)
}
func ShowStructure(s interface{}) {
    a := reflect.ValueOf(s)
    numfield := reflect.ValueOf(s).Elem().NumField()
    if a.Kind() != reflect.Ptr {
        log.Fatal("wrong type struct")
    }
    for x := 0; x < numfield; x++ {
        fmt.Printf("Name field: `%s`  Type: `%s`\n", reflect.TypeOf(s).Elem().Field(x).Name,
            reflect.ValueOf(s).Elem().Field(x).Type())
    }
}
