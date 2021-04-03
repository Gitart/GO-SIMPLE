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
        ArtTicles(D, "article", "news", w)                 



        fmt.Println(ArtTiclesStr(D, "article"))
}



func ArtTicles(txt, tmplfile, outfile  string, w http.ResponseWriter) {
    var Data []Mst
    f, err := os.Create("./tmpl/" + outfile + ".html")
    defer f.Close()
    json.Unmarshal( []byte(txt), &Data)

    Dt := Mst{"Dts": Data, "Title": "Новости", "Descript": "Описание новостей за последнее время в мире", "Datrep": "sss"}
    fp := path.Join("tmpl", tmplfile + ".html")
    tmpl, err := template.ParseFiles(fp, "tmp/main.html")
    ut.Err(err, "Error template execute.")

    
    var tpl bytes.Buffer



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




func ArtTiclesStr(txt, tmplfile string) string {
    var tpl bytes.Buffer
    var Data []Mst
    json.Unmarshal( []byte(txt), &Data)

    Dt := Mst{"Dts": Data, "Title": "Новости", "Descript": "Описание новостей за последнее время в мире", "Datrep": "sss"}
    fp := path.Join("tmpl", tmplfile + ".html")
    tmpl, err := template.ParseFiles(fp, "tmp/main.html")
    ut.Err(err, "Error template execute.")

    // Выввод в текстовую строку
    tmpl.Execute(&tpl, Dt)
    s := tpl.String()

    return s   

}
```
