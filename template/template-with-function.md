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



// 
// Color from template
// 
func Clrs(Txt string) string{
    if Txt=="Waiting"     {return "grey"}
    if Txt=="Done"        {return "green"}
    if Txt=="Progress"    {return "red"}
	return "yellow"
}
```

### Template 
Здесь используется обработка функций

```html

  <table id="tabledata" class='table table-sm table-striped table-hover'>
      <thead>
         <tr>
            <th>Id</th>
            <th>Tasks</th>
            <th>Module</th>
            <th>%</th>
            <th>Status</th>
         </tr>
      </thead>

      <tbody>
         {{range .Dat}}
              <tr >
                 <td class="veral">{{.Id}}</td>
                 <td><b style="color:#3E9C9C;">{{.Title}}</b><br>  <small>{{.Descript}}</small></td>
                 <td class="veral">{{.Module}}</td>
                 <td class="veral">{{.Percent}}</td>
                 <td class="veral" style="color:{{.Status|Fad}}">{{.Status}}</td>
              </tr>
         {{end}}
      </tbody>

  </table>
```

