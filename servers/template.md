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

