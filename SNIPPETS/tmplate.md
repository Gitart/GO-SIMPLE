
## Построение стрингового значения на основе указанного темплейта

```go 
/******************************************************************** 
 *
 * Построение стрингового значения на основе указанного темплейта
 * templates = template.Must(template.ParseFiles("tmpl/edit.html", "tmpl/view.html"))
 * func TemplString(Templatename string, Data Mst) string {
 * 	
 ********************************************************************/
func TemplString(Dat Mst) string {

	var tpl bytes.Buffer

	tmpl, err := template.ParseFiles("tmp/article.html")
	if err != nil {
		return err.Error()
	}

	errs := tmpl.Execute(&tpl, Dat)
	if errs != nil {
		return errs.Error()
	}
	return tpl.String()
}
```

## Template

```go
/********************************************************************
 *
 * Render html pages
 * RenderHtml("templname.html", Data, w)  
 *
 ********************************************************************/
func RenderHtml(Template string, Data Mst, w http.ResponseWriter) {
	// fp := path.Join("tmp/", Template)
	tmpl, err := template.ParseFiles("tmp/"+Template, "tmp/main.html")

	// Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	tmpl.Execute(w, Data)
}
```
