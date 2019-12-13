## Work wit Gorilla

```golang
main.go

package main

import (

  "route/page"
  "github.com/gorilla/mux"
  "log"
  "net/http"
)

func main() {

    router := mux.NewRouter()
    router.HandleFunc("/page", page.Search).Methods("GET","OPTIONS")
    log.Fatal(http.ListenAndServe(":8000", router))

}


//route/page.go

package page
import (
  "net/http"
  "log"
)

func Search(w http.ResponseWriter, r *http.Request) {

  w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, Organization")
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Content-Type", "application/json")

  log.Println(r.Header.Get("authorization"))

  log.Println("Hello World")
  w.Write([]byte("{}"))
  return
}
```

## Tes sample
```golang
// main.go
package main

import (

  "route/page"
  "github.com/gorilla/mux"
  "github.com/gorilla/handlers"
  "log"
  "net/http"
)

func main() {

    router := mux.NewRouter()

    router.HandleFunc("/page", page.Search).Methods("GET")

    headersOk := handlers.AllowedHeaders([]string{"Content-Type","Authorization","Organization"})
    originsOk := handlers.AllowedOrigins([]string{"*"})
    methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
    log.Fatal(http.ListenAndServe(":8000", handlers.CORS(originsOk, headersOk, methodsOk)(router)))

}

// route/page.go
package page
import (
  "net/http"
  "log"
)


func Search(w http.ResponseWriter, r *http.Request) {

  log.Println(r.Header.Get("authorization"))

  log.Println("Hello World")
  w.Write([]byte("{}"))
  return
}
```

## Full code
```golang
// main.go

package main

import (

  "route/page"
  "github.com/gorilla/mux"
  "log"
  "net/http"
)

func main() {

    router := mux.NewRouter()
    router.HandleFunc("/page", page.PreFlightHandler).Methods("OPTIONS")

    router.HandleFunc("/page", page.Search).Methods("GET")

    log.Fatal(http.ListenAndServe(":8000", router))

}

// route/page.go
package page
import (
  "net/http"
  "log"
)

func PreFlightHandler(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Vary", "Origin")
    w.Header().Set("Vary", "Access-Control-Request-Method")
    w.Header().Set("Vary", "Access-Control-Request-Headers")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Origin, Accept, Authorization, Organization")
    w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
    w.Header().Set("Access-Control-Allow-Credentials", "true")
}
func Search(w http.ResponseWriter, r *http.Request) {

  log.Println(r.Header.Get("authorization"))

  log.Println("Hello World")
  w.Write([]byte("{}"))
  return
}
```
