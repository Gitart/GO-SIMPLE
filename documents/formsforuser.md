### Форма для ввод нового пользователя
```go
/********************************************************************************************************************************
 *
 *  TITLE       : Форма для ввод нового пользователя
 *  DATE        : Savchenko Arthur
 *	PATH        : /tst/in/
 *	DESCRIPTION : Форма для ввода нового пользователя :
 *  PATH        : http://10.0.3.24:5555/static/newuser.html
 *
 *********************************************************************************************************************************/
func Add_Form(w http.ResponseWriter, req *http.Request) {
	/*
		W:=`<!DOCTYPE html>
			<html lang="en">
			<head>
			<meta charset="UTF-8">
			<title>Success Alert Message</title>
			<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.2/css/bootstrap.min.css">
			<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.2/css/bootstrap-theme.min.css">
			<script src="https://ajax.googleapis.com/ajax/libs/jquery/1.11.2/jquery.min.js"></script>
			<script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.2/js/bootstrap.min.js"></script>
			<style type="text/css">
				 .bs-example{margin: 20px;}
			</style>
			</head>
			<body>
			<div class="bs-example">
			    <div class="alert alert-success">
			        <a href="#" class="close" data-dismiss="alert">&times;</a>
			        <strong>CОХРАНЕНИЕ !</strong> Ваш запрос принят !!! Мы благодарны.
			    </div>
			</div>
			</body>
			</html>`
	*/
	CurTime = time.Now().Format(time.RFC3339)

	// Control
	t := true

	if t {
	   s := req.URL.String()
	   u, _ := url.Parse(s)
	   m, _ := url.ParseQuery(u.RawQuery)
	   fmt.Println(s, m)
	   fmt.Println(u.User)
	}

	now := time.Now()
	// secs  := now.Unix()
	// nanos := now.UnixNano()

	R := "Gest"

	// Если роль не определена считается как гость системы
	// с минимальными правами только на чтение
	if len(req.FormValue("Role")) == 0 {
 	   R = req.FormValue("Role")
	}

	var FN USER
	FN.Os           = req.UserAgent()
	FN.Ip           = basicAuth(req.FormValue("Fname"), req.FormValue("Password"))
	FN.ID           = now.Unix()
	FN.Insert_at    = now.String()
	FN.Update_at    = now.String()
	FN.Visit_at     = now.String()
	FN.Deadline     = now.AddDate(1, 0, 0).String() // + 1 Year добавление года для срока действия пользователя в системе
	FN.Structure    = STI(req.FormValue("Structure"))
	FN.Coorporation = STI(req.FormValue("Coorporation"))
	FN.Name         = req.FormValue("Fname") + " " + req.FormValue("Fname")
	FN.Fname        = req.FormValue("Fname")
	FN.Lname        = req.FormValue("Lname")
	FN.Mname        = req.FormValue("Mname")
	FN.Password     = req.FormValue("Password")
	FN.Telephone    = req.FormValue("Phonenum")
	FN.Telephonemob = req.FormValue("Phonemob")
	FN.Email        = req.FormValue("email")
	FN.Position     = req.FormValue("Position")
	FN.Role         = R
	FN.Description  = req.FormValue("Description")
	FN.Flag         = 1
	FN.Status       = "New"
	FN.Right.Access = true
	FN.Right.Delete = false
	FN.Right.Update = true
	FN.Right.Insert = true
	FN.Right.View   = true

	// TY    := req.PostForm
	// Запись в Response Headers системной информации
	// w.Header().Set("Content-Type", "application/x-www-form-urlencoded")

	w.Header().Set("Head-Office-version", "ver. 12.03")
	w.Header().Set("Head-Office", "Head Office")
	w.Header().Set("Head-result", "OK")
	w.Header().Set("Server", "A Go Web Server")
	w.Header().Set("Status-API", "200")

	// wwa:=w.Header().Get("Author")
	// var U Users
	// U:= req.URL.String()
	// Z:= req.PostForm
	// w.Write([]byte(W))
	// r.DB("HO").Table("Users").Delete().Run(sessionArray[0])
	// r.DB("HO").Table("Task").Insert(Mst{"Description": email, "Tasks":T, "URL":req.URL, "form":req.Form, "Header":req.Header,"Traller":req.TransferEncoding, "Post":req.PostForm, "Sss":req.PostFormValue("nn1"), "sdd":req.Body }).RunWrite(sessionArray[0])

	// Если не пустые основные поля производится запись
	if len(FN.Fname) > 0 && len(FN.Lname) > 0 {
		r.DB("HO").Table("Users").Update(map[string]interface{}{"Status": ""}).RunWrite(sessionArray[0])
		// r.DB("HO").Table("Users").Insert(Mst{"FULLNAME": FN, "STATUS": "New Users", "CREATED": CurTime,  "FIELDS": req.PostForm}).Merge(req.PostForm).RunWrite(sessionArray[0])
		// r.DB("HO").Table("Users").Insert(Mst{"FULLNAME": FN, "STATUS": "New Users", "CREATED": CurTime, "Fields":TY, "Nano": nanos, "Secs":secs, "now":now}).RunWrite(sessionArray[0])
		r.DB("HO").Table("Users").Insert(FN).RunWrite(sessionArray[0])
	}

	// fmt.Println(U,Z, FN )
	// Code 205 позволяет не прятать форму после добавления записи в базу
	// w.WriteHeader(205)

	// if len(P)==0 {
	//  fp   := path.Join("templates", "registration.htm")
	// tmpl, _ := template.ParseFiles(fp)
	// tmpl.Execute(w, FN)
	// }

	// fmt.Println("ok", http.StatusInternalServerError )

	// http.Redirect(w, req, "confirm.html", http.StatusSeeOther)
	// fmt.Println("redirect")
	// template.ParseFiles("templates/registration.htm")
	//render(w, "templates/registration.htm", nil)

	// Обновление страницы
	// render(w,"templates/confirm.htm",nil)

	// w.WriteHeader(301)
	// fmt.Println()
	// http://10.0.3.24:5555/api/system/users/viewall/

	// Возврат на страницу списка пользователей
	// http.Redirect(w, req, "http://"+AdresPort+"/api/system/users/viewall/", 301)

	// 301 Moved Permanently
	http.Redirect(w, req, "/api/system/users/viewall/", 301)
}
```

### Форма для ввод нового пользователя
```go
/********************************************************************************************************************************
 *
 * TITLE         : Запись в форму
 * DATE          :
 * DESCRIPTION   :
 * AUTHOR        : Savchenko Arthur
 * USAGE         :
 * ORGANIZATION  :
 *
 *********************************************************************************************************************************/
func foo(w http.ResponseWriter, r *http.Request) {

	profile := User{"Alex", "programming"}
	fp := path.Join("templates", "userform.html")
	tmpl, err := template.ParseFiles(fp)

	// Error
	if err != nil {
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}

	if err := tmpl.Execute(w, profile); err != nil {
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
```
