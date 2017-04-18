## Nested template

```golang
package main
import (
"html/template"
"net/http"
)

var t *template.Template
func init() {
t = template.Must(template.ParseFiles("index.html", "head.html"))
}

type Page struct {
Title, Content string
}

func diaplayPage(w http.ResponseWriter, r *http.Request) {
p := &Page{
Title: "An Example",
Content: "Have fun stormin’ da castle.",
}

t.ExecuteTemplate(w, "index.html", p)
}

func main() {
http.HandleFunc("/", diaplayPage)
http.ListenAndServe(":8080", nil)
}
```

## index.html
```html
<!DOCTYPE HTML>
<html>
{{template "head.html" .}}
<body>
<h1>{{.Title}}</h1>
<p>{{.Content}}</p>
</body>
</html>
```


## Head.html
```html
<head>
<meta charset="utf-8">
<title>{{.Title}}</title>
</head>
```


## Base.html

```html
{{define "base"}}<!DOCTYPE HTML>

<html>
<head>
<meta charset="utf-8">
<title>{{template "title" .}}</title>
{{ block "styles" . }}<style>
h1 {
color: #400080
}
</style>{{ end }}
</head>

<body>
<h1>{{template "title" .}}</h1>
{{template "content" .}}
{{block "scripts" .}}{{end}}
</body>

</html>{{end}}
```

## User.html
```html
{{define "title"}}User: {{.Username}}{{end}}
{{define "content"}}
<ul>
  <li>Userame: {{.Username}}</li>
  <li>Name: {{.Name}}</li>
</ul>
{{end}}
```


## Page.html
```html
{{define "title"}}{{.Title}}{{end}}
{{define "content"}}
<p>
{{.Content}}
</p>
{{end}}
{{define "styles"}}
<style>
h1 {
color: #800080
}
</style>
{{end}}
```

## Inherit go
```golang

package main

import (

"html/template"

"net/http"

)

var t map[string]*template.Template

func init() {
  t = make(map[string]*template.Template)
  temp := template.Must(template.ParseFiles("base.html", "user.html"))
  t["user.html"] = temp
  temp = template.Must(template.ParseFiles("base.html", "page.html"))
  t["page.html"] = temp
}

type Page struct {
  Title, Content string
}

type User struct {
  Username, Name string
}

func displayPage(w http.ResponseWriter, r *http.Request) {
  p := &Page{
  Title: "An Example",
  Content: "Have fun stormin’ da castle.",
  }
  t["page.html"].ExecuteTemplate(w, "base", p)
}


func displayUser(w http.ResponseWriter, r *http.Request) {
  u := &User{
  Username: "swordsmith",
  Name: "Inigo Montoya",
  }
  t["user.html"].ExecuteTemplate(w, "base", u)
}

func main() {
  http.HandleFunc("/user", displayUser)
  http.HandleFunc("/", displayPage)
  http.ListenAndServe(":8080", nil)
}
```
