## HTTP Response Snippets for Go

Taking inspiration from the Rails layouts and rendering guide, 
I thought it'd be a nice idea to build a snippet collection 
illustrating some common HTTP responses for Go web applications.

- Sending Headers Only
- Rendering Plain Text
- Rendering JSON
- Rendering XML
- Serving a File
- Rendering a HTML Template
- Rendering a HTML Template to a String
- Using Layouts and Nested Templates

### Sending Headers Only


#### File: main.go

```golang
package main

import (
  "net/http"
)

func main() {
  http.HandleFunc("/", foo)
  http.ListenAndServe(":3000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Server", "A Go Web Server")
  w.WriteHeader(200)
}
```


```
$ curl -i localhost:3000
HTTP/1.1 200 OK
Server: A Go Web Server
Content-Type: text/plain; charset=utf-8
Content-Length: 0
```


### Rendering Plain Text
File: main.go

```golang
package main

import (
  "net/http"
)

func main() {
  http.HandleFunc("/", foo)
  http.ListenAndServe(":3000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("OK"))
}
```

```
$ curl -i localhost:3000
HTTP/1.1 200 OK
Content-Type: text/plain; charset=utf-8
Content-Length: 2

OK
```



### Rendering JSON
File: main.go

```golang
package main

import (
  "encoding/json"
  "net/http"
)

type Profile struct {
  Name    string
  Hobbies []string
}

func main() {
  http.HandleFunc("/", foo)
  http.ListenAndServe(":3000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
  profile := Profile{"Alex", []string{"snowboarding", "programming"}}

  js, err := json.Marshal(profile)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(js)
}
```

```
$ curl -i localhost:3000
HTTP/1.1 200 OK
Content-Type: application/json
Content-Length: 56

{"Name":"Alex",Hobbies":["snowboarding","programming"]}
```


### Rendering XML
File: main.go

```golang
package main

import (
  "encoding/xml"
  "net/http"
)

type Profile struct {
  Name    string
  Hobbies []string `xml:"Hobbies>Hobby"`
}

func main() {
  http.HandleFunc("/", foo)
  http.ListenAndServe(":3000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
  profile := Profile{"Alex", []string{"snowboarding", "programming"}}

  x, err := xml.MarshalIndent(profile, "", "  ")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/xml")
  w.Write(x)
}
```

```
$ curl -i localhost:3000
HTTP/1.1 200 OK
Content-Type: application/xml
Content-Length: 128

<Profile>
  <Name>Alex</Name>
  <Hobbies>
    <Hobby>snowboarding</Hobby>
    <Hobby>programming</Hobby>
  </Hobbies>
</Profile>
```



### Serving a File
File: main.go

```golang
package main

import (
  "net/http"
  "path"
)

func main() {
  http.HandleFunc("/", foo)
  http.ListenAndServe(":3000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
  // Assuming you want to serve a photo at 'images/foo.png'
  fp := path.Join("images", "foo.png")
  http.ServeFile(w, r, fp)
}
```


```
$ curl -I localhost:3000
HTTP/1.1 200 OK
Accept-Ranges: bytes
Content-Length: 236717
Content-Type: image/png
Last-Modified: Thu, 10 Oct 2013 22:23:26 GMT
```


### Rendering a HTML Template
File: templates/index.html

```html
<h1>Hello {{ .Name }}</h1>
<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit.</p>
```


### File: main.go

```golang
package main

import (
  "html/template"
  "net/http"
  "path"
)

type Profile struct {
  Name    string
  Hobbies []string
}

func main() {
  http.HandleFunc("/", foo)
  http.ListenAndServe(":3000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
  profile := Profile{"Alex", []string{"snowboarding", "programming"}}

  fp := path.Join("templates", "index.html")
  tmpl, err := template.ParseFiles(fp)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  if err := tmpl.Execute(w, profile); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}
```

```
$ curl -i localhost:3000
HTTP/1.1 200 OK
Content-Type: text/html; charset=utf-8
Content-Length: 84

<h1>Hello Alex</h1>
<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit.</p>
````


### Rendering a HTML Template to a String
Instead of passing in the http.ResponseWriter when executing your 
template (like in the above snippet) use a buffer instead:

File: main.go

```golang
...
buf := new(bytes.Buffer)
if err := tmpl.Execute(buf, profile); err != nil {
  http.Error(w, err.Error(), http.StatusInternalServerError)
}
templateString := buf.String()
...
```

#### Using Layouts and Nested Templates
File: templates/layout.html

```html
<html>
  <head>
    <title>{{ template "title" . }}</title>
  </head>
  <body>
    {{ template "content" . }}
  </body>
</html>
```


#### File: templates/index.html

```html
{{ define "title" }}An example layout{{ end }}

{{ define "content" }}
<h1>Hello {{ .Name }}</h1>
<p>Lorem ipsum dolor sit amet, consectetur adipisicing elit.</p>
{{ end }}
```



File: main.go

```golang
package main

import (
  "html/template"
  "net/http"
  "path"
)

type Profile struct {
  Name    string
  Hobbies []string
}

func main() {
  http.HandleFunc("/", foo)
  http.ListenAndServe(":3000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
  profile := Profile{"Alex", []string{"snowboarding", "programming"}}

  lp := path.Join("templates", "layout.html")
  fp := path.Join("templates", "index.html")

  // Note that the layout file must be the first parameter in ParseFiles
  tmpl, err := template.ParseFiles(lp, fp)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  if err := tmpl.Execute(w, profile); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}
```

```

$ curl -i localhost:3000
HTTP/1.1 200 OK
Content-Type: text/html; charset=utf-8
Content-Length: 180

<html>
  <head>
    <title>An example layout</title>
  </head>
  <body>
    <h1>Hello Alex</h1>
    <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit.</p>
  </body>
</html>
```


