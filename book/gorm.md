## Golang RESTful API с использованием GORM и Gorilla Mux
REST расшифровывается как представительский государственный трансферт. Название «Представительный государственный перевод» (REST) ​​было придумано Роем Филдингом из Калифорнийского университета. Это очень упрощенный и легкий веб-сервис по сравнению с SOAP или WSDL. Производительность, масштабируемость, простота, мобильность являются основными принципами REST API.

## RESTful веб-API
Суть архитектуры REST состоит из клиента и сервера. API REST позволяет различным системам подключаться и отправлять / получать данные прямым способом. Сервер принимает входящие сообщения, затем отвечает на них, в то время как клиент создает соединение, а затем доставляет сообщения на сервер.

RESTful-клиент будет HTTP-клиентом, а RESTful-сервер будет HTTP-сервером. Каждый вызов REST API имеет отношение между глаголом HTTP и URL. Резервы (данные или бизнес-логика) в базе данных в приложении могут быть определены с помощью конечной точки API в REST.

### Установка
Установка пакетов GORM, Gorilla Mux и MySQL довольно проста. Вам просто нужно запустить ниже трех команд, используя GIT Bash или Putty:

```
go get -u github.com/gorilla/mux
go get -u github.com/jinzhu/gorm
go get -u github.com/go-sql-driver/mysql
```

## Develop
Теперь мы готовы к работе. Давайте создадим новую папку api внутри папки github.com . В папке api мы создадим наше приложение RESTful API.

**src/github.com/api/app/model** follow the folder location and create two new folders **app** and **model**.

Inside **model** create a file called **model.go** and add the following code:

```golang
package model

import (
	"github.com/jinzhu/gorm"
	\_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Employee struct {
	gorm.Model
	Name   string \`gorm:"unique" json:"name"\`
	City   string \`json:"city"\`
	Age    int    \`json:"age"\`
	Status bool   \`json:"status"\`
}

func (e \*Employee) Disable() {
	e.Status \= false
}

func (p \*Employee) Enable() {
	p.Status \= true
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db \*gorm.DB) \*gorm.DB {
	db.AutoMigrate(&Employee{})
	return db
}
```

---

**src/github.com/api/app/handler**

Create a file called **common.go** and add the following code:

```golang
package handler

import (
	"encoding/json"
	"net/http"
)

// respondJSON makes the response with payload as json format
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err :\= json.Marshal(payload)
	if err !\= nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(\[\]byte(err.Error()))
		return
	}
	w.Header().Set("Content\-Type", "application/json")
	w.WriteHeader(status)
	w.Write(\[\]byte(response))
}

// respondError makes the error response with payload as json format
func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, map\[string\]string{"error": message})
}
```

---

**src/github.com/api/app/handler**

Create a file called **employees.go** and add the following code:

```golang
package handler

import (
	"encoding/json"
	"net/http"

	"github.com/api/app/model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllEmployees(db \*gorm.DB, w http.ResponseWriter, r \*http.Request) {
	employees :\= \[\]model.Employee{}
	db.Find(&employees)
	respondJSON(w, http.StatusOK, employees)
}

func CreateEmployee(db \*gorm.DB, w http.ResponseWriter, r \*http.Request) {
	employee :\= model.Employee{}

	decoder :\= json.NewDecoder(r.Body)
	if err :\= decoder.Decode(&employee); err !\= nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err :\= db.Save(&employee).Error; err !\= nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, employee)
}

func GetEmployee(db \*gorm.DB, w http.ResponseWriter, r \*http.Request) {
	vars :\= mux.Vars(r)

	name :\= vars\["name"\]
	employee :\= getEmployeeOr404(db, name, w, r)
	if employee \== nil {
		return
	}
	respondJSON(w, http.StatusOK, employee)
}

func UpdateEmployee(db \*gorm.DB, w http.ResponseWriter, r \*http.Request) {
	vars :\= mux.Vars(r)

	name :\= vars\["name"\]
	employee :\= getEmployeeOr404(db, name, w, r)
	if employee \== nil {
		return
	}

	decoder :\= json.NewDecoder(r.Body)
	if err :\= decoder.Decode(&employee); err !\= nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err :\= db.Save(&employee).Error; err !\= nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, employee)
}

func DeleteEmployee(db \*gorm.DB, w http.ResponseWriter, r \*http.Request) {
	vars :\= mux.Vars(r)

	name :\= vars\["name"\]
	employee :\= getEmployeeOr404(db, name, w, r)
	if employee \== nil {
		return
	}
	if err :\= db.Delete(&employee).Error; err !\= nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func DisableEmployee(db \*gorm.DB, w http.ResponseWriter, r \*http.Request) {
	vars :\= mux.Vars(r)

	name :\= vars\["name"\]
	employee :\= getEmployeeOr404(db, name, w, r)
	if employee \== nil {
		return
	}
	employee.Disable()
	if err :\= db.Save(&employee).Error; err !\= nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, employee)
}

func EnableEmployee(db \*gorm.DB, w http.ResponseWriter, r \*http.Request) {
	vars :\= mux.Vars(r)

	name :\= vars\["name"\]
	employee :\= getEmployeeOr404(db, name, w, r)
	if employee \== nil {
		return
	}
	employee.Enable()
	if err :\= db.Save(&employee).Error; err !\= nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, employee)
}

// getEmployeeOr404 gets a employee instance if exists, or respond the 404 error otherwise
func getEmployeeOr404(db \*gorm.DB, name string, w http.ResponseWriter, r \*http.Request) \*model.Employee {
	employee :\= model.Employee{}
	if err :\= db.First(&employee, model.Employee{Name: name}).Error; err !\= nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &employee
}
```

---

**src/github.com/api/app**

Create a file called **app.go** and add the following code:

```golang
package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/api/app/handler"
	"github.com/api/app/model"
	"github.com/api/config"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// App has router and db instances
type App struct {
	Router \*mux.Router
	DB     \*gorm.DB
}

// App initialize with predefined configuration
func (a \*App) Initialize(config \*config.Config) {
	dbURI :\= fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
		config.DB.Charset)

	db, err :\= gorm.Open(config.DB.Dialect, dbURI)
	if err !\= nil {
		log.Fatal("Could not connect database")
	}

	a.DB \= model.DBMigrate(db)
	a.Router \= mux.NewRouter()
	a.setRouters()
}

// Set all required routers
func (a \*App) setRouters() {
	// Routing for handling the projects
	a.Get("/employees", a.GetAllEmployees)
	a.Post("/employees", a.CreateEmployee)
	a.Get("/employees/{title}", a.GetEmployee)
	a.Put("/employees/{title}", a.UpdateEmployee)
	a.Delete("/employees/{title}", a.DeleteEmployee)
	a.Put("/employees/{title}/disable", a.DisableEmployee)
	a.Put("/employees/{title}/enable", a.EnableEmployee)
}

// Wrap the router for GET method
func (a \*App) Get(path string, f func(w http.ResponseWriter, r \*http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Wrap the router for POST method
func (a \*App) Post(path string, f func(w http.ResponseWriter, r \*http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Wrap the router for PUT method
func (a \*App) Put(path string, f func(w http.ResponseWriter, r \*http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Wrap the router for DELETE method
func (a \*App) Delete(path string, f func(w http.ResponseWriter, r \*http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Handlers to manage Employee Data
func (a \*App) GetAllEmployees(w http.ResponseWriter, r \*http.Request) {
	handler.GetAllEmployees(a.DB, w, r)
}

func (a \*App) CreateEmployee(w http.ResponseWriter, r \*http.Request) {
	handler.CreateEmployee(a.DB, w, r)
}

func (a \*App) GetEmployee(w http.ResponseWriter, r \*http.Request) {
	handler.GetEmployee(a.DB, w, r)
}

func (a \*App) UpdateEmployee(w http.ResponseWriter, r \*http.Request) {
	handler.UpdateEmployee(a.DB, w, r)
}

func (a \*App) DeleteEmployee(w http.ResponseWriter, r \*http.Request) {
	handler.DeleteEmployee(a.DB, w, r)
}

func (a \*App) DisableEmployee(w http.ResponseWriter, r \*http.Request) {
	handler.DisableEmployee(a.DB, w, r)
}

func (a \*App) EnableEmployee(w http.ResponseWriter, r \*http.Request) {
	handler.EnableEmployee(a.DB, w, r)
}

// Run the app on it's router
func (a \*App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
```

---

**src/github.com/api/config**

Use your database connection credentials and create a file called **config.go** and add the following code:

```golang
package config

type Config struct {
	DB \*DBConfig
}

type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Name     string
	Charset  string
}

func GetConfig() \*Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Username: "root",
			Password: "root",
			Name:     "employee",
			Charset:  "utf8",
		},
	}
}
```

---

Now create **src** at root location.

Create a file called **main.go** and add the following code:

```golang
package main

import (
	"github.com/api/app"
	"github.com/api/config"
)

func main() {
	config :\= config.GetConfig()

	app :\= &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
```


#### Using command line or putty run command "go run main.go"

---

## Testing RESTful API methods

## GET

The GET method is a very common HTTP method in web applications used to retrieve a resource and should never be used to update any record. Typically, a body is never passed with a GET request to request a resource from our HTTP web servers.

#### Request

GET /employees HTTP/1.1
GET /employees/title HTTP/1.1

GET request type used in our scenario to request the data of all employees.

![](https://www.golangprograms.com/media/wysiwyg/api/get.jpg)

---

## POST

We used an HTML form to create a new resource via a non\-idempotent action which we called POST method. The POST method to send data to the server either through an asynchronous call or to execute a controller.

#### Request

POST /employees/data HTTP/1.1

Content\-Type: application/json
Content\-Length: xxxx

{"name": "John", "age": 42, "city": "Tokyo"}

POST request type used in our scenario to create new employee record.

![](https://www.golangprograms.com/media/wysiwyg/api/post.jpg)

---

## PUT

The PUT method is used to update a changeable resource and include the resource locator. Whenever we have to update a record that we have created earlier or want to create a record if it does not exist, using idempotent method calls which we refer as PUT method.

#### Request

PUT /employees/name HTTP/1.1

Content\-Type: application/json
Content\-Length: xxxx

{ "age": 30}

PUT request type used in our scenario to update employee age.

![](https://www.golangprograms.com/media/wysiwyg/api/put.jpg)

---

## DELETE

The DELETE method is always used to remove a record that is no longer required. In DELETE method we pass the ID of the resource as part of the path rather than in the body of the request.

#### Request

DELETE /employees/name HTTP/1.1

DELETE request type used in our scenario to remove employee record.

![](https://www.golangprograms.com/media/wysiwyg/api/delete.jpg)
