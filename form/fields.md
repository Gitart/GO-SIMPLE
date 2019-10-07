## Работа с полями формы
---
Можно получить перечень полей от формы в виде массива полей  




```golang
// /company/add/
func Company_add(rw http.ResponseWriter, req *http.Request){
	var Dat Company
    req.ParseForm()
    fmt.Printf("%+v\n", req.Form)
	rf      := req.FormValue
 

      fmt.Println("Value :",rf)

     rfs:=req.Form

     for v,r:=range(rfs){
         fmt.Println(v,r,r[0],"-------------\n")     	
     }

     r.DB("System").Table("Log").Insert(rfs).Run(sessionArray[0])
    fmt.Println(rfs)
    return

    pn      := rf("name")
    cn      := rf("Code")
    yp      := Times[0].Id


    if pn == "" {
       return 
    }


    Dat.Title  = pn
    Dat.name  = pn
    Dat.Code   = cn
    Dat.Date   = CTM()
    Dat.Remark = rf("Remark") 


     // Insert table and return id
     Comapnyid:=Insert_Company("Companyes", Dat)
     Cookies_write(rw, "ProjectId",  Comapnyid,  yp)
     
     fmt.Println("Добавлена новая компания ID: "+Comapnyid)

     // Redirect to journal projects
     redirect(rw,req,"/company/card/") 
}
```


### Output req.Form

```json

{ "Address": ["10"] ,
  "City": ["10"] ,
  "Director": ["ssssaa"] ,
  "Finish": ["16.09.2019"] ,
  "Project": ["/"] ,
  "Remark": ["www"] ,
  "Start": ["08.10.2019"] ,
  "Taskids": ["/"] ,
  "User": [""] ,
  "id": "cf655aa2-6c42-4e34-9613-ca87b306560f" ,
  "name": ["ssss"]
}

```
