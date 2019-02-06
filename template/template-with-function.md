# Использование темплейтов с функциями



### Функция для темплейта

```golang
// *****************************************************
// Report plan
// /api/report/plan/
// *****************************************************
func Rep_plan(w http.ResponseWriter, r *http.Request){
     
    db:=Dbcc()

    var records []Plan
    db.Find(&records)
    Dt := Mst{"Dat": records , "Title":"Task report", "Descript": "Current stage ", "Datrep": " " + Ctm()}

      
 // Maping function
    funcMap := template.FuncMap{"Fad": Clrs, "Fsd": Clrs_c}
	fp      := path.Join("tmp", "plan.html")                               
	tmpl,err:= template.New("plan.html").Funcs(funcMap).ParseFiles(fp, "tmp/main.html")
	Err(err, "Error template execute.")

    // tmpl.Funcs(fncmap)
	errf    := tmpl.Execute(w, Dt)
	Err(errf, "Error templates execute.")

    // fp := path.Join("tmp", "plan.html")
    // tmpl, err := template.ParseFiles(fp, "tmp/main.html")
    // tmpl.Execute(w, Dts)
    // RenderHtml("plan.html",Dts,w)
}
```
