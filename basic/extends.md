## Main 

```golang
package main

import (
	"html/template"
	"net/http"
)

var hogeTmpl = template.Must(template.New("hoge").ParseFiles("base.html", "hoge.html"))

func hogeHandler(w http.ResponseWriter, r *http.Request) {
	hogeTmpl.ExecuteTemplate(w, "base", "Hoge")
}

var piyoTmpl = template.Must(template.New("piyo").ParseFiles("base.html", "piyo.html"))

func piyoHandler(w http.ResponseWriter, r *http.Request) {
	piyoTmpl.ExecuteTemplate(w, "base", "Piyo")
}

func main() {
	// hoge
	http.HandleFunc("/", hogeHandler)
	http.HandleFunc("/hoge", hogeHandler)

	// piyo
	http.HandleFunc("/piyo", piyoHandler)

	http.ListenAndServe(":8080", nil)
}
```

## base.html
```html
{{define "base"}}
<!DOCTYPE html>
<html>
    <body>

      <header>
        <h1>Title of {{.}}</h1>
        <nav>
          <ul>
            <li><a href="hoge">Hoge</a></li>
            <li><a href="piyo">Piyo</a></li>
          </ul>
        </nav>
      </header>

      <article id="content">
        {{template "content"}}
      </article>

      <footer>
        &copy; Copyright 2013 by golang-samples.
      </footer>
    </body>
</html>
{{end}}
```

## hoge.html
```html
{{define "content"}}
<h1>Hoge</h1>
<p>
  I'm Hoge.
</p>
{{end}}
```

## pioy.html

```html
{{define "content"}}
<h1>Piyo</h1>
<p>
  I'm Piyo.
</p>
{{end}}
```


