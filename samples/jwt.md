# Создание и развёртывание REST API с помощью Go, PostgreSQL, JWT и GORM

[Переводы](https://tproger.ru/translations/ "Все посты рубрики «Переводы»"), 28 июня 2019 в 08:09 41 986

13

Поделиться

В этом руководстве расскажем, как разрабатывать и развёртывать защищённые REST API, используя язык программирования Go.

## Почему именно Go

Go — очень интересный язык. Он обладает строгой типизацией, очень быстро компилируется, а его производительность сравнима с C++. Go имеет goroutines — гораздо более эффективную замену для Threads, а также даёт возможность использовать статическую типизацию для создания web-приложений.

## Что будем создавать

Мы собираемся создать приложение для управления контактами. Созданный API позволит пользователям добавлять контакты в свои профили и восстановить их, если телефон потеряется.

## Что для этого понадобится

У вас уже должны быть установлены следующие пакеты:

*   Go;
*   PostgreSQL;
*   GoLand IDE  (не обязательно). Но в этом руководстве мы будем использовать её.

Также необходимо [настроить](https://github.com/golang/go/wiki/SettingGOPATH) переменную окружения `GOPATH`.

## Что такое REST

REST расшифровывается как Representational State Transfer. Это [механизм](https://ru.wikipedia.org/wiki/REST), используемый современными клиентскими приложениями для связи с базами данных и серверами через HTTP.

## Сборка приложения

Мы начнём с определения зависимостей пакетов, которые понадобятся для проекта. К счастью, стандартная библиотека Go достаточно богата ими, чтобы мы могли создать полноценный веб-сайт без использования сторонних фреймворков — смотрите на [этой](https://golang.org/pkg/net/http/) странице в разделе *Packages*.

Следующие пакеты облегчат нам работу:

*   gorilla / mux — мощный URL-маршрутизатор и диспетчер. Мы будем использовать этот пакет для сопоставления путей URL с их обработчиками.
*   jinzhu / gorm — отличная [ORM](https://ru.wikipedia.org/wiki/ORM)\-библиотека для Golang. Мы используем этот пакет, чтобы взаимодействовать с базой данных.
*   dgrijalva / jwt-go — используется для подписи и проверки токенов [JWT](https://ru.wikipedia.org/wiki/JSON_Web_Token).
*   joho / godotenv — используется для загрузки файлов `.env` в проект.

Чтобы установить любой из этих пакетов, откройте терминал и запустите команду

```bash
go get github.com/{package-name}
```

Эта команда установит пакеты в ваш `GOPATH`.

## Структура проекта

[![](https://tproger.ru/s3/uploads/2019/06/rest-api-golang-pic-1.jpg)](https://tproger.ru/s3/uploads/2019/06/rest-api-golang-pic-1.jpg)

Структура проекта отображается в левой боковой панели.

Файл `utils.go`:

```go
package utils

import (
	"encoding/json"
	"net/http"
)

func Message(status bool, message string) (map[string]interface{}) {
	return map[string]interface{} {"status" : status, "message" : message}
}

func Respond(w http.ResponseWriter, data map[string] interface{})  {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
```

Файл `utils.go` содержит удобные функции для работы с JSON. Обратите внимание на функции `Message()` и `Respond()` , прежде чем мы продолжим.

## Подробнее о JWT

[JSON Web Tokens](https://ru.wikipedia.org/wiki/JSON_Web_Token) — это открытый стандарт RFC 7519 для создания токенов доступа. Используется в передаче данных для аутентификации в клиент-серверных приложениях. В обычных веб-приложениях легко идентифицировать пользователей с помощью сессий, однако, когда API вашего веб-приложения взаимодействует, скажем, с клиентом Android или IOS, сессии становятся малопригодными для использования. С помощью JWT мы можем создать уникальный токен для каждого аутентифицированного пользователя. Этот токен будет включён в заголовок последующего запроса к API. Этот метод позволяет идентифицировать всех пользователей, которые выполняют вызовы API. Давайте посмотрим реализацию:

```go
package app

import (
	"net/http"
	u "lens/utils"
	"strings"
	"go-contacts/models"
	jwt "github.com/dgrijalva/jwt-go"
	"os"
	"context"
	"fmt"
)

var JwtAuthentication = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		notAuth := []string{"/api/user/new", "/api/user/login"} //Список эндпоинтов, для которых не требуется авторизация
		requestPath := r.URL.Path //текущий путь запроса

		//проверяем, не требует ли запрос аутентификации, обслуживаем запрос, если он не нужен
		for _, value := range notAuth {

			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		response := make(map[string] interface{})
		tokenHeader := r.Header.Get("Authorization") //Получение токена

		if tokenHeader == "" { //Токен отсутствует, возвращаем  403 http-код Unauthorized
			response = u.Message(false, "Missing auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		splitted := strings.Split(tokenHeader, " ") //Токен обычно поставляется в формате `Bearer {token-body}`, мы проверяем, соответствует ли полученный токен этому требованию
		if len(splitted) != 2 {
			response = u.Message(false, "Invalid/Malformed auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		tokenPart := splitted[1] //Получаем вторую часть токена
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil { //Неправильный токен, как правило, возвращает 403 http-код
			response = u.Message(false, "Malformed authentication token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		if !token.Valid { //токен недействителен, возможно, не подписан на этом сервере
			response = u.Message(false, "Token is not valid.")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		//Всё прошло хорошо, продолжаем выполнение запроса
		fmt.Sprintf("User %", tk.Username) //Полезно для мониторинга
		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r) //передать управление следующему обработчику!
	});
}

```

Комментарии внутри кода объясняют всё, что нужно знать, но в основном код создаёт *Middleware*, чтобы перехватывать все запросы, проверять наличие токена аутентификации (токена JWT), проверять, является ли он подлинным и действительным, а затем отправлять ошибку клиенту, если возникли какие-то проблемы.

[

Тестировщик мобильных и web-приложений

ООО «Табер Трейд», сеть магазинов «Подружка», Удалённо, До 120 000 ₽

tproger.ru

](https://tproger.ru/jobs/testirovshhik-mobilnogo-prilozhenija-i-sajta-for-podruzhka/?utm_source=in_text)

[Вакансии на tproger.ru](https://tproger.ru/jobs/)

Ниже вы увидите, как из запроса можно получить доступ к пользователю, который взаимодействует с API.

## Построение системы регистрации пользователей и входа

Необходимо, чтобы пользователи могли зарегистрироваться и войти в систему. Первое, что нужно сделать, это подключиться к базе данных. В проекте используется файл `.env` для хранения учётных данных для доступа к базе данных.

`.env` может выглядеть вот так:

```go
db_name = gocontacts
db_pass = **** //Это пароль по умолчанию для текущего пользователя в Windows для Postgresql
db_user = postgres
db_type = postgres
db_host = localhost
db_port = 5434
token_password = thisIsTheJwtSecretPassword //Не передавайте это через git!

```

Затем можно подключиться к базе данных, используя следующий код:

```go
package models

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jinzhu/gorm"
	"os"
	"github.com/joho/godotenv"
	"fmt"
)

var db *gorm.DB //база данных

func init() {

	e := godotenv.Load() //Загрузить файл .env
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Создать строку подключения
	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(&Account{}, &Contact{}) //Миграция базы данных
}

// возвращает дескриптор объекта DB
func GetDB() *gorm.DB {
	return db
}

```

Код делает очень простую вещь. Функция `init()` автоматически вызывается Go, код извлекает информацию о соединении из `.env` файла, затем строит строку соединения и использует её для соединения с базой данных.

## Создание точки входа в приложение

Итак, нам удалось создать промежуточный обработчик запроса (*Middleware*) для проверки JWT-токена и подключиться к базе данных. Следующим шагом будет создание точки входа в приложение. Фрагмент кода расположен ниже.

```go
package main

import (
	"github.com/gorilla/mux"
	"go-contacts/app"
	"os"
	"fmt"
	"net/http"
)

func main() {

	router := mux.NewRouter()
	router.Use(app.JwtAuthentication) // добавляем middleware проверки JWT-токена

	port := os.Getenv("PORT") //Получить порт из файла .env; мы не указали порт, поэтому при локальном тестировании должна возвращаться пустая строка
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":" + port, router) //Запустите приложение, посетите localhost:8000/api

	if err != nil {
		fmt.Print(err)
	}
}

```

Мы создаём новый объект `Router`, подключаем наше промежуточное программное обеспечение *JWT auth*, используя функцию маршрутизатора `Use()`, а затем приступаем к прослушиванию входящих запросов.

[![](https://tproger.ru/s3/uploads/2019/06/rest-api-golang-pic-4.jpg)](https://tproger.ru/s3/uploads/2019/06/rest-api-golang-pic-4.jpg)

Нажмите кнопку *play* слева от `func main()`, чтобы скомпилировать и запустить приложение. Если всё хорошо, вы не должны видеть ошибки в консоли. Если ошибка всё же возникла, ещё раз посмотрите на параметры подключения к базе данных, чтобы убедиться, что они корректны.

[![](https://tproger.ru/s3/uploads/2019/06/rest-api-golang-pic-3.png)](https://tproger.ru/s3/uploads/2019/06/rest-api-golang-pic-3.png)

Результаты. Произошла миграция БД, GORM преобразовал go struct в таблицы базы данных.

## Создание и аутентификация пользователей

Создайте новый файл `models/accounts.go`.

```go
package models

import (
	"github.com/dgrijalva/jwt-go"
	"lens/utils"
	"strings"
	"github.com/jinzhu/gorm"
	"os"
	"golang.org/x/crypto/bcrypt"
)

/*
Структура прав доступа JWT
*/
type Token struct {
	UserId uint
	jwt.StandardClaims
}

//структура для учётной записи пользователя
type Account struct {
	gorm.Model
	Email string `json:"email"`
	Password string `json:"password"`
	Token string `json:"token";sql:"-"`
}

//Проверить входящие данные пользователя ...
func (account *Account) Validate() (map[string] interface{}, bool) {

	if !strings.Contains(account.Email, "@") {
		return u.Message(false, "Email address is required"), false
	}

	if len(account.Password) < 6 {
		return u.Message(false, "Password is required"), false
	}

	//Email должен быть уникальным
	temp := &Account{}

	//проверка на наличие ошибок и дубликатов электронных писем
	err := GetDB().Table("accounts").Where("email = ?", account.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}
	if temp.Email != "" {
		return u.Message(false, "Email address already in use by another user."), false
	}

	return u.Message(false, "Requirement passed"), true
}

func (account *Account) Create() (map[string] interface{}) {

	if resp, ok := account.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)

	GetDB().Create(account)

	if account.ID <= 0 {
		return u.Message(false, "Failed to create account, connection error.")
	}

	//Создать новый токен JWT для новой зарегистрированной учётной записи
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	account.Password = "" //удалить пароль

	response := u.Message(true, "Account has been created")
	response["account"] = account
	return response
}

func Login(email, password string) (map[string]interface{}) {

	account := &Account{}
	err := GetDB().Table("accounts").Where("email = ?", email).First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email address not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Пароль не совпадает!!
		return u.Message(false, "Invalid login credentials. Please try again")
	}
	//Работает! Войти в систему
	account.Password = ""

	//Создать токен JWT
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString // Сохраните токен в ответе

	resp := u.Message(true, "Logged In")
	resp["account"] = account
	return resp
}

func GetUser(u uint) *Account {

	acc := &Account{}
	GetDB().Table("accounts").Where("id = ?", u).First(acc)
	if acc.Email == "" { //Пользователь не найден!
		return nil
	}

	acc.Password = ""
	return acc
}

```

В `account.go` на первый взгляд много загадок, давайте немного разберёмся с ними.

Первая часть создаёт две структуры: *Token* и *Account*. Они представляют токен JWT и учётную запись пользователя соответственно.

Функция `Validate()` проверяет данные, отправленные клиентами, а функция `Create()` создаёт новую учетную запись и генерирует токен JWT, который будет отправлен обратно клиенту, сделавшему запрос.

Функция `Login(username, password)` аутентифицирует существующего пользователя, затем генерирует токен JWT, если аутентификация прошла успешно.

Файл `authController.go`:

```go
package controllers

import (
	"net/http"
	u "go-contacts/utils"
	"go-contacts/models"
	"encoding/json"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //декодирует тело запроса в struct и завершается неудачно в случае ошибки
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := account.Create() //Создать аккаунт
	u.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //декодирует тело запроса в struct и завершается неудачно в случае ошибки
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := models.Login(account.Email, account.Password)
	u.Respond(w, resp)
}

```

Содержание очень простое. В нём имеется `handler` для `/user/new` и эндпоинт /`user/login`.

Добавьте следующий фрагмент `main.go`, чтобы зарегистрировать новые маршруты.

```go
router.HandleFunc("/api/user/new",
controllers.CreateAccount).Methods("POST")

router.HandleFunc("/api/user/login",
controllers.Authenticate).Methods("POST")

```

Этот код регистрирует `/user/new` и эндпоинт `/user/login`, а затем передаёт их соответствующие обработчики запросов.

Теперь перекомпилируйте код и зайдите на `localhost:8000/api/user/new` с помощью инструмента Postman, установите тело запроса `application/json`, как показано ниже:

[![](https://tproger.ru/s3/uploads/2019/06/rest-api-golang-pic-5.jpg)](https://tproger.ru/s3/uploads/2019/06/rest-api-golang-pic-5.jpg)

Если вы попытаетесь вызвать `/user/new` дважды с одними и теми же параметрами, вы получите ответ о том, что email уже существует.

## Создание контактов

Часть функциональности этого приложения позволяет нашим пользователям создавать и хранить контакты. Контакт будет иметь поля *name* и *phone*, и мы определим их как свойства структуры. Следующий код содержится в `models/contact.go`:

```go
package models

import (
	u "go-contacts/utils"
	"github.com/jinzhu/gorm"
	"fmt"
)

type Contact struct {
	gorm.Model
	Name string `json:"name"`
	Phone string `json:"phone"`
	UserId uint `json:"user_id"` //Пользователь, которому принадлежит этот контакт
}

/*
 Эта структурная функция проверяет обязательные параметры, отправленные через тело http-запроса
возвращает сообщение и true, если требование выполнено
*/
func (contact *Contact) Validate() (map[string] interface{}, bool) {

	if contact.Name == "" {
		return u.Message(false, "Contact name should be on the payload"), false
	}

	if contact.Phone == "" {
		return u.Message(false, "Phone number should be on the payload"), false
	}

	if contact.UserId <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	//Все обязательные параметры присутствуют
	return u.Message(true, "success"), true
}

func (contact *Contact) Create() (map[string] interface{}) {

	if resp, ok := contact.Validate(); !ok {
		return resp
	}

	GetDB().Create(contact)

	resp := u.Message(true, "success")
	resp["contact"] = contact
	return resp
}

func GetContact(id uint) (*Contact) {

	contact := &Contact{}
	err := GetDB().Table("contacts").Where("id = ?", id).First(contact).Error
	if err != nil {
		return nil
	}
	return contact
}

func GetContacts(user uint) ([]*Contact) {

	contacts := make([]*Contact, 0)
	err := GetDB().Table("contacts").Where("user_id = ?", user).Find(&contacts).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return contacts
}

```

Так же, как в `models/accounts.go`, мы создаём функцию `Validate()` для проверки переданных входных данных, возвращаем сообщение об ошибке, если происходит что-то, что нам не нужно, затем пишем функцию `Create()` для добавления контакта в базу данных.

Осталась только часть поиска контактов. Давайте сделаем её.

```go
router.HandleFunc("/api/me/contacts",
controllers.GetContactsFor).Methods("GET")

```

Добавьте приведённый выше фрагмент, чтобы сообщить маршрутизатору `main.go` о регистрации эндпоинта `/me/contacts`. Давайте создадим обработчик `controllers.GetContactsFor` для обработки запроса к API.

Файл `contactsController.go`:

```go
package controllers

import (
	"net/http"
	"go-contacts/models"
	"encoding/json"
	u "go-contacts/utils"
	"strconv"
	"github.com/gorilla/mux"
	"fmt"
)

var CreateContact = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user") . (uint) //Получение идентификатора пользователя, отправившего запрос
	contact := &models.Contact{}

	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	contact.UserId = user
	resp := contact.Create()
	u.Respond(w, resp)
}

var GetContactsFor = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		//Переданный параметр пути не является целым числом
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := models.GetContacts(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

```

То, что он делает, очень похоже на `authController.go`, но в основном он обрабатывает тело JSON и декодирует его в структуру *Contact*, и, если произошла ошибка, немедленно возвращает ответ. Если всё прошло хорошо, то вставляет контакты в базу данных.

## Извлечение контактов, принадлежащих пользователю

Теперь пользователи могут сохранять свои контакты. А что, если они захотят восстановить контакт в случае потери телефона? Посещение `/me/contacts` должно вернуть JSON структуру для контактов вызывающего API (текущего пользователя).

Как правило, эндпоинт для получения контактов пользователя должен выглядеть следующим образом: `/user/{userId}/contacts`. Использование `userId` опасно по следующим причинам:

*   каждый прошедший проверку пользователь может обработать запрос по этому пути;
*   контакты других пользователей будут возвращены без каких-либо проблем, это может привести к хакерской атаке.

Эту проблему решает JWT.

Мы можем легко получить `id` обработчика API вызывающего `r.Context().Value("user")`, зная, что мы установили это значение внутри `auth.go`.

```go
package controllers

import (
	"net/http"
	"go-contacts/models"
	"encoding/json"
	u "go-contacts/utils"
	"strconv"
	"github.com/gorilla/mux"
	"fmt"
)

var CreateContact = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user") . (uint) //Получение идентификатора пользователя, отправившего запрос
	contact := &models.Contact{}

	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	contact.UserId = user
	resp := contact.Create()
	u.Respond(w, resp)
}

var GetContactsFor = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		//Переданный параметр пути не является целым числом
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := models.GetContacts(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

```

[![](https://tproger.ru/s3/uploads/2019/06/rest-api-golang-pic-6.jpg)](https://tproger.ru/s3/uploads/2019/06/rest-api-golang-pic-6.jpg)

Код для этого проекта находится на [GitHub](https://github.com/adigunhammedolalekan/go-contacts).

## Развёртывание

Мы можем легко развернуть наше приложение на Heroku. Для этого скачайте [godep](https://github.com/tools/godep) — менеджер зависимостей для Golang, аналогичный npm для Node.js.

```bash
go get -u github.com/tools/godep
```

*   Откройте *GoLand terminal* и запустите *godep save*. Это создаст папки с названием *Godeps* и *vendor*.
*   Создайте аккаунт на [heroku](https://www.heroku.com/) и загрузите Heroku Cli, затем войдите под своими учётными данными.
*   После этого запустите `heroku create gocontacts`. Это создаст приложение на панели инструментов heroku, а также удалённый репозиторий git.
*   Выполните следующие команды для вставки вашего кода в heroku:

```bash
git add .
git commit -m "First commit"
git push heroku master

```

Если всё прошло хорошо, ваш экран должен выглядеть примерно так:

[![](https://tproger.ru/s3/uploads/2019/06/rest-api-golang-pic-7.png)](https://tproger.ru/s3/uploads/2019/06/rest-api-golang-pic-7.png)

Всё, ваше приложение было развёрнуто. Теперь нужно настроить удалённую базу данных PostgreSQL.

Для этого запустите `heroku addons:create heroku-postgresql:hobby-dev` для создания базы данных.

Почти всё готово. Остаётся лишь связаться с нашей удалённой базой данных.

Зайдите на [heroku.com](https://www.heroku.com/) и войдите под своими учётными данными. Найдите только что созданное приложение на панели инструментов и кликните на него. После этого нажмите «Настройки», затем выберите «Reveal Config Vars».

### URI-формат подключения PostgreSQL

`postgres://username:password@host/dbName`

В конфигурации `vars` вы обнаружите `DATABASE_URL`, который было автоматически добавлен в ваш файл `.env` при создании базы данных PostgreSQL.

Примечание Heroku автоматически заменяет локальную версию `.env` при развёртывании приложения. Из `var` мы будем извлекать параметр подключения к базе данных.

[![](https://tproger.ru/s3/uploads/2019/06/rest-api-golang-pic-8.jpg)](https://tproger.ru/s3/uploads/2019/06/rest-api-golang-pic-8.jpg)

[![](https://tproger.ru/s3/uploads/2019/06/rest-api-golang-pic-9.png)](https://tproger.ru/s3/uploads/2019/06/rest-api-golang-pic-9.png)

Мы извлекли параметры подключения к базе данных из автоматически сгенерированных переменных `DATABASE_URL`.

Если всё прошло хорошо, ваш API сейчас должен быть активен.
[![](https://tproger.ru/s3/uploads/2019/06/rest-api-golang-pic-10.jpg)](https://tproger.ru/s3/uploads/2019/06/rest-api-golang-pic-10.jpg)
