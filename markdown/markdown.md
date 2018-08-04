## Markdown

```go
package main

import (
 "github.com/russross/blackfriday"
 "html/template"
 "io/ioutil"
 "log"
 "net/http"
 "path"
 "strings"
)

type Post struct {
 Title string
 Body  template.HTML
}
var (
 // компилируем шаблоны, если не удалось, то выходим
 post_template = template.Must(template.ParseFiles(path.Join("templates", "layout.html"), path.Join("templates", "post.html")))
)

func main() {
 // для отдачи сервером статичный файлов из папки public/static
 fs := http.FileServer(http.Dir("./public/static"))
 http.Handle("/static/", http.StripPrefix("/static/", fs))
 http.HandleFunc("/", postHandler)
 log.Println("Listening...")
 http.ListenAndServe(":3000", nil)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
 // обработчик запросов
 fileread, _ := ioutil.ReadFile("posts/index.md")
 lines       := strings.Split(string(fileread), "\n")
 title       := string(lines[0])
 body        := strings.Join(lines[1:len(lines)], "\n")
 body         = string(blackfriday.MarkdownCommon([]byte(body)))
 post        := Post{title, template.HTML(body)}
 
 if err := post_template.ExecuteTemplate(w, "layout", post); err != nil {
   log.Println(err.Error())
   http.Error(w, http.StatusText(500), 500)
 }
}
```
