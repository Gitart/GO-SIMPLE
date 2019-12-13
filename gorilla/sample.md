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
