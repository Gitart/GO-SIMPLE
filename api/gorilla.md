## Gorila example

```go
package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func allUsers(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "All Users Endpoint Hit")
}

func newUser(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "New User Endpoint Hit")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Delete User Endpoint Hit")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Update User Endpoint Hit")
}

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/users", allUsers).Methods("GET")
    myRouter.HandleFunc("/user/{name}", deleteUser).Methods("DELETE")
    myRouter.HandleFunc("/user/{name}/{email}", updateUser).Methods("PUT")
    myRouter.HandleFunc("/user/{name}/{email}", newUser).Methods("POST")
    log.Fatal(http.ListenAndServe(":8081", myRouter))
}


func main() {
    fmt.Println("Go ORM Tutorial")

    // Handle Subsequent requests
    handleRequests()
}
```
