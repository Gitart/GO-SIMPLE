
```go
func readForm(r *http.Request) *User {
    r.ParseForm()
    user := new(User)
    decoder := schema.NewDecoder()
    decodeErr := decoder.Decode(user, r.PostForm)
    if decodeErr != nil {
    log.Printf("error mapping parsed form data to struct : ", decodeErr)
 }
return user
}
```


```go
func login(w http.ResponseWriter, r *http.Request){
if r.Method == "GET"{
  parsedTemplate, _ := template.ParseFiles("templates/
  login-form.html")
  parsedTemplate.Execute(w, nil)
}
else
{
user := readForm(r)
fmt.Fprintf(w, "Hello "+user.Username+"!")
}
}
```

