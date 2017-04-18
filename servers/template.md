## Simple template

suimple.html
```html
<!DOCTYPE HTML>
<html>
<head>
<meta charset="utf-8">
<title>{{.Title}}</title>
</head>
<body>
<h1>{{.Title}}</h1>
<p>{{.Content}}</p>
</body>
</html>
```

## Simple go

```golang
package main

import (
  "html/template"
  "net/http"
)

type Page struct {
     Title, Content string
}

func displayPage(w http.ResponseWriter, r *http.Request) {
p := &Page{
            Title: "An Example",
            Content: "Have fun stormin’ da castle.",
}

t := template.Must(template.ParseFiles("templates/simple.html"))
t.Execute(w, p)
}

func main() {
  http.HandleFunc("/", displayPage)
  http.ListenAndServe(":8080", nil)
}
```


Being context-aware is more than knowing that these are HTML templates. The
package understands what’s happening inside the templates. Take the following tem-
plate snippet:

```html
<a href="/user?id={{.Id}}">{{.Content}}</a>
```

The html/template package expands this intelligently. For escaping purposes, it adds
context-appropriate functionality. The preceding snippet is automatically expanded

to look like this:
```html
<a href="/user?id={{.Id | urlquery}}">{{.Content | html}}</a>
```

