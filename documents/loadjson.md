## Инструкция
### Загружает файл JSON в базу RETHINKDB в опредленную базу - таблицу удаляя (не удаляя) перед закачкой данные





### run.bat
```bat
@echo off

* rem 1 аргумент - база данных (test)
* rem 2 аргумент - таблица (Post)
* rem 3 аргумент - имя json файла  (Holders.json)
* rem 4 аргумент - D удалять данные перед загрузкой

loadjsonfile.exe test Post Holders.json D >> log.txt
rem pause
```

### Compilation
```bat
@echo off
SETLOCAL
:: start

rem Path to current Programm API Service
SET GOPATH=D:\ORION\RETHINKDB\GO

rem путь к компилятору
SET GOROOT=C:\GO
SET PATH=%GOROOT%\BIN;%PATH%;
cls

title Run "GO" 
color 0f

rem echo "  "                 
rem echo ....................................................................
rem echo gopath = %gopath%
rem echo ....................................................................

rem go get -u "github.com/dancannon/gorethink"
go get github.com/beego/bee

rem go env
rem go build -o bin\app.exe d:\ORION\RETHINKDB\GO\main.go
rem go build -o loadjsonfile.exe loadjson.go
go build  -o service.exe service.go 
service.exe

rem Запись в лог файл о действиях запуска
rem echo %Date% %Time%  %date:~-10,2% Run >> d:\morion\rethinkdb\go\bin\log.txt

rem go run selectgo.go
rem go run d:\MORION\RETHINKDB\GO\main.go
@pause
```

## Programm

```golang
/*
 Copyright (C) <2014>  SERVICE HD OFFICE
 Description      : Services and Calculation
 Version          : Version 1.0
 Date Started     : 03.11.2014
 Author           : Savchenko Arthur
 Last Upadte Date : 10-11-2014
*/

package main

import (
	
	"encoding/json"
	"fmt"
	"io/ioutil"
	r "github.com/dancannon/gorethink"
	"log"
	"os"
	"time"
)

// Cесси¤
var sessionArray []*r.Session

// Declaration inetrfaces
type Mst map[string]interface{} // map - structure - interface
type Mif []interface{}          // interface

//****************************************************************************************************
//
// Connect ini DB
//
//****************************************************************************************************
func Dbini() {
	session, err := r.Connect(r.ConnectOpts{Address: "10.0.50.16:28015", Database: "test"})

	// Обработка ошибок
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

// Загрузка файла JSON
func loadPostsJson() string {
	// os.Args[0]  -- полный путь без аргументов
	// os.Args[1]  -- первый аргумент база
	// os.Args[2]  -- первый аргумент таблица
	// os.Args[3]  -- Json файл
	// os.Args[4]  -- 1 - нужно ли очищать таблицу перед добавлением информации

	// Параметры
	NameBase   := os.Args[1]  // База
	NameTables := os.Args[2]  // Таблица
	NameFile   := os.Args[3]  // Файл JSON
	DelFile    := os.Args[4]  // A=добавление D=удаление

	// Текущее время
	var CurTime = time.Now().Format("2006-01-02 15:04:05")

	//content, _ := ioutil.ReadFile("/Users/wangbin/Downloads/dump.json")
	//content, _ := ioutil.ReadFile("mp.json")
	content, _ := ioutil.ReadFile(NameFile)

	// var posts []*Post
	// Автоматически подходит для всех форматов Json
	var posts []*Mst 
	
	// Описание структуры Json
	// var posts []*Pt 
	json.Unmarshal(content, &posts)

	//r.Table("Post").Insert(posts).RunWrite(sess)
	
	// Очистка таблицы
	if DelFile == "D" {
	   r.Db(NameBase).Table(NameTables).Delete().Run(sessionArray[0])
	}

	// Добавление поля Insert Time
	// Добавление поля с текущем временем вставки
	// r.Table("Post").Insert(r.Expr(posts).Merge(Mst{"InsertTime": CurTime})).Run(sessionArray[0])
	// r.Db(NameBase).Table(NameTables).Insert(r.Expr(posts).Merge(Mst{"InsertTime": CurTime})).Run(sessionArray[0])
	defer r.Db(NameBase).Table(NameTables).Insert(posts).Run(sessionArray[0])

	if DelFile == "D" {
	   fmt.Printf(CurTime+".... Данные из таблицы %s сначала были удалены а потом добавлены...\n", NameTables)
	} else {
	   fmt.Printf(CurTime+".... Данные добавлены в таблицу %s . \n", NameTables)
	}

	return "Ok"
	// Return col записей всталенных
	// r.Db(NameBase).Table(NameTables).Count()
	/*
		var users []*models.User
		rows, _ := r.Table(models.UserTable).GetAllByIndex("username", "wangbin").Run(sess)
		rows.ScanAll(&users)
		author := users[0]
		r.Table(models.PostTable).Update(models.RethinkMap{"user_id": author.Id}).RunWrite(sess)
	*/
}

/*
// Load CSV
func loadCsv() {
	r := csv.NewReader("")
	row, err := r.Read()
	for err != nil {
		row.err = r.Read()
	}

	if err != os.EOF {
		fmt.Println("Error", err)
	}
}
*/
```
