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

