### Веб-приложение на Go: редактирование страниц, пакет html/template

Вики это не настоящий вики, если нет возможности редактировать страницы. Давайте создадим два новых обработчика: один с именем `editHandler` для отображения формы редактирования страницы, и другой с именем `saveHandler` для сохранения данных, введенных через форму.

Сначала мы добавим их в `main()`:

```go
func main() {
    http.HandleFunc("/view/", viewHandler)
    http.HandleFunc("/edit/", editHandler)
    http.HandleFunc("/save/", saveHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

```

Функция `editHandler` загружает страницу (или, если он не существует, создает пустую структуру `Page`), и отображает HTML форму.

```go
func editHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/edit/"):]
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    fmt.Fprintf(w, "<h1>Editing %s</h1>"+
        "<form action=\"/save/%s\" method=\"POST\">"+
        "<textarea name=\"body\">%s</textarea><br>"+
        "<input type=\"submit\" value=\"Save\">"+
        "</form>",
        p.Title, p.Title, p.Body)
}

```

Эта функция будет хорошо работать, но весь этот жестко закодированный HTML ужасен. Конечно, есть лучший способ.

### Пакет `html/template`

Пакет `html/template` является частью стандартной библиотеки Go. Мы можем использовать `html/template`, чтобы хранить HTML в отдельном файле, что позволяет нам изменить макет нашей страницы редактирования без изменения основного Go кода.

Во-первых, мы должны добавить `html/template` в список импорта. Мы также не будем больше использовать `fmt`, поэтому мы должны удалить его.

```go
import (
    "html/template"
    "io/ioutil"
    "net/http"
)

```

Давайте создадим файл шаблона, содержащий HTML форму. Откройте новый файл с именем `edit.html`и добавьте следующие строки:

```go
<h1>Editing {{.Title}}</h1>

<form action="/save/{{.Title}}" method="POST">
<div><textarea name="body" rows="20" cols="80">
{{printf "%s" .Body}}
</textarea></div>
<div><input type="submit" value="Save"></div>
</form>

```

Измените `editHandler`, чтобы использовать шаблон вместо жестко заданного HTML кода:

```go
func editHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/edit/"):]
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    t, _ := template.ParseFiles("edit.html")
    t.Execute(w, p)
}

```

Функция `template.ParseFiles` будет читать содержимое `edit.html` и возвращать `*template.Template`.

Метод `t.Execute` выполняет шаблон, записывая сгенерированный HTML для `http.ResponseWriter`. Точечные идентификаторы `.Title` и `.Body` относятся к `p.Title` и `p.Body`.

Шаблонные директивы заключены в двойные фигурные скобки. Инструкция `printf "%s" .Body` является вызовом функции, которая выводит `.Body` в виде строки вместо потока байтов, такой же, как вызов `fmt.Printf`. Пакет `html/template` помогает гарантировать, что только безопасный и правильно выглядящий HTML генерируется действиями шаблона. Например, он автоматически экранирует знак «больше» (`>`), заменяя его с помощью `&gt;`, чтобы убедиться, что данные пользователя не повреждают HTML форму.

Поскольку сейчас мы работаем с шаблонами, давайте создадим шаблон для нашего `viewHandler` называемый `view.html`:

```go
<h1>{{.Title}}</h1>

<p>[<a href="/edit/{{.Title}}">edit</a>]</p>
<div>{{printf "%s" .Body}}</div>

```

Измените `viewHandler` соответственно:

```go
func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, _ := loadPage(title)
    t, _ := template.ParseFiles("view.html")
    t.Execute(w, p)
}

```

Обратите внимание, что мы использовали почти одинаковый шаблонный код в обоих обработчиках. Давайте удалим это дублирование, переместив шаблонный код в свою собственную функцию:

```go
func renderTemplate(w http.ResponseWriter,
                    tmpl string, p *Page) {
    t, _ := template.ParseFiles(tmpl + ".html")
    t.Execute(w, p)
}

```

И измените обработчики для использования этой функции:

```go
func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, _ := loadPage(title)
    renderTemplate(w, "view", p)
}
func editHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/edit/"):]
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    renderTemplate(w, "edit", p)
}

```

Если мы закомментируем регистрацию нашего невыполненного обработчика сохранения в `main`, мы можем еще раз собрать и протестировать нашу программу. Взглянем на весь код, который мы уже написали.

```go
package main

import (
  "html/template"
  "io/ioutil"
  "log"
  "net/http"
)

type Page struct {
  Title string
  Body  []byte
}

func (p *Page) save() error {
  filename := p.Title + ".txt"
  return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
  filename := title + ".txt"
  body, err := ioutil.ReadFile(filename)
  if err != nil {
    return nil, err
  }
  return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter,
                    tmpl string, p *Page) {
  t, _ := template.ParseFiles(tmpl + ".html")
  t.Execute(w, p)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
  title := r.URL.Path[len("/view/"):]
  p, _ := loadPage(title)
  renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
  title := r.URL.Path[len("/edit/"):]
  p, err := loadPage(title)
  if err != nil {
    p = &Page{Title: title}
  }
  renderTemplate(w, "edit", p)
}

func main() {
  http.HandleFunc("/view/", viewHandler)
  http.HandleFunc("/edit/", editHandler)
  //http.HandleFunc("/save/", saveHandler)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
```


[Template description](https://golang-blog.blogspot.com/2019/02/go-web-app-html-template.html)
![image](https://user-images.githubusercontent.com/3950155/152210248-2b02488a-2a88-4b47-9b96-b604b6eaf904.png)

