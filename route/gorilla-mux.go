package main
import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)
 
func productsHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    cat := vars["category"]
    response := fmt.Sprintf("Product category=%s id=%s", cat, id)
    fmt.Fprint(w, response)
}
 
func main() {
      
    router := mux.NewRouter()
    router.HandleFunc("/products/{category}/{id:[0-9]+}", productsHandler)
    http.Handle("/",router)
 
    fmt.Println("Server is listening...")
    http.ListenAndServe(":8181", nil)
}
