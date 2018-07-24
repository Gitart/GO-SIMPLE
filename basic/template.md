

## Parse files

```golang
package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	t := template.Must(template.ParseFiles("templates/main.tmpl", "templates/header.tmpl", "templates/footer.tmpl"))
	err := t.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}
}
```


## main.html

```html
{{template "header"}}
<p>main content</p>
{{template "footer"}}
```




## header.html

```html
{{define "header"}}
<!doctype html>
<html lang="ja">
<head>
	<meta charset="UTF-8">
	<title>Document</title>
</head>
<body>
{{end}}
```


## footer.html
```html
{{define "footer"}}
</body>
</html>
{{end}}
```
