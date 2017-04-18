## Add in template function


```golang
package main
import (
  "html/template"
  "net/http"
  "time"
)

var tpl = `<!DOCTYPE HTML>
<html>
<head>
<meta charset="utf-8">
<title>Date Example</title>
</head>
<body>
<p>{{.Date | dateFormat "Jan 2, 2006"}}</p>
</body>
</html>`

var funcMap = template.FuncMap{
    "dateFormat": dateFormat,

}

func dateFormat(layout string, d time.Time) string {
     return d.Format(layout)
}

func serveTemplate(res http.ResponseWriter, req *http.Request) {
t := template.New("date")
t.Funcs(funcMap)
t.Parse(tpl)
data := struct{ Date time.Time }{
                Date: time.Now(),
              }

t.Execute(res, data)
}

func main() {
http.HandleFunc("/", serveTemplate)
http.ListenAndServe(":8080", nil)
}
```


## Cashing

```golang
package main

import (
"html/template"
"net/http"
)

var t = template.Must(template.ParseFiles("templates/simple.html"))
type Page struct {
Title, Content string
}

func diaplayPage(w http.ResponseWriter, r *http.Request) {
p := &Page{
Title: "An Example",
Content: "Have fun stormin’ da castle.",
}

t.Execute(w, p)
}

func main() {
http.HandleFunc("/", diaplayPage)
http.ListenAndServe(":8080", nil)
}
```

## Teemplate

```golang
package main

import (
"bytes"
"fmt"
"html/template"
"io"
"net/http"
)

var t *template.Template
func init() {
t = template.Must(template.ParseFiles("./templates/simple.html"))
}

type Page struct {
Title, Content string
}

func diaplayPage(w http.ResponseWriter, r *http.Request) {
p := &Page{
Title: "An Example",
Content: "Have fun stormin’ da castle.",
}
var b bytes.Buffer
err := t.Execute(&b, p)
if err != nil {
  fmt.Fprint(w, "A error occured.")
  return
}

b.WriteTo(w)
}

func main() {
http.HandleFunc("/", diaplayPage)
http.ListenAndServe(":8080", nil)
}
```
