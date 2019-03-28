package main
import (
    "fmt"
    "net/http"
)
 
type httpHandler struct{
    message string
}
func (h httpHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
    fmt.Fprint(resp, h.message) 
 }
 
func main() {
      
    h1 := httpHandler{ message:"Index"}
    h2 := httpHandler{ message:"About"}
 
    http.Handle("/", h1)
    http.Handle("/about", h2)
 
    fmt.Println("Server is listening...")
    http.ListenAndServe(":8181", nil)
}
