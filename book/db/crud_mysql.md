## Пример Golang CRUD с использованием MySQL с нуля
В этом уроке мы увидим пример программы, чтобы узнать, как выполнять операции CRUD с базой данных с использованием Golang и MySQL. CRUD - это аббревиатура для создания, чтения, обновления и удаления. Операции CRUD - это основные манипуляции с данными для базы данных.

В этом примере мы собираемся создать интерфейс в качестве интерфейса базы данных для обработки этих операций. У нас есть таблица сотрудников, содержащая информацию о сотрудниках, такую ​​как идентификатор, имя и город. С этой таблицей мы должны выполнить CRUD, используя MySQL.

### Шаг 1: Подготовьте и импортируйте драйвер MySQL в ваш проект
Используя Git Bash, сначала установите драйвер для пакета базы данных MySQL от Go. Запустите команду ниже и установите драйвер MySQL

```
go get  - u github.com / go - sql - driver / mysql
```

Теперь создайте базу данных Goblog

1\. Откройте PHPMyAdmin / SQLyog или любой другой инструмент управления базами данных MySQL, который вы используете.
2\. Создайте новую базу данных "goblog"

---

## Шаг 2: Создание таблицы сотрудников

Выполните следующий запрос SQL, чтобы создать таблицу с именем Employee в базе данных MySQL. Мы будем использовать эту таблицу для всех наших будущих операций.

```sql
DROP TABLE IF EXISTS `employee`;
CREATE TABLE `employee` (
  `id` int(6) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL,
  `city` varchar(30) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_
```

## Шаг 3: Создание структуры, обработчика и функции обработчика

Давайте создадим файл с именем `main.go` и поместим в него следующий код.
Обычно мы импортируем базу данных / sql и используем sql для выполнения запросов к базе данных.
Функция `dbConn` открывает соединение с драйвером MySQL.
Мы создадим `Employee` структуру, которая имеет следующие свойства: Id, Name и City.


```golang
package main

import (
    "database/sql"
    "log"
    "net/http"
    "text/template"

    _ "github.com/go-sql-driver/mysql"
)

type Employee struct {
    Id    int
    Name  string
    City string
}

func dbConn() (db *sql.DB) {
    dbDriver := "mysql"
    dbUser := "root"
    dbPass := "root"
    dbName := "goblog"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
    if err != nil {
        panic(err.Error())
    }
    return db
}

var tmpl = template.Must(template.ParseGlob("form/*"))

func Index(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    selDB, err := db.Query("SELECT * FROM Employee ORDER BY id DESC")
    if err != nil {
        panic(err.Error())
    }
    emp := Employee{}
    res := []Employee{}
    for selDB.Next() {
        var id int
        var name, city string
        err = selDB.Scan(&id, &name, &city)
        if err != nil {
            panic(err.Error())
        }
        emp.Id = id
        emp.Name = name
        emp.City = city
        res = append(res, emp)
    }
    tmpl.ExecuteTemplate(w, "Index", res)
    defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")
    selDB, err := db.Query("SELECT * FROM Employee WHERE id=?", nId)
    if err != nil {
        panic(err.Error())
    }
    emp := Employee{}
    for selDB.Next() {
        var id int
        var name, city string
        err = selDB.Scan(&id, &name, &city)
        if err != nil {
            panic(err.Error())
        }
        emp.Id = id
        emp.Name = name
        emp.City = city
    }
    tmpl.ExecuteTemplate(w, "Show", emp)
    defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
    tmpl.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")
    selDB, err := db.Query("SELECT * FROM Employee WHERE id=?", nId)
    if err != nil {
        panic(err.Error())
    }
    emp := Employee{}
    for selDB.Next() {
        var id int
        var name, city string
        err = selDB.Scan(&id, &name, &city)
        if err != nil {
            panic(err.Error())
        }
        emp.Id = id
        emp.Name = name
        emp.City = city
    }
    tmpl.ExecuteTemplate(w, "Edit", emp)
    defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
        name := r.FormValue("name")
        city := r.FormValue("city")
        insForm, err := db.Prepare("INSERT INTO Employee(name, city) VALUES(?,?)")
        if err != nil {
            panic(err.Error())
        }
        insForm.Exec(name, city)
        log.Println("INSERT: Name: " + name + " | City: " + city)
    }
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
        name := r.FormValue("name")
        city := r.FormValue("city")
        id := r.FormValue("uid")
        insForm, err := db.Prepare("UPDATE Employee SET name=?, city=? WHERE id=?")
        if err != nil {
            panic(err.Error())
        }
        insForm.Exec(name, city, id)
        log.Println("UPDATE: Name: " + name + " | City: " + city)
    }
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    emp := r.URL.Query().Get("id")
    delForm, err := db.Prepare("DELETE FROM Employee WHERE id=?")
    if err != nil {
        panic(err.Error())
    }
    delForm.Exec(emp)
    log.Println("DELETE")
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}

func main() {
    log.Println("Server started on: http://localhost:8080")
    http.HandleFunc("/", Index)
    http.HandleFunc("/show", Show)
    http.HandleFunc("/new", New)
    http.HandleFunc("/edit", Edit)
    http.HandleFunc("/insert", Insert)
    http.HandleFunc("/update", Update)
    http.HandleFunc("/delete", Delete)
    http.ListenAndServe(":8080", nil)
}
```


#### Шаг 4: Создание файлов шаблонов

Теперь пришло время создать файлы шаблонов нашего приложения CRUD. Создайте `form` папку в том же месте, где мы создали `main.go` .

а) Давайте создадим файл с именем `Index.tmpl` внутри `form` папки и поместим в него следующий код.

```golang
{{ define "Index" }}
  {{ template "Header" }}
    {{ template "Menu"  }}
    <h2> Registered </h2>
    <table border="1">
      <thead>
      <tr>
        <td>ID</td>
        <td>Name</td>
        <td>City</td>
        <td>View</td>
        <td>Edit</td>
        <td>Delete</td>
      </tr>
       </thead>
       <tbody>
    {{ range . }}
      <tr>
        <td>{{ .Id }}</td>
        <td> {{ .Name }} </td>
        <td>{{ .City }} </td> 
        <td><a href="/show?id={{ .Id }}">View</a></td>
        <td><a href="/edit?id={{ .Id }}">Edit</a></td>
        <td><a href="/delete?id={{ .Id }}">Delete</a><td>
      </tr>
    {{ end }}
       </tbody>
    </table>
  {{ template "Footer" }}
{{ end }}
```

b) Now create another file named Header.tmpl inside the same form folder and put the following code inside it.

```golang
{{ define "Header" }}
<!DOCTYPE html>
<html lang="en-US">
    <head>
        <title>Golang Mysql Curd Example</title>
        <meta charset="UTF-8" />
    </head>
    <body>
        <h1>Golang Mysql Curd Example</h1>   
{{ end }}

c) Now create another file named Footer.tmpl inside the same form folder and put the following code inside it.

{{ define "Footer" }}
    </body>
</html>
{{ end }}
```

d) Now create another file named Menu.tmpl inside the same form folder and put the following code inside it.
```golang
{{ define "Menu" }}
<a href="/">HOME</a> | 
<a href="/new">NEW</a>
{{ end }}
```

e) Next, we have to create Show.tmpl file for item details page, so again create this file in form folder.


```golang
{{ define "Show" }}
  {{ template "Header" }}
    {{ template "Menu"  }}
    <h2> Register {{ .Id }} </h2>
      <p>Name: {{ .Name }}</p>
      <p>City:  {{ .City }}</p><br /> <a href="/edit?id={{ .Id }}">Edit</a></p>
  {{ template "Footer" }}
{{ end }}
```


f) Now we create new blade file for create new item, it's call New.tmpl file inside form.


```golang
{{ define "New" }}
  {{ template "Header" }}
    {{ template "Menu" }} 
   <h2>New Name and City</h2>  
    <form method="POST" action="insert">
      <label> Name </label><input type="text" name="name" /><br />
      <label> City </label><input type="text" name="city" /><br />
      <input type="submit" value="Save user" />
    </form>
  {{ template "Footer" }}
{{ end }}
```

g) At last, we need to create Edit.tmpl file for update item, so again create this file in form folder.
```golang
{{ define "Edit" }}
  {{ template "Header" }}
    {{ template "Menu" }} 
   <h2>Edit Name and City</h2>  
    <form method="POST" action="update">
      <input type="hidden" name="uid" value="{{ .Id }}" />
      <label> Name </label><input type="text" name="name" value="{{ .Name }}"  /><br />
      <label> City </label><input type="text" name="city" value="{{ .City }}"  /><br />
      <input type="submit" value="Save user" />
    </form><br />    
  {{ template "Footer" }}
{{ end }}
```

After a long journey finally we've created all files of our CRUD application with Golang and MySQL.

![CRUD application with Golang and MySQL](https://www.golangprograms.com/media/wysiwyg/1.JPG)

![CRUD application with Golang and MySQL](https://www.golangprograms.com/media/wysiwyg/2.JPG)

---

Run the following command

go run main.go

---

Load the following URL

http://localhost:8080/

![CRUD application with Golang and MySQL](https://www.golangprograms.com/media/wysiwyg/3.JPG)

![CRUD application with Golang and MySQL](https://www.golangprograms.com/media/wysiwyg/4.JPG)

![CRUD application with Golang and MySQL](https://www.golangprograms.com/media/wysiwyg/5.JPG)
