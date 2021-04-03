## Вывод в разные форматы

## Опредление статических страниц
```go
/*
 *   Статические странички
 *   c установкой разрешений и доступов на операции
 *
 *   http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {http.ServeFile(w, r, r.URL.Path[1:])})
 *   http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {http.FileServer(http.Dir("/static/"))})
 *
 *   /static/....
 *
 */
func StaticPage(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", "*") // Allows
  // Allows
  // if origin := r.Header().Get("Origin"); origin != "" {
  //  w.Header().Set("Access-Control-Allow-Origin", origin)
  //  w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
  //  w.Header().Set("Access-Control-Allow-Headers",  "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
  //  w.Header().Set("Access-Control-Max-Age", "86400") // 24 hours
  // }

  //  File static page
  http.ServeFile(w, r, r.URL.Path[1:])
}
```


## Public
```go
/*
 *
 *   Статические странички
 *   c установкой разрешений и доступов на операции
 *
 *   http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {http.ServeFile(w, r, r.URL.Path[1:])})
 *   http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {http.FileServer(http.Dir("/static/"))})
 *   /static/....
 *
 */
func PublicPage(w http.ResponseWriter, r *http.Request) {
	// Allows
	w.Header().Set("Access-Control-Allow-Origin", "*") 

	//  File static page
	http.ServeFile(w, r, r.URL.Path[1:])
}
```


## Вывод в разные форматы
```go
package main

import (
	"html/template"
    "net/http"
    "encoding/json"
    "path"
    "os"
    "fmt"
    "bytes"
 ut "./lib/util"
)

/*
 *   Новости
 */
func Articles_new(w http.ResponseWriter, r *http.Request) {
        D :=` [ 
                {"id":1, "title":"Новость 1", "Descript":"Описание сайта 1"},
                {"id":2, "title":"Новость 3", "Descript":"Описание сайта 4"},
                {"id":3, "title":"Новость 4", "Descript":"Описание сайта 5"},
                {"id":7, "title":"Новость 7", "Descript":"Описание сайта 6"}
             ]
           `


      content := `
                  <h3> Работа с технологичным модулем ардурино</h3>
                  <p>В этом скетче объявлен пин 13 для управления светодиодом, но на плате NodeMCU  нет такого пина, на плате все пины начинаются с буквы. Ниже привел картинку, сопоставления пинов NodeMCU с пинами программирования в Arduino IDE</p>
                  <p>Давайте подключим датчик движения. Перед подключением его необходимо настроить. Ставим перемычку на букву L, и тогда он будет слать одиночные импульсы. Первый потенциометр ставьте на минимум, чтобы импульсы были короткие (нам важно засечь движение). А второй потенциометр — это дальность срабатыватывания. Советую сразу ставить на максимум.</p>
                  <p>Теперь подключите все по схеме на картинке. Светодиод необязательный, но подключив его, удобно видеть срабатывание датчика.</p>                  

      `
      

        ArtTicles(D, "article", "news", content,  w)                 


        // fmt.Println(ArtTiclesStr(D, "article"))
}




func ArtTicles(txt, tmplfile, outfile, content  string, w http.ResponseWriter) {
    var tpl bytes.Buffer
    var Data []Mst
    f, err := os.Create("./tmpl/" + outfile + ".html")
    defer f.Close()

    json.Unmarshal( []byte(txt), &Data)


    ContentHtml:=template.HTML(content)

    Dt := Mst{ "Dts":      Data, 
               "Content":  ContentHtml, 
               "Title":    "Новости", 
               "Descript": "Описание новостей за последнее время в мире", 
               "Datrep"  : "Дата 02.03.2021"}

    fp := path.Join("tmpl", tmplfile + ".html")
    tmpl, err := template.ParseFiles(fp, "tmp/main.html")
    ut.Err(err, "Error template execute.")

    // Формипрование локального файла на диске
    tmpl.Execute(f, Dt)

    //выод файла на страницу сайта
    tmpl.Execute(w, Dt)

    // Выввод в текстовую строку
    tmpl.Execute(&tpl, Dt)
    s := tpl.String()
    fmt.Println(s)
   // ut.Writfile(File, r)    
}




// *********************************************************
// Получение в строковом виде на основании шаблона
// *********************************************************
func ArtTiclesStr(txt, tmplfile string) string {
    var tpl bytes.Buffer
    var Data []Mst
    json.Unmarshal([]byte(txt), &Data)

    Dt := Mst{"Dts": Data, "Title": "Новости", "Descript": "Описание новостей за последнее время в мире", "Datrep": "sss"}
    
    // 
    fp := path.Join("tmpl", tmplfile + ".html")
    fs := path.Join("tmp",  "main.html") 

    // Основной шаблон с темплейтами 
    // tmpl, err := template.ParseFiles(fp, "tmp/main.html")         
    tmpl, err := template.ParseFiles(fp, fs)         
    ut.Err(err, "Error template execute.")

    // Выввод в текстовую строку
    tmpl.Execute(&tpl, Dt)
    s := tpl.String()
    return s   
}
```
