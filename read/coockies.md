## Coockies
Now, to see how the cookie is being created, read and deleted. You need to follow this sequence.

Run the program and point your web browser to localhost:8080
Change the URL to localhost:8080/createcookie to create new cookie.
Read the cookie by changing to localhost:8080/createcookie.
Finally, delete the cookie by changing to localhost:8080/deletecookie.
If you happen to view with Chrome browser, open up the inspector to see the cookie value

```go
package main

 import (
         "net/http"
 )

 func SayHelloWorld(w http.ResponseWriter, r *http.Request) {
         html := "Hello World! "
         w.Write([]byte(html))
 }

 func ReadCookie(w http.ResponseWriter, r *http.Request) {
         c, err := r.Cookie("ithinkidroppedacookie")
         if err != nil {
                 w.Write([]byte("error in reading cookie : " + err.Error() + "\n"))
         } else {
                 value := c.Value
                 w.Write([]byte("cookie has : " + value + "\n"))
         }
 }

 // see https://golang.org/pkg/net/http/#Cookie
 // Setting MaxAge<0 means delete cookie now.

 func DeleteCookie(w http.ResponseWriter, r *http.Request) {
         c := http.Cookie{
                 Name:   "ithinkidroppedacookie",
                 MaxAge: -1}
         http.SetCookie(w, &c)

         w.Write([]byte("old cookie deleted!\n"))
 }

 func CreateCookie(w http.ResponseWriter, r *http.Request) {
         c := http.Cookie{
                 Name:   "ithinkidroppedacookie",
                 Value:  "thedroppedcookiehasgoldinit",
                 MaxAge: 3600}
         http.SetCookie(w, &c)

         w.Write([]byte("new cookie created!\n"))
 }

 func main() {
         mux := http.NewServeMux()
         mux.HandleFunc("/", SayHelloWorld)
         mux.HandleFunc("/readcookie", ReadCookie)
         mux.HandleFunc("/deletecookie", DeleteCookie)
         mux.HandleFunc("/createcookie", CreateCookie)
         http.ListenAndServe(":8080", mux)
 }
```
