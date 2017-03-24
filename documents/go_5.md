# Веб-программирование в Go
## Часть 2

Рассмотрим простейшее веб-приложение:

```golang
 package main
     
     import (
         "fmt"
         "net/http"
     )
     
     func handler(w http.ResponseWriter, r *http.Request) {
         fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
     }
     
     func main() {
         http.HandleFunc("/", handler)
         http.ListenAndServe(":8000", nil)
     }
```

Функция handler импортирует из базового пакета http два обьекта - http.ResponseWriter и указатель на http.Request. Функция http.HandleFunc будет перенаправлять все запросы для корня на хэндлер. Затем мы включаем прослушивание порта 8000, не используя хэндлер. Системный пакет net/http написан все на том же go. Если глянуть на его исходники, то мы увидим, что ResponseWriter - это интерфейс, в котором определены 3 функии:

```golang
     type ResponseWriter interface {
             // заголовок возвращает коллекцию типа map
             Header() Header
     
             // пишет данные для клиента после вызова WriteHeader()
             Write([]byte) (int, error)
     
             // отсылает клиенту заголовок вместе с кодом статуса
             WriteHeader(int)
     }
```

Хэндлер можно было бы расшифровать и переписать на более низком уровне так:

```golang
 func handler(w http.ResponseWriter, r *http.Request) {
         w.Header().Set("Content-Type", "application/json; charset=utf-8") 
 
         myItems := []string{"item1", "item2", "item3"}
         a, _ := json.Marshal(myItems)
 
         w.Write(a)
         return
     }
```

Иногда возникает необходимость после того, как респонс отдан клиенту, дописать в конец этого респонса что-то еще. Для этого нужно вызвать функцию Write() этого самого респонса. Нужно создать структуру, в которой будет одно поле - хэндлер, создать ссылку на обьект этой структуры, присвоив ее хэндлеру нужное значение, и эту ссылку передать в нужный обработчик:

```golang
  package main
 
 import (
 	"net/http"
 )
 
 type AppendMiddleware struct {
 	handler http.Handler
 }
 
 func (a *AppendMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
 	a.handler.ServeHTTP(w, r)
 	w.Write([]byte(""))
 }
 
 func rootHandler(w http.ResponseWriter, r *http.Request) {
 	w.Write([]byte("Success!"))
 }
 
 func aboutHandler(w http.ResponseWriter, r *http.Request) {
 	w.Write([]byte("About !"))
 }
 
 func main() {
 
 	rmd := &AppendMiddleware{http.HandlerFunc(rootHandler)}
 	amd := &AppendMiddleware{http.HandlerFunc(aboutHandler)}
 
         http.Handle("/", rmd)
         http.Handle("/about/", amd)
 	http.ListenAndServe(":8000", nil)
 }
 ```
 
А что делать, если нужно что-то добавить не в конец, а в начало респонса ? Нам нужен буфер респонса, который можно модифицировать и потом передать в хэндлер. Функция ResponseRecorder хранит этот буфер, и кроме этого хранит заголовки:

```golang
 package main
 
 import (
 	"net/http"
 	"net/http/httptest"
 	"strconv"
 )
 
 type ModifierMiddleware struct {
 	handler http.Handler
 }
 
 func (m *ModifierMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
 	rec := httptest.NewRecorder()
 	// подменяем респонс на ResponseRecorder
 	m.handler.ServeHTTP(rec, r)
 
 	// копируем хидер респонса
 	for k, v := range rec.Header() {
 		w.Header()[k] = v
 	}
 	// добавляем свой собственный хидер
 	w.Header().Set("X-We-Modified-This", "Yup")
 	// status code
 	w.WriteHeader(418)
         // вставляем в начало респонса нужный текст
         data := []byte("Middleware says hello again. ")
 
         // у заголовка изменился Content-Length
         // нужно его пересчитать
         clen, _ := strconv.Atoi(r.Header.Get("Content-Length"))
         clen += len(data)
         r.Header.Set("Content-Length", strconv.Itoa(clen))
 
         // пишем то, что хотели добавить в начало
         w.Write(data)
 	// пишем все остальное
 	w.Write(rec.Body.Bytes())
 }
 
 func myHandler(w http.ResponseWriter, r *http.Request) {
 	w.Write([]byte("Success!"))
 }
 
 func main() {
 	mid := &ModifierMiddleware{http.HandlerFunc(myHandler)}
 
 	println("Listening on port 8000")
 	http.ListenAndServe(":8000", mid)
 }
 ```
 
 
Обьект http.Request представляет из себя структуру, в которой содержатся параметры клиентского запроса, данные и т.д:

```golang
    type Request struct {
             Method string
             URL *url.URL
             Proto      string // "HTTP/1.0"
             ProtoMajor int    // 1
             ProtoMinor int    // 0
             Header Header
             Body io.ReadCloser
             ContentLength int64
             TransferEncoding []string
             Close bool
             Host string
             Form url.Values
             PostForm url.Values
             MultipartForm *multipart.Form
             Trailer Header
             RemoteAddr string
             RequestURI string
             TLS *tls.ConnectionState
     }
```

Системная функция http.HandleFunc регистрирует хэндлер:

```golang
    func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
         DefaultServeMux.HandleFunc(pattern, handler)
     }
```     

* ServeMux - это http-мультиплексор или переключатель. Он берет урл из каждого запроса и вызывает тот хэндлер, который ему соответствует. Урл может быть простым и составным. 

Исходный код функции http.ListenAndServe - она устанавливает tcp-коннект с keep-alive таймаутом, создавая для каждого входящего коннекта новую goroutine:

```golang
     func ListenAndServe(addr string, handler Handler) error {
         server := &Server{Addr: addr, Handler: handler}
         return server.ListenAndServe()
     }
     
    func (srv *Server) ListenAndServe() error {
         addr := srv.Addr
         if addr == "" {
             addr = ":http"
         }
         ln, err := net.Listen("tcp", addr)
         if err != nil {
             return err
         }
         return srv.Serve(tcpKeepAliveListener{ln.(*net.TCPListener)})
     }
     
    func (srv *Server) Serve(l net.Listener) error {
         defer l.Close()
         var tempDelay time.Duration // how long to sleep on accept failure
         for {
             rw, e := l.Accept()
             if e != nil {
                 if ne, ok := e.(net.Error); ok && ne.Temporary() {
                     if tempDelay == 0 {
                         tempDelay = 5 * time.Millisecond
                     } else {
                         tempDelay *= 2
                     }
                     if max := 1 * time.Second; tempDelay > max {
                         tempDelay = max
                     }
                     srv.logf("http: Accept error: %v; retrying in %v", e, tempDelay)
                     time.Sleep(tempDelay)
                     continue
                 }
                 return e
             }
             tempDelay = 0
             c, err := srv.newConn(rw)
             if err != nil {
                 continue
             }
             c.setState(c.rwc, StateNew) // before Serve can return
             go c.serve()
         }
     }
```

Пакет http имеет несколько встроенных хэндлеров, таких, как FileServer, NotFoundHandler, RedirectHandler. Последний можно использовать для редиректа - в следующем примере мы создаем два новых обьекта - мультиплексор и редирект-хэндлер, регистрируем мультиплексор и передаем его в качестве параметра - в результате при загрузке корневой страницы сразу произойдет редирект:

```golang
   mux := http.NewServeMux()
 
   rh := http.RedirectHandler("http://example.org", 307)
   mux.Handle("/", rh)
 
   log.Println("Listening...")
   http.ListenAndServe(":3000", mux)
````

В следующем примере мы рассмотрим, как делать вложенные хэндлеры. Пусть у нас имеются два хэндлера, в каждом засекается время его выполнения, после чего это время логируется:

```golang
 func aboutHandler(w http.ResponseWriter, r *http.Request) {
   t1 := time.Now()
   fmt.Fprintf(w, "You are on the about page.")
   t2 := time.Now()
   log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
 }
 
 func indexHandler(w http.ResponseWriter, r *http.Request) {
   t1 := time.Now()
   fmt.Fprintf(w, "Welcome!")
   t2 := time.Now()
   log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
 }
 
 func main() {
   http.HandleFunc("/about", aboutHandler)
   http.HandleFunc("/", indexHandler)
   http.ListenAndServe(":8080", nil)
 }
 ```
 
Здесь мы видим повторное использование кода, от которого надо избавиться. Надо написать хэндлер, который будет в качестве параметра принимать другой хэндлер, чтобы это выглядело как-то так:

```
 loggingHandler(indexHandler)
 ```
 
Для функции логирования мы создадим отдельный хэндлер, который в качестве параметра принимает другой хэндлер:

```golang
 func loggingHandler(next http.Handler) http.Handler {
   fn := func(w http.ResponseWriter, r *http.Request) {
     t1 := time.Now()
     next.ServeHTTP(w, r)
     t2 := time.Now()
     log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
   }
 
   return http.HandlerFunc(fn)
 }
 ```
 
Окончательно программа выглядит так:

```golang
 package main
 
 import (
   
   "net/http"
   "time"
   "log"
   "fmt"
   
 )  
 
 type Constructor func(http.Handler) http.Handler
 
 // Chain - immutable коллекция хэндлеров
 type Chain struct {
 	constructors []Constructor
 }
 
 // создает коллекцию Chain
 func New(constructors ...Constructor) Chain {
 	c := Chain{}
 	c.constructors = append(c.constructors, constructors...)
 
 	return c
 }
 
 // функция возвращает из коллекции нужный хэндлер
 // берет в качестве параметра хэндлер
 // может вызываться несколько раз подряд, 
 // т.е. уровень вложенности хэндлеров может быть больше двух
 func (c Chain) Then(h http.Handler) http.Handler {
 	var final http.Handler
 	if h != nil {
 		final = h
 	} else {
 		final = http.DefaultServeMux
 	}
 
 	for i := len(c.constructors) - 1; i >= 0; i-- {
 		final = c.constructors[i](final)
 	}
 
 	return final
 }
 
 // эта функция является оберткой для предыдущей функции 
 // берет в качестве параметра хэндлер-функцию
 func (c Chain) ThenFunc(fn http.HandlerFunc) http.Handler {
 	if fn == nil {
 		return c.Then(nil)
 	}
 	return c.Then(http.HandlerFunc(fn))
 }
 
 
 func loggingHandler(next http.Handler) http.Handler {
   fn := func(w http.ResponseWriter, r *http.Request) {
     t1 := time.Now()
     next.ServeHTTP(w, r)
     t2 := time.Now()
     log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
   }
 
   return http.HandlerFunc(fn)
 }
 
 func aboutHandler(w http.ResponseWriter, r *http.Request) {
   fmt.Fprintf(w, "You are on the about page.")
 }
 
 func indexHandler(w http.ResponseWriter, r *http.Request) {
   fmt.Fprintf(w, "Welcome!")
 }
 
 func main() {
   commonHandlers := New(loggingHandler)
   http.Handle("/about/", commonHandlers.ThenFunc(aboutHandler))
   http.Handle("/", commonHandlers.ThenFunc(indexHandler))
   http.ListenAndServe(":8000", nil)
 }
 ```
 
Загружаем веб-сервер и смотрим лог, при этом все логируется:

```
 2015/01/06 20:13:37 [GET] "/" 39.091µs
 2015/01/06 20:13:48 [GET] "/about/" 9.115µs
```

Добавим в этот пример еще один обработчик верхнего уровня, который будет обрабатывать ошибки хэндлеров и поддерживать сервер на плаву:

```golang
 func recoverHandler(next http.Handler) http.Handler {
     fn := func(w http.ResponseWriter, r *http.Request) {
     defer func() {
       if err := recover(); err != nil {
         log.Printf("panic: %+v", err)
         http.Error(w, http.StatusText(500), 500)
       }
     }()
 
     next.ServeHTTP(w, r)
   }
 
   return http.HandlerFunc(fn)
 }
 ```
 
Главная функция будет выглядеть так:

```golang
 func main() {
   commonHandlers := New(loggingHandler, recoverHandler)
   http.Handle("/about/", commonHandlers.ThenFunc(aboutHandler))
   http.Handle("/", commonHandlers.ThenFunc(indexHandler))
   http.ListenAndServe(":8000", nil)
 }
```

Если теперь в каком-то хэндлере случится непредвиденная серверная ошибка,
она будет обработана новым хэндлером и сервер не упадет, а будет работать дальше, только в логе появится сообщение.
