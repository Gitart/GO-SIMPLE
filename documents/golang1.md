# Веб-программирование в Go
## Часть 1
В go есть стандартный пакет net/http, который позволяет написать свой собственный веб-сервер. Начнем с написания простейшего примера: сервер будет обслуживать всего одну статическую страницу, на которой будет подсчитываться и выводиться счетчик ссылок:

```golang
 package main
 
 import (
     "fmt"
     "net/http"
 )
  
 type webCounter struct {
   count chan int
 }
 
 func NewCounter() *webCounter {
   counter := new(webCounter)
   counter.count = make(chan int, 1)
   go func() {
     for i:=1 ;; i++ { counter.count <- i }
   }()
   return counter
 }
 
 func (w *webCounter) ServeHTTP(r http.ResponseWriter, rq *http.Request) {
   if rq.URL.Path != "/" {
     r.WriteHeader(http.StatusNotFound)
   return
 }
 fmt.Fprintf(r, "You are visitor %d", <-w.count)
 }
 
 func main() {
   err := http.ListenAndServe(":8000", NewCounter())
   if err != nil {
     fmt.Printf("Server failed: ", err.Error())
   }
 }
 ```
 
Компилируем:
 #### go build
 
Запускаем откомпилированный бинарник, открываем броузер и набираем адрес:

#### http://localhost:8000/
 
и обновляем несколько раз страницу. При каждом обновлении страницы будет срабатывать метод ServeHTTP(), и для обработки запроса каждый раз будет генериться goroutine. Все они будут обслуживаться одним и тем же каналом до окончания работы веб-сервера, поэтому счетчик не теряет свое значение.
Напишем веб-клиента, который будем делать запрос к нашему серверу и выводить содержимое заглавной страницы:

```golang
 package main
 
 import "fmt"
 import "net/http"
 import "os"
 import "io"
  
 func main() {
 	client := &http.Client{}
 	client.CheckRedirect =
 		func(req *http.Request, via []*http.Request) error {
 		fmt.Fprintf(os.Stderr, "Redirect: %v\n", req.URL);
 		return nil
 	}
 	var url string
 	url = "http://localhost:8000/"
 	page, err := client.Get(url)
 	if err != nil {
 		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
 		return
 	}
 	io.Copy(os.Stdout, page.Body)
 	page.Body.Close()
 }
 ```
 
Структура http.Client содержит метод Get(url), которая получает содержимое этой страницы.
Стандартный пакет "text/template" позволяет использовать технологию шаблонов при генерации HTML. В следующем примере имеется шаблон для главной страницы:
```html
 <html>
 	<head>
 		<title>Go Web Counter</title>
 	</head>
 	<body>
 		<h1>A Simple Example</h1>
 		<p>You are visitor: {{.Counter}}</p>
 	</body>
 </html>
 ```
 
Соответственно меняется код самого сервера:

```golang
 package main
 import "fmt"
 import "net/http"
 import "text/template"
 
 type webCounter struct {
 	count chan int
 	template *template.Template
 }
 func NewCounter() *webCounter {
 	counter := new(webCounter)
 	counter.count = make(chan int, 1)
 	go func() {
 		for i:=1 ;; i++ { counter.count <- i }
 	}()
 	counter.template, _ = template.ParseFiles("counter.html")
 	return counter
 }
 func (w *webCounter) ServeHTTP(r http.ResponseWriter, rq *http.Request) {
 	if rq.URL.Path != "/" {
 		r.WriteHeader(http.StatusNotFound)
 		return
 	}
 	w.template.Execute(r, struct{Counter int}{<-w.count})
 }
 func main() {
 	err := http.ListenAndServe(":8000", NewCounter())
 	if err != nil {
 		fmt.Printf("Server failed: ", err.Error())
 	}
 }
 ```
 
Создается анонимная структура counter, хранящая счетчик. Она читает шаблон с диска и заполняет в нем {{.Counter}} значением счетчика.
В следующем примере показано, как работают вложенные шаблоны. Имеется базовый шаблон base.html, в который могут быть вложены два других - index.html либо about.html. Исходники можно найти на гитхабе. В одном каталоге с кодом сервера нужно создать подкаталог templates, куда положить 3 шаблона - base.html, index.html, about.html. Код сервера:

```golang
 package main
 
 import (
 	"fmt"
 	"html/template"
 	"io"
 	"log"
 	"net/http"
 	"time"
 )
 
 const STATIC_URL string = "/static/"
 const STATIC_ROOT string = "static/"
 
 type Context struct {
 	Title  string
 	Static string
 }
 
 func Home(w http.ResponseWriter, req *http.Request) {
 	context := Context{Title: "Welcome!"}
 	render(w, "index", context)
 }
 
 func About(w http.ResponseWriter, req *http.Request) {
 	context := Context{Title: "About"}
 	render(w, "about", context)
 }
 
 func render(w http.ResponseWriter, tmpl string, context Context) {
 	context.Static = STATIC_URL
 	tmpl_list := []string{"templates/base.html",
 		fmt.Sprintf("templates/%s.html", tmpl)}
 	t, err := template.ParseFiles(tmpl_list...)
 	if err != nil {
 		log.Print("template parsing error: ", err)
 	}
 	err = t.Execute(w, context)
 	if err != nil {
 		log.Print("template executing error: ", err)
 	}
 }
 
 func StaticHandler(w http.ResponseWriter, req *http.Request) {
 	static_file := req.URL.Path[len(STATIC_URL):]
 	if len(static_file) != 0 {
 		f, err := http.Dir(STATIC_ROOT).Open(static_file)
 		if err == nil {
 			content := io.ReadSeeker(f)
 			http.ServeContent(w, req, static_file, time.Now(), content)
 			return
 		}
 	}
 	http.NotFound(w, req)
 }
 
 func main() {
 	http.HandleFunc("/", Home)
 	http.HandleFunc("/about/", About)
 	http.HandleFunc(STATIC_URL, StaticHandler)
 	err := http.ListenAndServe(":8000", nil)
 	if err != nil {
 		log.Fatal("ListenAndServe: ", err)
 	}
 }
 
 ```
 
Любой веб-сервер должен уметь заполнять при пост-бэке формы ее заполненные вручную поля. В go для этого можно использовать словари:

```golang
 package main
 
 import (
 	"html/template"
 	"net/http"
 )
 
 var templateString = `
 
 <html>
 <body>
 {{ if .name }}
 <p>Your name: {{ .name }}</p>
 {{ end }}
 <form action="/" method="POST">
 <input type="text" name="name" value="{{ .name }}">
 <input type="submit" value="Send">
 </form>
 </body>
 </html>
 `
 var templ = template.Must(template.New("t1").Parse(templateString))
 
 func myFunc(w http.ResponseWriter, r *http.Request) {
 	context := make(map[string]string)
 	if r.Method == "POST" {
 		context["name"] = r.FormValue("name")
 	}
 	templ.Execute(w, context)
 }
 
 func main() {
 	myHandler := http.HandlerFunc(myFunc)
 	http.ListenAndServe(":8000", myHandler)
 }
 
 ```golang
 
### Если нам нужно вывести контент в формате XML:

```golang
 package main
 
 import (
   "encoding/xml"
   "net/http"
 )
 
 type Profile struct {
   Name    string
   Hobbies []string `xml:"Hobbies>Hobby"`
 }
 
 func main() {
   http.HandleFunc("/", foo)
   http.ListenAndServe(":3000", nil)
 }
 
 func foo(w http.ResponseWriter, r *http.Request) {
   profile := Profile{"Alex", []string{"snowboarding", "programming"}}
 
   x, err := xml.MarshalIndent(profile, "", "  ")
   if err != nil {
     http.Error(w, err.Error(), http.StatusInternalServerError)
     return
   }
 
   w.Header().Set("Content-Type", "application/xml")
   w.Write(x)
 }
 ```
 
### В следующем примере показано, как записать и прочитать куку:

```golang
 package main
 
 import (
 	"fmt"
 	"strconv"
 	"log"
 	"net/http"
 )
 
 func SetMyCookie(response http.ResponseWriter){
 	cookie := http.Cookie{Name: "testcookiename", Value:"testcookievalue"}
 	http.SetCookie(response, &cookie)
 }
 
 func rootHandler(response http.ResponseWriter, request *http.Request){
 
 	SetMyCookie(response)
 	response.Header().Set("Content-type", "text/plain")
 	fmt.Fprint(response,  "FooWebHandler says ... \n")
 	fmt.Fprintf(response, " request.Method     '%v'\n", request.Method)
 	fmt.Fprintf(response, " request.RequestURI '%v'\n", request.RequestURI)
 	fmt.Fprintf(response, " request.URL.Path   '%v'\n", request.URL.Path)
 	fmt.Fprintf(response, " request.Form       '%v'\n", request.Form)
 	fmt.Fprintf(response, " request.Cookies()  '%v'\n", request.Cookies())
 }
 ```
 
 ```golang
 
 func main(){
 	port := 8000
 	portstring := strconv.Itoa(port)
 
 	mux := http.NewServeMux()
 	mux.Handle("/", http.HandlerFunc( rootHandler ))
 
 	log.Print("Listening on port " + portstring + " ... ")
 	err := http.ListenAndServe(":" + portstring, mux)
 	if err != nil {
 		log.Fatal("ListenAndServe error: ", err)
 	}
 }
 ```
 
## Сделать upload файла:

```golang
 package main
 
 import (
 	"fmt"
 	"html/template"
 	"io/ioutil"
 	"net/http"
 	"os"
 )
 
 var size int64 = 5 * 1024 * 1024
 var html = template.Must(template.New("html").Parse(`
 <html>
 	<head>
 		<meta charset="UTF-8"/>
 		<title>Golang File Upload</title>
 	</head>
 	<body>
 		<form action="/upload" method="POST" enctype="multipart/form-data">
 			<label for="file">File: </label>
 			<input name="file" type="file"></input>
 			<button type="submit">upload</button>
 		</form>
 	</body>
 </html>
 `))
 
 func root(w http.ResponseWriter, r *http.Request) {
 	err := html.Execute(w, nil)
 	if err != nil {
 		fmt.Print(err)
 	}
 }
 
 func upload(w http.ResponseWriter, r *http.Request) {
 	var path string
 	if err := r.ParseMultipartForm(size); err != nil {
 		fmt.Println(err)
 		http.Error(w, err.Error(), http.StatusForbidden)
 	}
 
 	for _, fileHeaders := range r.MultipartForm.File {
 		for _, fileHeader := range fileHeaders {
 			file, _ := fileHeader.Open()
 			path = fmt.Sprintf("%s", fileHeader.Filename)
 			buf, _ := ioutil.ReadAll(file)
 			ioutil.WriteFile(path, buf, os.ModePerm)
 		}
 	}
 	fmt.Printf("File \"%v\" uploaded\n", path)
 }
 ```
 
 
 ```golang
 func main() {
 	http.HandleFunc("/upload", upload)
 	http.HandleFunc("/", root)
 	fmt.Print(http.ListenAndServe(":8000", nil))
 }
 ```
 
В следующем примере дана реализация аутентификации на основе сессии с использованием куки. Сервер будет обслуживать две страницы - корневую по умолчанию, и вторую - internal - куда пользователь попадает после того, как набирает логин и пароль. Для этого примера нужно на гитхабе забрать тулкит под названием горилла. Из этого тулкита понадобится два пакета - mux и securecookie. Их можно установить с помощью команд:
   go get github.com/gorilla/mux
   go get github.com/gorilla/securecookie
после чего дать ссылку на установленные компоненты в виде

 import (
 	"github.com/gorilla/mux"
 	"github.com//gorilla/securecookie"
Можно пойти другим путем: в локальном каталоге, в котором будет лежать текст, приведенный ниже, создать каталог gorilla и скопировать туда два подкаталога mux и securecookie вместе с файлами. 
В примере регистрируются 4 хэндлера. Для логина и лог-аута разрешен пост. Функция loginHandler читает логин и пароль, которые приходят из формы. Имя пишется в сессию и происходит редирект на internal страницу. Если логин и пароль пусты, возвращаемся на главную. logoutHandle удаляет сессию и редиректит на главную. setSession ложит логин и пароль в сессию, зашифрованное значение сессии сохраняется в куку. getUserName читает куку. В примере используется т.н. client-side сессия. Другой вариант хранения сессии - в базе.

```golang
 package main
 
 import (
 	"fmt"
 	"./gorilla/mux"
 	"./gorilla/securecookie"
 	"net/http"
 )
 
 // cookie handling
 
 var cookieHandler = securecookie.New(
 	securecookie.GenerateRandomKey(64),
 	securecookie.GenerateRandomKey(32))
 
 func getUserName(request *http.Request) (userName string) {
 	if cookie, err := request.Cookie("session"); err == nil {
 		cookieValue := make(map[string]string)
 		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
 			userName = cookieValue["name"]
 		}
 	}
 	return userName
 }
 
 func setSession(userName string, response http.ResponseWriter) {
 	value := map[string]string{
 		"name": userName,
 	}
 	if encoded, err := cookieHandler.Encode("session", value); err == nil {
 		cookie := &http.Cookie{
 			Name:  "session",
 			Value: encoded,
 			Path:  "/",
 		}
 		http.SetCookie(response, cookie)
 	}
 }
 
 func clearSession(response http.ResponseWriter) {
 	cookie := &http.Cookie{
 		Name:   "session",
 		Value:  "",
 		Path:   "/",
 		MaxAge: -1,
 	}
 	http.SetCookie(response, cookie)
 }
 
 // login handler
 
 func loginHandler(response http.ResponseWriter, request *http.Request) {
 	name := request.FormValue("name")
 	pass := request.FormValue("password")
 	redirectTarget := "/"
 	if name != "" && pass != "" {
 		// .. check credentials ..
 		setSession(name, response)
 		redirectTarget = "/internal"
 	}
 	http.Redirect(response, request, redirectTarget, 302)
 }
 
 // logout handler
 
 func logoutHandler(response http.ResponseWriter, request *http.Request) {
 	clearSession(response)
 	http.Redirect(response, request, "/", 302)
 }
 
 // index page
 
 const indexPage = `
 <h1>Login</h1>
 <form method="post" action="/login">
     <label for="name">User name</label>
     <input type="text" id="name" name="name">
     <label for="password">Password</label>
     <input type="password" id="password" name="password">
     <button type="submit">Login</button>
 </form>
 `
 
 func indexPageHandler(response http.ResponseWriter, request *http.Request) {
 	fmt.Fprintf(response, indexPage)
 }
 
 // internal page
 
 const internalPage = `
 <h1>Internal</h1>
 <hr>
 <small>User: %s</small>
 <form method="post" action="/logout">
     <button type="submit">Logout</button>
 </form>
 `
 
 func internalPageHandler(response http.ResponseWriter, request *http.Request) {
 	userName := getUserName(request)
 	if userName != "" {
 		fmt.Fprintf(response, internalPage, userName)
 	} else {
 		http.Redirect(response, request, "/", 302)
 	}
 }
 
 // server main method
 
 var router = mux.NewRouter()
 
 func main() {
 
 	router.HandleFunc("/", indexPageHandler)
 	router.HandleFunc("/internal", internalPageHandler)
 
 	router.HandleFunc("/login", loginHandler).Methods("POST")
 	router.HandleFunc("/logout", logoutHandler).Methods("POST")
 
 	http.Handle("/", router)
 	http.ListenAndServe(":8000", nil)
 }
 
 ```
 
В следующем примере мы напишем веб-сервис и клиента к нему. Материал взят отсюда. 
Прежде всего нам понадобится пакет net/http, поскольку наш сервис будет работать по протоколу http. Также нам понадобится пакет мартини. Сервис будет выполнять основные прототипы функций для интерфейса гостевой книги: он позволяет добавлять записи в гостевую книгу и читать их после добавления. Записи хранятся в памяти. Запись в гостевой книге представлена структурой:

```golang
 type GuestBookEntry struct {
         Id      int
         Email   string
         Title   string
         Content string
 }
 ```
 
Сама гостевая книга представлена структурой, реализованной в форме коллекции:

```
 type GuestBook struct {
         guestBookData []*GuestBookEntry
 }
 ```
 
Добавление записи в гостевую книгу:
```golang
 func (g *GuestBook) AddEntry(email, title, content string) int
 ```
 
Прочитать запись:
```golang
 func (g *GuestBook) GetEntry(id int) (*GuestBookEntry, error)
 ```
 
Как вы уже знаете, в go методу можно передать в качестве параметра интерфейс. В этом интерфейсе можно реализовать целый список других методов и произвольный набор типов. Реализация интерфейса WebService, в котором определены 4 метода, в частности методы для добавления, чтения и удаления записи в гостевой книге:

```golang
 type WebService interface {
         GetPath() string
  
         // если параметр отсутсвует, удяляются все записи
         WebDelete(params martini.Params) (int, string)
  
         // реализация http-метода GET 
         // если параметр отсутствует, возвращаются все записи
         WebGet(params martini.Params) (int, string)
  
         // реализация http-метода POST method. 
         WebPost(params martini.Params, req *http.Request) (int, string)
 
 }
 ```
 
Функция, регистрирующая вебсервис, в котором инициализируются эти 4 метода:

   func RegisterWebService(webService WebService, classicMartini *martini.ClassicMartini)
Реализация http.Post:

```golang
 func (g *GuestBook) WebPost(params martini.Params,
         req *http.Request) (int, string) {
         defer req.Body.Close()
  
         // читаем тело запроса
         requestBody, err := ioutil.ReadAll(req.Body)
         if err != nil {
                 return http.StatusInternalServerError, “internal error”
         }
  
         if len(params) != 0 {
                 return http.StatusMethodNotAllowed, “method not allowed”
         }
  
         // расшифровываем данные от клиента
         var guestBookEntry GuestBookEntry
         err = json.Unmarshal(requestBody, &guestBookEntry)
         if err != nil {
                 return http.StatusBadRequest, “invalid JSON data”
         }
  
         // добавляем запись
         g.AddEntry(guestBookEntry.Email, guestBookEntry.Title,
                 guestBookEntry.Content)
  
         return http.StatusOK, “new entry created”
 }
 ```
 
Стандартный пакет encoding/json переводит присланные данные из формата json в структуры go. 
Код главной функции на сервере:

```golang
 func main() {
 	martiniClassic := martini.Classic()
 	guestBook := guestbook.NewGuestBook()
 	guestbook.RegisterWebService(guestBook, martiniClassic)
 	martiniClassic.Run()
 }
 ```
 
#### Здесь мы создаем 2 обьекта - мартини и гостевую книгу - и передаем эти обьекты в качестве параметров для регистрации сервиса. Исходные коды примера лежат тут . Распаковав архив, в корневом каталоге собираем сервер и клиента:
 go build server.go
 go build client.go

#### Запускаем сервер:
  server
 
#### Открываем второй терминал и запускаем клиента со следующими параметрами, выполняя пост:
 client --request_url=http://127.0.0.1:8000/guestbook --request_method=post \
        --request_data='{"Id":0,"Email":"my-email@blablabla.com"}'

#### Читаем добавленную запись:
 client --request_url=http://127.0.0.1:8000/guestbook/0 --request_method=get

#### Ее также можно посмотреть в броузере:
 http://localhost:8000/guestbook/0
 
 
