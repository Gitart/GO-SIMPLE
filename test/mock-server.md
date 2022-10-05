## Golang реализует HTTP Mock Server

Теги:  [mock](https://russianblogs.com/tag/mock/ "mock")  [httpserver](https://russianblogs.com/tag/httpserver/ "httpserver")

Исходный код от[https://github.com/deis/mock-http-server](https://github.com/deis/mock-http-server)

Основная роль - открыть HTTP-сервер локального порта прослушивателя 8080, который может распечатать запрос клиента, который удобен для отладки.

```Go
   package main
 
import (
	"fmt"
	"log"
	"net/http"
)
 
// Log the HTTP request
func logHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
}
 
// mockHandler responds with "ok" as the response body
func mockHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok\n")
}
 
// rootHandler used to process all inbound HTTP requests
func rootHandler(w http.ResponseWriter, r *http.Request) {
	logHandler(w, r)
	mockHandler(w, r)
}
 
// Start an HTTP server which dispatches to the rootHandler
func main() {
	http.HandleFunc("/", rootHandler)
	port := "8080"
	log.Printf("server is listening on %v\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
```

Если вам нужно печатать заголовок, пожалуйста, обратитесь к:[https://www.cnblogs.com/5bug/p/8494953.html](https://www.cnblogs.com/5bug/p/8494953.html)

```Go
func helloFunc(w http.ResponseWriter, r *http.Request)  {
   FMT.Println («Список параметров заголовков печати:»)
   if len(r.Header) > 0 {
      for k,v := range r.Header {
         fmt.Printf("%s=%s\n", k, v[0])
      }
   }
       FMT.PrintLn («Список параметров формы печати:»)
   r.ParseForm()
   if len(r.Form) > 0 {
      for k,v := range r.Form {
         fmt.Printf("%s=%s\n", k, v[0])
      }
   }
```
