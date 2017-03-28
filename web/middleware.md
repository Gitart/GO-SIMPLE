## Middleware

Middleware функция которая должна выполниться всегда со многими другими функциями.    
Для этого необходимо создать функцию котрая на вход принимает имя второй функции   
и после выполнения своих инструкций передает управление внутренней функции.    
[Основано на источнике](http://www.alexedwards.net/blog/making-and-using-middleware)

### Сервер с использованием middleware

```golang
package main

import (
  "log"
  "net/http"
)

/*
func middlewareOne(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    log.Println("Executing middlewareOne")
    next.ServeHTTP(w, r)
    log.Println("Executing middlewareOne again")
  })
}

func middlewareTwo(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    log.Println("Executing middlewareTwo")
    
    if r.URL.Path != "/" {
       return
    }
    next.ServeHTTP(w, r)
    log.Println("Executing middlewareTwo again")
  })
}
*/

// Middleware
func middleware(next http.Handler) http.Handler {
	 return http.HandlerFunc(
	 func  (w http.ResponseWriter, r *http.Request) {
            log.Println("Мидлеваре")
            next.ServeHTTP(w, r)
      })
}

// Функция первая
func final(w http.ResponseWriter, r *http.Request) {
     log.Println("Выполнение процедуры")
     w.Write([]byte("OK"))
}

// Функция вторая
func finalT(w http.ResponseWriter, r *http.Request) {
     log.Println("Вы зашли в другую процедуру после мидлеваре.")
     w.Write([]byte("OK"))
}

func main() {
   // finalHandler := http.HandlerFunc(final)
  
  // Если необходимо выполнить вложенные мидлеваре
  // http.Handle("/", middlewareOne(middlewareTwo(finalHandler)))

  http.Handle("/",     middleware(http.HandlerFunc(final)))
  http.Handle("/test", middleware(http.HandlerFunc(finalT)))

  log.Println("Старт сервиса мидлеваре.")
  http.ListenAndServe(":3000", nil)
}
```

## Test

```bat
c:\curl\curl.exe -XPOST  -i localhost:3000
c:\curl\curl.exe -XGET   -i localhost:3000
c:\curl\curl.exe -XPUT   -i localhost:3000
c:\curl\curl.exe -XPUT   -i localhost:3000/test
c:\curl\curl.exe -XPOST  -i localhost:3000/test
```


## Второй вариант

```golang
package main

import (
  "github.com/gorilla/handlers"
  "net/http"
  "os"
)

func main() {
  finalHandler := http.HandlerFunc(final)

  logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
  if err != nil {
    panic(err)
  }

  http.Handle("/", handlers.LoggingHandler(logFile, finalHandler))
  http.ListenAndServe(":3000", nil)
}

func final(w http.ResponseWriter, r *http.Request) {
     w.Write([]byte("OK"))
}
```

### Output
#### server.log
```
127.0.0.1 - - [28/Mar/2017:10:47:19 +0300] "POST / HTTP/1.1" 200 2
127.0.0.1 - - [28/Mar/2017:10:47:19 +0300] "GET / HTTP/1.1" 200 2
127.0.0.1 - - [28/Mar/2017:10:47:20 +0300] "PUT / HTTP/1.1" 200 2
```
