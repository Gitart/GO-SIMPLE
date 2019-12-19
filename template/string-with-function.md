# Формирование строковой переменной на основе шаблона 

Формирование строки на основе шаблона который хранится в директории tmp/sea.html и с использованием функций


```golang
 
/ ******************************************************************
// Тестирование
// https://gowebexamples.com/hello-world/
// https://blog.gopheracademy.com/advent-2017/using-go-templates/
// https://github.com/Gitart/hr/blob/master/main.go#L2271
// ******************************************************************
func Creat_string(w http.ResponseWriter, r *http.Request) {


  var tpl bytes.Buffer
  D:=[]Mst{
  	          {"Descript":"d1",   "Note":"Пример 1",  "Done": false, "Summ":22.00  , "Navigation":"pending" },
  	          {"Descript":"d2",   "Note":"Пример 2",  "Done": true , "Summ":32.02  , "Navigation":"completed"},
  	          {"Descript":"d3",   "Note":"Пример 3",  "Done": false, "Summ":56.01  , "Navigation":"deleted"},
  	          {"Descript":"d4",   "Note":"Пример 4",  "Done": false, "Summ":112.89 , "Navigation":"edit"},
  	          {"Descript":"long", "Note":"Пример 5",  "Done": false, "Summ":112.89 , "Navigation":"deleted"},
  	        } 

 
 
    Dt:= Mst{"Title": "Поиск по сайту.", "Dat": "Test", "Dts":D, "Yes":"status"}

    // Maping function
    funcMap := template.FuncMap{"Fad": Tmp_a, "Fsd": Tmp_c, "Fcc": Tmp_cc}
	fp      := path.Join("tmp", "sea.html")                               
	tmpl,err:= template.New("sea.html").Funcs(funcMap).ParseFiles(fp, "tmp/main.html")
	Err(err, "Error template execute.")

    // tmpl.Funcs(fncmap)
	errf    := tmpl.Execute(&tpl, Dt)
	Err(errf, "Error templates execute.")


  // Page 
  // tmpl, err := template.ParseFiles("tmp/sea.html", "tmp/main.html")
  // if err != nil {fmt.Println("Template error", err.Error()) }

  // errs := tmpl.Execute(&tpl, Dt)
  // if errs != nil {    fmt.Println("Execute error", errs.Error())  }

  s := tpl.String()
   // s:=TemplString(Dt)
   // p := template.HTML(rr)
    fmt.Println(s)   

}

// ***********************************************************
// Color tasks for 
// ***********************************************************
func Tmp_a(t float64) string {
	if t>32{
	   return "Ok"	
	}
	return "bad"
}


// ***********************************************************
// Color for Active Release
// ***********************************************************
func Tmp_c(t string) (string,error) {
   r:=""
  if t=="d2"{
     r = "table-danger" //success
   }else{
	 r= ""
   }

   return r,nil
}

// ***********************************************************
// Color for Active Release
// ***********************************************************
func Tmp_cc(t string) (string,error) {
   r:=""
  if t=="long"{
     r = "To long line" //success
   }else{
	 r= "Normal line"
   }

   return r, nil
}

```

## Template
sea.html

```html
<h1>    {{.Dat}}     </h1>
<title> {{.Title}}   </title>
{{$ri := .Title}}

{{block "start" .}} Блок start ! {{end}}


	

# Работа в цикле 
{{- range .Dts}}
      

     <title> {{if eq .Navigation "pending"}}           Tasks
             {{ else if eq .Navigation "completed"}}   Completed
             {{ else if eq .Navigation "deleted"}}     Deleted
             {{ else if eq .Navigation "edit"}}        Edit
             {{end}}
     </title>



      *************************************************************************************************
      👽 Title : {{$ri}}
      В цикле что бы не менялось общее значение выставляем через переменную
      **********************************************************************
      Fad: {{.Summ|Fad}}
      Len -------->{{.Descript|Fcc}}

       {{.Descript|Fsd}}
       -----------------------

      {{if .Done}}
              <li > Yes  😃{{printf "DONE !** %-20s ***" .Descript}}</li>
      {{else}}
               <li> No 📗  {{.Note}}</li>
      {{end}}


      Примечание :  {{.Note}} 
      Описание   :  {{.Descript}}
    *************************************************************************************************
{{- end}}

{{if .Yes }} 
     Yes !
{{end}}

{{block "content" .}}
Блок финиш
{{end}}

```


