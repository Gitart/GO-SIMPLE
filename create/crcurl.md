# Curl usage examples with Golang
Tags : curl golang api-server 
Writing down these examples for my own future references. They are for my own usage and to show how to use curl to test my own Golang programs.

1. curl POST with and without data

Normal POST without data:
```
$ curl -X POST http://localhost:8888/api/login
```

Normal POST with data:

```
$ curl -d "username=adamng&password=abc123" http://localhost:8888/api/login/
```

The Golang API server handling POST data source code:

```go
 package main

 import (
         "fmt"
         "net/http"
 )

 func Home(w http.ResponseWriter, r *http.Request) {
         w.Write([]byte("use curl command!"))
 }

 func HandleLogin(w http.ResponseWriter, r *http.Request) {
         // get the POST data
         username := r.PostFormValue("username")
         password := r.PostFormValue("password")

         received := username + " " + password
         fmt.Println(received)
         w.Write([]byte(received))
 }

 func main() {
         http.HandleFunc("/", Home)
         http.HandleFunc("/api/login", HandleLogin)
         http.ListenAndServe(":8888", nil)
 }
 ```
 
### 2. curl upload file example
To upload file with curl:

```
$ curl -F fileUploadName=@"/path/filename.txt" http://localhost:8888/api/upload
```

NOTE: Pay attention to fileUploadName, it must match the r.FormFile("fileUploadName").
This snippet does not save the uploaded file. See https://www.socketloop.com/tutorials/golang-command-line-file-upload-program-to-server-example for complete program.

```go
 package main

 import (
         "fmt"
         "net/http"
 )

 func Home(w http.ResponseWriter, r *http.Request) {
         w.Write([]byte("use curl command!"))
 }

 func HandleUpload(w http.ResponseWriter, r *http.Request) {

         // For complete program, see
         // https://www.socketloop.com/tutorials/golang-command-line-file-upload-program-to-server-example

         // the FormFile function takes in the POST input id file
         file, header, err := r.FormFile("fileUploadName")

         if err != nil {
                 fmt.Fprintln(w, err)
                 return
         }

         defer file.Close()

         fmt.Println(header.Filename)
         w.Write([]byte(header.Filename))
 }

 func main() {
         http.HandleFunc("/", Home)
         http.HandleFunc("/api/upload", HandleUpload)
         http.ListenAndServe(":8888", nil)
 }
 ```
 
## 3. curl post JSON data example
To POST with JSON data, use this parameter with curl
```
-H "Content-Type: application/json"
```

Example:
```
$ curl -H "Content-Type: application/json" -X POST -d '{"username":"adamng","password":"abc123"}' http://localhost:8888/api/login
```

## The server source code to process JSON POST data by curl:

```go
 package main

 import (
         "fmt"
         "io/ioutil"
         "net/http"
 )

 func Home(w http.ResponseWriter, r *http.Request) {
         w.Write([]byte("use curl command!"))
 }

 func HandleJSON(w http.ResponseWriter, r *http.Request) {
         // get JSON POST data
         json, err := ioutil.ReadAll(r.Body)
         if err != nil {
                 panic(err)
         }
         fmt.Println(string(json))

         // see how to unmarshal json data
         // at https://www.socketloop.com/tutorials/golang-unmarshal-json-from-http-response

         // this will echo back to curl
         //w.Write([]byte(string(json)))
 }

 func main() {
         http.HandleFunc("/", Home)
         http.HandleFunc("/api/login", HandleJSON)
         http.ListenAndServe(":8888", nil)
 }
 ```
 
