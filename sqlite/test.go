package main


import (
	// "C"
	"database/sql"
	"encoding/json"
	"fmt"
  _ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"strconv"
	"time"
	"html/template"
	"path"
)



// *****************************************************************************
// Стартовая процедура
// *****************************************************************************
func main() {

	// dir := http.Dir("./files")

	http.HandleFunc("/",               Static_Page)              // Статические страницы
	http.HandleFunc("/addrecord/",     DisplayPage)              // Добавление записи

    http.HandleFunc("/addlinks/",      AddLinks)                 // Добавление записи
    http.HandleFunc("/replinks/",      Db_lnk_report)            // Добавление записи

	http.HandleFunc("/adddoc/",        Addform)                  // Добавление документа
	http.HandleFunc("/cldoc/",         Clear_docs)               // Удаление всех записей
	http.HandleFunc("/delone/",        Delonerec)                // Удаление одной записи
	http.HandleFunc("/docscan/",       DbScan)                   // сканирование документов
	http.HandleFunc("/returnrec/",     Retrec)                   // Возврат записи
	http.HandleFunc("/updaterec/",     Updaterec)                // Обновление записи
	http.HandleFunc("/testrec/",       TestInsertRecord)         // Тест на вставку 100 000 записей
	http.HandleFunc("/testrec2/",      TestInsertRecord2)        // Тест на вставку 100 000 записей



	// http.ListenAndServe(":5555", http.FileServer(dir))
	fmt.Println("OK Start Server.")
	http.ListenAndServe(":5555", nil)
}


// *******************************************************************************************
//   Статические странички
//   c установкой разрешений и доступов на операции
//   http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {http.ServeFile(w, r, r.URL.Path[1:])})
//   http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {http.FileServer(http.Dir("/static/"))})
//   /static/....
// *******************************************************************************************
func Static_Page(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	http.ServeFile(w, r, r.URL.Path[1:])
	// http.FileServer(http.Dir("/barsetka/"))
}




type Mreturn struct {
     
     HDF_SEQ         int64                                    // Счетчик модификации данных
     HDF_TIME_STR    string                                   // Время модификации (2006-01-02T15:04:05.000)
     
}	


// *****************************************************************************
// https://astaxie.gitbooks.io/build-web-application-with-golang/en/07.4.html
// 
// *****************************************************************************
func DisplayPage2(w http.ResponseWriter, r *http.Request) {
    // var Perm []map[string]interface{}
    
     var m []Mst
     
     // m[0]["ggg"]="dddd"
     // m[1]["ddggg"]="dddd"
     // m=Mst{"ddd":"ddd"}
     // m=Mst{"ddd2":"ddd"}
     m=append(m,Mst{"ffff":"dfff"})
    //   Perm[0].HDF_TIME_STR="ssss"
    fmt.Println(m)
}




// *****************************************************************************
// 
// Дата 23.11.2017
// Отчет по линкам в табличном виде
// https://astaxie.gitbooks.io/build-web-application-with-golang/en/07.4.html
// 
// *****************************************************************************
func DisplayPage(w http.ResponseWriter, r *http.Request) {
    var p []Mst

    db, err := sql.Open("sqlite3", "wrk.db")
	checkErr(err)
	defer db.Close()

	rows, err := db.Query("SELECT Id, Dt, Nam, Lnk, Status FROM Links")
	checkErr(err)
	defer rows.Close()

    i:=0


	// Fetch record
	for rows.Next() {
		var id, dt, nam, lnk, status string
        
        // 
		err = rows.Scan(&id, &dt, &nam, &lnk, &status)
		checkErr(err)
        
        // Добавление записей  в масив для отчета
        p=append(p, Mst{"Id":id, "Dat":dt, "Nam":nam, "Lnk":lnk, "Status": status})
        i=i+1
     }
    
    
    fp := path.Join("tmp", "main.html")
	tmpl, err := template.ParseFiles(fp)

	// Error
	if err != nil {
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
	}

	tmpl.Execute(w, p) 
}



// *****************************************************************************
// Удаление очистка всей таблицы
// и возврат
// *****************************************************************************
func Clear_docs(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DELETE")

	EXC("DELETE FROM 'userinfo' ")
	// EXC("DELETE FROM 'userinfo'")
	http.Redirect(w, r, "/docscan/", 301)
}


// *****************************************************************************
// Удаление по ИД
// *****************************************************************************
func Delonerec(w http.ResponseWriter, r *http.Request) {
	 EXC("DELETE FROM 'userinfo' WHERE id=" + Lr(r, "delone"))
}


// *****************************************************************************
//
// *****************************************************************************
func Lr(r *http.Request, T string) string {
	id := r.URL.Path[len(T)+2:]
	return id
}


// *****************************************************************************
// Обновление одной записи
// *****************************************************************************
func Updaterec(w http.ResponseWriter, r *http.Request) {
	// id:=r.URL.Path[len("/updaterec/"):]

	id := Urp(r, "/updaterec/")
	EXC("UPDATE 'userinfo' SET username='зависимости' WHERE id=" + id)
	fmt.Println("Update", id)
}


// *****************************************************************************
// Получение параметра
// *****************************************************************************
func Urp(r *http.Request, Len string) string {
	return r.URL.Path[len(Len):]
}



// *****************************************************************************
// Date : 01.08.2017 19:12
// Возврат одной записи
// *****************************************************************************
func Retrec(w http.ResponseWriter, r *http.Request) {
	var Dt vv
	id := r.URL.Path[len("/returnrec/"):]

	if id == "" {
	   fmt.Fprintln(w, "ERROR ID NUMBER. Cannot be null.")
	   return
	}

	var idd, name, dep, dat string

	db, err := sql.Open("sqlite3", "wrk.db")
	checkErr(err)

	errs := db.QueryRow("SELECT id, username, departname, Dateinsert FROM userinfo WHERE id="+id).Scan(&idd, &name, &dep, &dat)
	checkErr(errs)

	// Control output data
	// fmt.Println(name,"\n",idd,"\n", dep,"\n", dat)
	// Load data to structure
	Dt.Name = name
	Dt.Dep  = dep
	Dt.Idd  = idd
	Dt.Dat  = dat

	// Variant 2
	// Конструкция для Mst
	// response:=Mst{"name":name}
	data, err := json.Marshal(Dt)
	checkErr(err)

	// fmt.Fprintf(w, &Dt)
	w.Write(data)

	// Error
	// w.Write([]byte(Dt))
}



// *****************************************************************************
// Date : 01.08.2017 19:12
// Добааление новой записи
// *****************************************************************************
func AddLinks(w http.ResponseWriter, r *http.Request) {

	// Fields for form
	s := r.FormValue("Ses")
	n := r.FormValue("Nam")
	// d := r.FormValue("Dt")
	t := r.FormValue("Status")
	l := r.FormValue("Lnk")
	j := r.FormValue("Descript")

	

	if s != "Secret" {
		log.Println("INFO : NO SECRET")
		http.Redirect(w, r, "/no.html", 301)
	}

	// Добавление 10 дней до следующей выполненной задачи
	// x := TimeStrAdd(10)

	// Добавлена новая запись в базу данных
	Db_add_link(n, t, l, j)

	// Запись в лог файл
	log.Println("ADD REC : ", n, l,j)

	// После записи возвращаемся обратно в форму
	http.Redirect(w, r, "/links.html", 301)

	// http.Redirect(w, r, "/docscan/", 301)
	// http://localhost:5555/docscan/
}


// *****************************************************************************
// Добавление записи в базу данных
// *****************************************************************************
func Db_add_link(Nam, Status, Lnk, Descript string) {
	db, errdb := sql.Open("sqlite3", "wrk.db")
	checkErr(errdb)
	defer db.Close()
    db.Exec("PRAGMA synchronous=OFF;")

	// insert
	stmt, errs := db.Prepare("INSERT INTO Links(Flag, Nam, Dt, Status, Lnk, Descript) VALUES (?,?,?,?,?,?)")
	checkErr(errs)

    // Active
    f:="A"

	// Добавление новой записи в базу
	stmt.Exec(f, Nam, TimeStr(), Status, Lnk, Descript)
    log.Println("Record for links saved")  	
}



// *****************************************************************************
// Сканирование документов из базы данных
// И сохранение в базу
// в зависимости от необходимости
// *****************************************************************************
func Db_lnk_report(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "wrk.db")
	checkErr(err)
	defer db.Close()

	rows, err := db.Query("SELECT id, Dt, Nam, Lnk FROM Links")
	checkErr(err)
	defer rows.Close()

	tr := ""

	// Fetch record
	for rows.Next() {
		var id, dt, nam, lnk string
		err = rows.Scan(&id, &dt, &nam, &lnk )
		checkErr(err)

		tr +=`<tr> 
         			<td style='text-align:center; width:50px;'>  <input id="` + id + `" type="checkbox"></td> 
                    <td style='text-align:center; width:50px;'>  <i onclick="Mmodal(` + id + `); Insertdata([` + id + `,'` + dt + `','` + nam + `','` + lnk + `']);" style='cursor:pointer;' class='fa fa-cog fa-lg'></i></td> 
                    <td style='text-align:center; width:40px;'>  `  + id + `</td> 
                    <td style='text-align:center; width:150px;'> ` + dt + `</td> 
                    <th >` + nam + `</th> 
                    <td style='text-align:center; width:50px; '> <i onclick='delrec(` + id + `);' style='cursor: pointer;' class='fa fa-trash-o fa-lg'></i> </td> 
                    <td style='text-align:center; width:40px; '> <a href='` + lnk + `'>Link</a> </td> 
              </tr>`
	}


	bd := h2 + tr + foo
	fmt.Fprintln(w, bd)

	err = rows.Err()
	checkErr(err)
}



// *****************************************************************************
// Date : 01.08.2017 19:12
// Добааление новой записи
// *****************************************************************************
func Addform(w http.ResponseWriter, r *http.Request) {

	// Fields for form
	s := r.FormValue("Ses")
	f := r.FormValue("Fam")
	n := r.FormValue("Nam")
	d := r.FormValue("Dt")
	l := r.FormValue("Lnk")
	t := r.FormValue("Status")
	p := r.FormValue("Project")
	j := r.FormValue("Description")

	fmt.Println("Description", j)

	if s != "Secret" {
		log.Println("INFO : NO SECRET")
		http.Redirect(w, r, "/no.html", 301)
	}

	// Добавление 10 дней до следующей выполненной задачи
	x := TimeStrAdd(10)

	// Добавлена новая запись в базу данных
	Db_add(f, n, d, l, t, p, j, x)

	// Запись в лог файл
	log.Println("ADD REC : ", f, n, d, l)

	// После записи возвращаемся обратно в форму
	http.Redirect(w, r, "/card.html", 301)

	// http.Redirect(w, r, "/docscan/", 301)
	// http://localhost:5555/docscan/
}




// *****************************************************************************
// Добавление записи в базу данных
// *****************************************************************************
func Db_add(N, D, C, L, S, P, J, X string) {
	fmt.Println(J)
	T := TimeStr()
	EXC("INSERT INTO userinfo(Dateinsert, username, departname, created, Lnk, Status, Project, Descript, End) VALUES ('" + T + "','" + N + "','" + D + "','" + C + "','" + L + "','" + S + "','" + P + "','" + J + "','" + X + "')")
	prn("INFO", "REC ADDED IN TABLE"+S+" : "+P)
}


// *****************************************************************************
// Добавление записи в базу данных
// Старый вариант - длинный
// *****************************************************************************
func Db_add_rec(Name, Depart, Create, Link string) {

	db, errdb := sql.Open("sqlite3", "wrk.db")
	checkErr(errdb)
	defer db.Close()

	// insert
	stmt, errs := db.Prepare("INSERT INTO userinfo(Dateinsert, username, departname, created, Lnk) VALUES (?,?,?,?,?)")
	checkErr(errs)

	// Добавление новой записи в базу
	stmt.Exec(TimeStr(), Name, Depart, Create, Link)
}




// *****************************************************************************
// Сканирование документов из базы данных
// И сохранение в базу
// в зависимости от необходимости
// *****************************************************************************
func DbScan(w http.ResponseWriter, r *http.Request) {
     Limit := r.URL.Path[len("/docscan/"):]

     if Limit=="" {
        Limit="20" 
    }


	db, err := sql.Open("sqlite3", "wrk.db")
	checkErr(err)
	defer db.Close()

 //db.Exec("PRAGMA synchronous=OFF;")

	// rows, err := db.Query("SELECT id, username, departname, Dateinsert, Lnk, Status, Project FROM userinfo WHERE Project='Дом' ORDER BY id DESC LIMIT "+Limit)
	rows, err := db.Query("SELECT id, username, departname, Dateinsert, Lnk, Status, Project FROM userinfo  ORDER BY id DESC LIMIT "+Limit)
	checkErr(err)
	defer rows.Close()

	tr := ""

	// Fetch record
	for rows.Next() {
		var id, name, dep, tm, lnk, status, project string
         
        // Ошибку вызывает пустое поле в базе
		err = rows.Scan(&id, &name, &dep, &tm, &lnk, &status, &project)
		checkErr(err)
		// rows.Scan(&id, &name, &dep, &tm, &lnk, &status, &project)

		tr += `<tr> 
         			<td style='text-align:center; width:50px; '><input id="` + id + `" type="checkbox"></td> 
                    <td style='text-align:center; width:50px; '><i onclick="Mmodal(` + id + `); Insertdata([` + id + `,'` + dep + `','` + name + `','` + lnk + `','` + tm + `']);" style='cursor:pointer;' class='fa fa-cog fa-lg'></i></td> 
                    <td style='text-align:center; width:40px; '>` + id + `</td> 
                    <td style='text-align:center; width:150px;'>` + tm + `</td> 
                    <th >` + name + `</th> 
                    <td >` + dep + `</td> 
                    <td >` + project + `</td> 
                    <td style='text-align:center; width:100px;'>` + Sts(status) + `</td>
                    <td style='text-align:center; width:50px; '> <i onclick='delrec(` + id + `);' style='cursor: pointer;' class='fa fa-trash-o fa-lg'></i> </td> 
                    <td style='text-align:center; width:40px; '> <a href='` + lnk + `'>Link</a> </td> 
                </tr>`
	}

	bd := h + tr + foo
	fmt.Fprintln(w, bd)

	// err = rows.Err()
	// checkErr(err)
}



// *****************************************************************************
// определение статуса по номеру
// *****************************************************************************
func Sts(Num string) string {
	t := []string{"Старт", "В работе", "Выполненно", "Планирование", "Откланено", "Важно", "Архив"}
	n, _ := strconv.Atoi(Num)
	return Stsl(Num, t[n])
}


// *****************************************************************************
// Покраска статусов в разный цвет для Bootsrup
// *****************************************************************************
func Stsl(Num, Text string) string {
	t := []string{"default", "primary", "success", "info", "warning", "danger"}
	n, _ := strconv.Atoi(Num)
	s := t[n]
	return `<span class="label label-` + s + `">` + Text + `</span>`
}


// *****************************************************************************
// Количество добавленных дней к выполнению задачи
// *****************************************************************************
func TimeStrAdd(Days int) string {
	Ttstr := time.Now().AddDate(0, 0, Days).Format("2006-01-02 15:04:05")
	return Ttstr
}

// *****************************************************************************
// Дата и время текущеее
// *****************************************************************************
func TimeStr() string {
	Ttstr := time.Now().Format("2006-01-02 15:04:05")
	return Ttstr
}

// *****************************************************************************
// Исполнитель запросов
// *****************************************************************************
func EXC(StrSql string) {
	db, err := sql.Open("sqlite3", "wrk.db")
	checkErr(err)
	defer db.Close()
    
    // db.Exec("PRAGMA journal_mode=MEMORY;")
    // db.Exec("PRAGMA locking_mode=EXCLUSIVE;")
    // db.Exec("PRAGMA temp_store=MEMORY;")

	db.Exec("PRAGMA synchronous=OFF;")

	ss, erre := db.Exec(StrSql)
	checkErr(erre)
	// prn("EXEC", StrSql)
	fmt.Println(ss)
}


// *****************************************************************************
// Обработка ошибок
// *****************************************************************************
func checkErr(err error) {
	if err != nil {
	   log.Println("ERROR : ", err)
	}
}


// *****************************************************************************
// Вывод сообщений
// *****************************************************************************
func prn(Notif, Text string) {
	 log.Println(Notif+" : ", Text)
}


// *****************************************************************************
// Пример вставки 10000
// *****************************************************************************
func TestInsertRecord(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "wrk.db?mode=ro")
	checkErr(err)
	defer db.Close()

	s := time.Now()
	count := 100000

     // db.Exec("PRAGMA journal_mode=MEMORY;")
     // db.Exec("PRAGMA locking_mode=EXCLUSIVE;")
     // db.Exec("PRAGMA temp_store=MEMORY;")
     db.Exec("PRAGMA synchronous=OFF;")

     // db.Exec("PRAGMA journal_mode=OFF;")

    

	for i := 0; i < count; i++ {
         
       //   stmt, errs := db.Prepare("INSERT INTO userinfo(username) VALUES (?)")
	      // checkErr(errs)
       //    stmt.Exec("естирование",true)


             
		      db.Exec("INSERT INTO userinfo(username,departname) VALUES ('Test','Department')")
             // log.Println("Всnfdrf")	    
	}

	f   := time.Now()
	rez := f.Sub(s)

	fmt.Println("Всего вреемни потрачено 1 процедурой	:", rez)
}




// *****************************************************************************
// Пример вставки 10000
// *****************************************************************************
func TestInsertRecord2(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "wrk.db?mode=ro")
	checkErr(err)
	defer db.Close()

	s := time.Now()
	count := 100

     // db.Exec("PRAGMA journal_mode=MEMORY;")
     // db.Exec("PRAGMA locking_mode=EXCLUSIVE;")
      db.Exec("PRAGMA temp_store=MEMORY;")
      db.Exec("PRAGMA journal_mode=OFF;")

     // Использовать обязательно
     // Без этой опции вставка происходит ооооооочень медленно !!!!
     db.Exec("PRAGMA synchronous=OFF;")

     

    
    // Loop
	for i := 0; i < count; i++ {
         
          // stmt, errs := db.Prepare("INSERT INTO userinfo(username) VALUES (?)")
	      // checkErr(errs)
          // stmt.Exec("естирование",true)
		  db.Exec("INSERT INTO userinfo(username, departname, Dateinsert,Lnk) VALUES ('Test 2','Department 2', '2019-11-11','')")
          // log.Println("Всnfdrf")	    
         
	}

	f   := time.Now()
	rez := f.Sub(s)



    // 26.28883072s Вставлено  1000000  Записей
	fmt.Println("Всего вреемни потрачено  второй:", rez, "Вставлено ", count, " Записей")
}
