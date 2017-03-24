## Загрузка JSON файла в базу

```golang
Shell "d:\MYSOLUTION\LoadJson\loadtxt.exe test Post " & "`{""id"":1222, ""sName"":""dddd""}`"

Copyright (C) <2014> SERVICE HOF
Description : Services and Calculation
Version : Version 1.0
Date Started : 03.11.2014
Author : Savchenko Arthur
Last Upadte Date : 10-11-2014
*/

package main

import (
"encoding/json"
// "fmt"
r "github.com/dancannon/gorethink"
"log"
"os"

// "time"
)

// Cесси¤
var sessionArray []*r.Session

// Declaration inetrfaces
type Mst map[string]interface{} // map - structure - interface
type Mif []interface{} // interface

//****************************************************************************************************
//
// Connect ini DB
//
//****************************************************************************************************

func Dbini() {
session, err := r.Connect(r.ConnectOpts{Address: "10.10.50.16:28015", Database: "test"})

// ќбработка ошибок
if err != nil {
log.Fatalln(err)
}

sessionArray = append(sessionArray, session)
}

//****************************************************************************************************
//
// Master Load Json File
// Listen Port 5555
// Кодировка обязательно должна быть Bom
//****************************************************************************************************
func main() {
Dbini()
loadPostsJson()
}

func loadPostsJson() {

// os.Args[0] -- полный путь без аргументов
// os.Args[1] -- первый аргумент база
// os.Args[2] -- первый аргумент таблица
// os.Args[3] -- Json файл
// os.Args[4] -- 1 - нужно ли очищать таблицу перед добавлением информации

// Параметры
NameBase := os.Args[1] // База
NameTables := os.Args[2] // Таблица

var StrJson string

StrJson = os.Args[3] // Файл JSON
StrJson = "`" + StrJson + "`"
//StrJsons = StrJson

//DelFile := os.Args[4] // A=добавление D=удаление

// Текущее время
//var CurTime = time.Now().Format("2006-01-02 15:04:05")

m := make(map[string]interface{})

byt := []byte(StrJson)
errj := json.Unmarshal(byt, &m)

//fmt.Println(StrJson)
// Обработка ошибок
if errj != nil {
//log.Fatalln(errj)
//panic("\n \n Problem with Load Document........................................!!!! \n \n")
}

//r.Db("test").Table("Post").Insert(r.Expr(m).Merge(Mst{"InsertTime": CurTime}).Merge(Mst{"Sequence": MaxID()})).Run(sessionArray[0])
r.Db(NameBase).Table(NameTables).Insert(m).Run(sessionArray[0])
}
```

