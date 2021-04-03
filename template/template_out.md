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
    "time"
 ut "./lib/util"
)

/*
 *   Новости
 */
func Articles_new(w http.ResponseWriter, r *http.Request) {
        D :=` [ 
                {"id":1, "title":"Новость 1", "Descript":"Описание сайта 1", "Author":"Stepan Abraham Linkoln", "Date":"02 june 2020", "Content":"Ставим"},
                {"id":2, "title":"Новость 3", "Descript":"Описание сайта 4", "Author":"Stepan Abraham Linkoln", "Date":"02 june 2020", "Content":"Перед"},
                {"id":3, "title":"Новость 4", "Descript":"Описание сайта 5", "Author":"Stepan Abraham Linkoln", "Date":"02 june 2020", "Content":"Получение"},
                {"id":1, "title":"Новость 1", "Descript":"Описание сайта 1", "Author":"Stepan Abraham Linkoln", "Date":"02 june 2020", "Content":"Новость"},
                {"id":2, "title":"Новость 3", "Descript":"Описание сайта 4", "Author":"Stepan Abraham Linkoln", "Date":"02 june 2020", "Content":"Светодиод"},
                {"id":3, "title":"Новость 4", "Descript":"Описание сайта 5", "Author":"Stepan Abraham Linkoln", "Date":"02 june 2020", "Content":"модулем"},
                {"id":7, "title":"Новость 7", "Descript":"Описание сайта 6", "Author":"Stepan Abraham Linkoln", "Date":"02 june 2020", "Content":"локального"}

             ]
           `

      contentard := `
                  <h3> Работа с технологичным модулем ардурино</h3>
                  <p>В этом скетче объявлен пин 13 для управления светодиодом, но на плате NodeMCU  нет такого пина, на плате все пины начинаются с буквы. Ниже привел картинку, сопоставления пинов NodeMCU с пинами программирования в Arduino IDE</p>
                  <p>Давайте подключим датчик движения. Перед подключением его необходимо настроить. Ставим перемычку на букву L, и тогда он будет слать одиночные импульсы. Первый потенциометр ставьте на минимум, чтобы импульсы были короткие (нам важно засечь движение). А второй потенциометр — это дальность срабатыватывания. Советую сразу ставить на максимум.</p>
                  <p>Теперь подключите все по схеме на картинке. Светодиод необязательный, но подключив его, удобно видеть срабатывание датчика.</p>                  
                  <p>Бывает, что нужно обеспечить автономное питание проекта, т.е. вдали от розетки, давайте рассмотрим варианты. Также для этих целей пригодится <a href="https://alexgyver.ru/lessons/power-sleep/" target="_blank" rel="noopener noreferrer">урок по энергосбережению</a>    и режимам сна микроконтроллера.</p>
                  `

    // Формирование стрингового контента на основании шаблона
    contents := ArtTiclesStr(D, "vid")
   
    ArtTicles(D, "article", "news", contents, contentard, w)                 
}


// ****************************************************************
// Конвертация стрингового вырвжения в HTML
// ****************************************************************
func HtmlConvert(content string) template.HTML {
      cnvert := template.HTML(content)
     return cnvert
}


// **********************************************************************
// Формирование страницы 
// **********************************************************************
func ArtTicles(txt, tmplfile, outfile, content, contentard  string, w http.ResponseWriter) {
    

    fmt.Println(content)

    var tpl bytes.Buffer
    var Data []Mst

    f, err := os.Create("./tmpl/" + outfile + ".html")
    defer f.Close()

    json.Unmarshal([]byte(txt), &Data)
    

    Dt := Mst{ "Dts":      Data,                                                    // Дата JSON  
               "Content":     HtmlConvert(content),                                 // Наполнение текст
               "Contentard":  HtmlConvert(contentard),                              // Наполнение текст
               "Vidjet":   "Виджет для теста ",                                     // Наполнение текст
               "Warning":  "Внимание внимание",
               "Title":    "Новости", 
               "Descript": "Описание новостей за последнее время в мире", 
               "Datrep"  : "Дата 02.03.2021"}

    // Основной блок 
    fp := path.Join("tmpl",  tmplfile + ".html")
    
    // Темплайты
    fs := path.Join("tmp",   "main.html") 
    ft := path.Join("tmpl",  "tmp.html") 

    
    tmpl, err := template.ParseFiles(fp, fs, ft)
    ut.Err(err, "Error template execute.")

    // Формипрование локального файла на диске
    tmpl.Execute(f, Dt)

    //выод файла на страницу сайта
    tmpl.Execute(w, Dt)

    // Выввод в текстовую строку
    tmpl.Execute(&tpl, Dt)
   
   // Control
   //  s := tpl.String()
    fmt.Println("Ready...")

   // // ut.Writfile(File, r)    
}


// *********************************************************
// Получение в строковом виде на основании шаблона
// *********************************************************
func ArtTiclesStr(txt, tmplfile string) string {
    var tpl bytes.Buffer
    var Data []Mst
    json.Unmarshal([]byte(txt), &Data)

    Dt := Mst{"Dts"      : Data, 
              "Title"    : "Новости", 
              "Descript" : "Описание новостей за последнее время в мире", 
              "Datrep"   : time.Now().Format("02-01-2006 15:40")}
    

     // fp := path.Join("tmpl", "tmp.html")
    // 
    // fp := path.Join("tmpl", tmplfile + ".html")
    // fs := path.Join("tmp",  "main.html") 

    // Основной шаблон с темплейтами 
    // tmpl, err := template.ParseFiles(fp, "tmp/main.html")         
    tmpl, err := template.ParseFiles("tmpl/" + tmplfile + ".html")         
    ut.Err(err, "Error template execute.")

    // Выввод в текстовую строку
    tmpl.Execute(&tpl, Dt)
    s := tpl.String()
    return s   
}
```

## /tmpl/vid.html
```html
{{range .Dts}} 
       <div style="background-color:#CD5C5C; color:white; border: 1px solid #CCC; padding: 5px; border-radius: 5px; text-align: center; margin-bottom: 10px;">  
	    <b>{{.title}} </b>            {{.Descript}}      
	  </div>
{{end}}
```

## /tmpl/tmp.html
```html
{{define "basstyle"}}

<style>
html, body {
    margin: 0;
    padding: 0;
    font-family: "Work Sans", "Helvetica Neue", "Helvetica", Helvetica, Arial, sans-serif;
    font-weight:   400;
    background:     #F5F6F7;
    color:          #1F2667;
    text-rendering: optimizeLegibility;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    font-size: 16px;
    overflow-x: hidden;
}
</style>
{{end}}

{{define "tyr"}}
      <h1></h1>
	  <hr>
	  <h3>Системное сообщение</h3>
	 
      <div style="background-color:#2ECC71; color:white; border: 1px solid #CCC; padding: 5px; border-radius: 5px; text-align: center;">  
	    {{.Vidjet}}
	  </div>
{{end}}

{{define "vidjet_warning"}}
       <div style="background-color:#2ECC71; color:white; border: 1px solid #CCC; padding: 5px; border-radius: 5px; text-align: center;">  
	    {{.Warning}}
	  </div>
{{end}}

{{define "WidjetArticle"}}
   {{range .Dts}} 
       <div style="background-color:#2ECC71; border: 1px solid #CCC; padding: 5px; border-radius: 5px; text-align: center; margin-bottom: 10px;">  
	    {{.title}}         <br>
	    {{.Descript}}      <br>
	    {{.Id}}        <br> 
	    {{.Date}}          <br>
	    <hr>
        {{.Content}}
        {{.Url}}
	  </div>
   {{end}}
{{end}}
```

## /tmpl/article.html
```html
<html>
<head>
  {{template "libs"}}
  {{template "MainStyle"}}
 
  <link href="https://fonts.googleapis.com/css?family=Open+Sans:300,300i,400,400i,600,600i,700,700i|Montserrat:300,400,500,600,700" rel="stylesheet">
  {{template  "basstyle"}}

  <style>
        body {font-family: "Open Sans"}
  </style>
</head>

<body>
<div class="container">
    <h1>ART TECH</h1>
    {{template "headertop"}}
    {{template "tyr" .}}
 
    <h4>Навигационное меню</h4>
 
	<ul>
			{{range .Dts}} 
				  <li>{{.id}} <b>{{.title}}</b> {{.Descript}}</li>
			{{end}}
	</ul>

    <br>
    <h4>Область контента</h4>
    <div>
	       {{.Content}}
	</div>
    <br>
 

    <h4>Область Arduino</h4>
    <div>
	       {{.Contentard}}
	</div>
    <br>

 
    <h4> {{.Descript}} </h4>

    {{template "WidjetArticle" .}}
    {{template "vidjet_warning" .}}
    {{template "a1" .}}
    {{template "footer"}}
</div>	
</body>
</html>
```


