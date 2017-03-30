## Web routing/multiplex example
Routing based on the URL's path can be useful in some cases like build RESTful API server.

### Problem :

You need to route/multiplex to different function/handler based on the URL's path. For example :    
 "/someresource/:id" ---> "code to do something with the resource"    
 "/users/:name/profile" ---> "code to do something with the profile"    

Solution :

Use "net/http" NewServeMux() function. It compares incoming requests against a list of predefined URL paths,    
and calls the associated handler for the path whenever a match is found.   
Code example ;

```golang
 package main

 import (
         "net/http"
         "strings"
 )

 func SayHelloWorld(w http.ResponseWriter, r *http.Request) {
         w.Write([]byte("Hello, World!"))
 }

 func ReplyName(w http.ResponseWriter, r *http.Request) {
         URISegments := strings.Split(r.URL.Path, "/")
         w.Write([]byte(URISegments[1]))
 }

 func main() {
         // http.Handler
         mux := http.NewServeMux()
         mux.HandleFunc("/", SayHelloWorld)
         mux.HandleFunc("/replyname", ReplyName)

         http.ListenAndServe(":8080", mux)
 }
 ```
 
there are couple of third parties packages that provides more features when come to routing. For example, 
Gorilla Mux for path pattern matching (useful for RESTful APIs)
In this example, we will use Gorilla's Mux. Go get from github.com/gorilla/mux before trying out the codes below.

```
 package main

 import (
         "fmt"
         "github.com/gorilla/mux"
         "net/http"
 )

 func SayHelloWorld(w http.ResponseWriter, r *http.Request) {
         w.Write([]byte("Hello, World!"))
 }

 func ReplyNameGorilla(w http.ResponseWriter, r *http.Request) {
         name := mux.Vars(r)["name"] // variable name is case sensitive
         w.Write([]byte(fmt.Sprintf("Hello %s !", name)))
 }

 func main() {
         mx := mux.NewRouter()
         mx.HandleFunc("/", SayHelloWorld)
         mx.HandleFunc("/{name}", ReplyNameGorilla)  // variable name is case sensitive

         http.ListenAndServe(":8080", mx)
 }
 ```
