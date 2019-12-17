## Пример работы с возвратом ошибок в GO
В примере показана работа с формой и обработки формы на сервере.   
Внимание обращено на возврат ошибок и корректной обработки данных.  


```js
<script type="text/javascript">

// Insert to database
function Addnews(){
                $.ajax({type: 'POST',
                        url:  '/reiting/add/',
                        data: $('#formentry').serialize(),
                        success:function(response){
                                 Showalert(response);
                                 $.notify(response);
                                // $('#formentry').find('.formre').html(response);
                                // $('#tabledata').load('/api/report/plan/ #tabledata tbody');
                                // $('#formentry').hide();
                        },
                        error:function(error) {
                           
                           // Check return resultate in case error
                           // console.log(JSON.stringify(error));
                           
                           // alert(error.responseText); 
                              Showalert("Неправильные или не полные данные.");
                              $.notify(error.responseText);
                        }
                      });
                  
                  // Reset data in form    
                  document.getElementById("formentry").reset();
                  
                  // Second variant message about rigth insert to database
                  // Showalert("Рейтинг добавлен в базу данных");
                                  
}

// Показ сообщения
function Showalert(Text){
       $('#sret').html(Text);
       $('#myalert').show(100).delay(5000).hide(200);
}

// Показ сообщения
function ShowBadAlert(Text){
       $('#sret').html(Text);
       $('#myalert').show(100).delay(5000).hide(200);
}

</script>
```


## Html code
```html
<body >

<!--Body page-->
<div class="container">

             <h1 class="clrhead"><i class='fas fa-satellite-dish'></i> {{.Title}}</h1>
             <hr>

             <div class="row">
                  <div class="col-sm-6 col-lg-6 mb-3">
                         <h3> Alexa <span class="cif">{{.Alexa.Alexa}}</span> Ukraine <span class="cifukr">{{.Alexa.Ukraine}}</span></h3>
                  </div>

                  <div class="col-sm-6 col-lg-6 mb-3">
                        <span>Id: {{.Alexa.Id}} | Date: {{.Alexa.Date}} | Note : {{.Alexa.Descr}}  </span>
                        <p>Результаты SEO, SEM и контент-маркетинга</p>
                 </div>
             </div>

             <!--
             <h3> Alexa <span class="cif">{{.Alexa.Alexa}}</span> Ukraine <span class="cifukr">{{.Alexa.Ukraine}}</span></h3>
             <br>
             <span>Id: {{.Alexa.Id}} | Date: {{.Alexa.Date}} | Note : {{.Alexa.Descr}}  </span>
             <p>Результаты SEO, SEM и контент-маркетинга</p>
             <hr>
             -->

             <!--Notification add to database-->
             <div id="myalert" class="alert alert-success" style="display: none;">
  	              <strong>Success!</strong> <span id="sret"><i class='far fa-paper-plane'>
                  </i> Рейтинг добавлен в базу.</span>
    	       </div>

            <form  id="formentry" class="windowpopup" role="form" >
                       <div class="col-md-12">
                          <input id="idrec" name="idrec" type="hidden" >

                          <div class="form-group"> <h3 style="color:#C15A3C;"><i class='far fa-paper-plane'></i> {{.Description}}<b id="numtask"></b></h3><hr> </div>
                                                 
                          <div class="form-row">
                               <label class="control-label control-label-left col-sm-2">Alexa</label>
                               <div class="controls col-sm-9"> <input id="alexa" name="alexa" type="text" class="form-control" value="{{.Dat.Alexa}}" autofocus></div>
                          </div>
                          
                          <div class="form-row">
                               <label class="control-label control-label-left col-sm-2">Ukraine</label>
                               <div class="controls col-sm-9"> <input id="ukraine" name="ukraine" type="text" class="form-control" value="{{.Dat.Ukraine}}"> </div>
                          </div>

                          <div class="form-row">
                               <label class="control-label control-label-left col-sm-2">Описание</label>
                               <div class="controls col-sm-9"> <input id="descr" name="descr" type="text" class="form-control" value="{{.Dat.Descr}}"> </div>
                          </div>

                          <div class="form-row">
                               <label class="control-label control-label-left col-sm-2">Дата</label>
                               <div class="controls col-sm-9"><input id="date" name="date" type="text" class="form-control" value="{{if .Id}}{{.Dat.Date}}{{else}} {{.Tm}} {{end}} "></div>
                          </div>
                                                 
                        <hr>
                        <div class="form-row">
                            <button id="but_update" type="button" class="btn btn-sm btn-success" onclick="Addnews()">Добавить рейтинг</button>
                            <a id="but_home"  class="btn btn-sm btn-info" href="/">Домашняя</a>
                            <a id="but_list"  class="btn btn-sm btn-secondary" href="/reiting/report/">Список рейтингов</a>
                        </div>
                   </div>
          </form>
</div>
```

## In GO function

### Start programm
```golang
package main

import (
  "context"
  "encoding/base64"
  "flag"
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "os"
  "os/signal"
  "strings"
  "syscall"
  "time"
 
  _ "github.com/jinzhu/gorm/dialects/sqlite"
)

// *******************************************************
// Start main procedure
// *******************************************************
func main() {
  // Restore in bad case
  defer func() {
    if r := recover(); r != nil {
      fmt.Println("Recovered in f", r)
    }
  }()

  // Set Logs
  f, err := os.OpenFile("log/log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
  if err != nil {
    fmt.Printf("Error opening file: %v", err)
  }

  defer f.Close()
  log.SetOutput(f)

  // Flags
  Port := flag.String("p", "1968", "Input Port") // Port by default
  // View := flag.String("v", "y",    "Previe error")

  flag.Parse()
  
  
  // Рейтинг
  http.HandleFunc("/reiting/",            Add_reiting)                // Форма       - Добавление рейтинга
  http.HandleFunc("/reiting/add/",        Add_new_raiting)            // Процедура   - Добавление рейтинга
  http.HandleFunc("/reiting/report/",     All_reitings)               // Отчет рейтингов



  // log.Fatal(sr.ListenAndServe())
  fmt.Printf(Pr, *Port, "2.003", Ct())

  // log.Println(http.ListenAndServe(":1968", nil))
  // Err(err, "Error start service.")

  // Settings portal
  srv := &http.Server{Addr:         ":" + *Port,
                      IdleTimeout:  120 * time.Second,
                      ReadTimeout:  10  * time.Second,
                      WriteTimeout: 10  * time.Second,
  }

  // srv :=makeServerFromMux()
  // Start Server
  go func() {
    FgGreen("\n        Start servers......")

    if err := srv.ListenAndServe(); err != nil {
      log.Fatal(err)
    }
  }()

  // Graceful Shutdown
  waitForShutdown(srv)
}


// ***************************************
// Graceful Shutdown server
// ***************************************
func waitForShutdown(srv *http.Server) {

  interruptChan := make(chan os.Signal, 1)
  signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

  // Block until we receive our signal.
  <-interruptChan

  // Create a deadline to wait for.
  ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
  defer cancel()
  srv.Shutdown(ctx)

  // Notify shutdown server
  FgRed("Shutting down server.....\n")
  os.Exit(0)
}
```


## Request form with template
```golang
/*
 * Форма для добавление рейтинга
 * /reiting/
 * Add_reiting
 */
func Add_reiting(w http.ResponseWriter, r *http.Request) {

	idp   := r.URL.Path[len("/reiting/"):]
	id    := Str2Int64(idp)      // 
	Dat   := GetOneRating(id)    //
	Alexa := GetOneRatingMax()   // Cвежий рейтинг сайта

	// Дата
	Dt := Mst{
		"Title":       "Рейтинг Alexa & Ukraine",
		"Description": "Добавление текущего ретинга сайта",
		"Dat":         Dat,
		"Tm":          time.Now().Format("02.01.2006 15:04:05"),
		"Id":          idp,
		"Alexa":       Alexa,
	}


    RenderHtml("addreiting.html",Dt,w)
}


/*
 * Render html pages
 */
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











### Insert procedure
```golang
package main
import (	
    "fmt"
    "log"
    "time"
    "net/http"
    "strings"
    // "github.com/fatih/color"
    "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
)

//**************************************************
// Db connect
//**************************************************
func Dbc()  *gorm.DB {
     // log.Println("Connecting to database...")
     db, err := gorm.Open("sqlite3", "db/zorg.db")
     if err!=nil{
        fmt.Println("Error connect to db", err.Error())
     }
     // defer db.Close()
     return db
}

// Рейтинг структура
type Ratings struct{
     Id, Alexa, Ukraine, Date, Descr string
}

//******************************************************************
//  Добавление рейтинга в журнал рейтингов
//  /reiting/add/
//******************************************************************
func Add_new_raiting(w http.ResponseWriter, r *http.Request) {
     var Rating Ratings
     tm := time.Now()
     np := tm.Format("02-01-2006 15:04:05")
     db := Dbc();defer db.Close()
     dt := strings.Trim(r.FormValue("date")," ")


     Rating.Alexa     = r.FormValue("alexa")              // By World 
     Rating.Ukraine   = r.FormValue("ukraine")            // By Ukraine
     Rating.Descr     = r.FormValue("descr")              // Description 

     // Chek input date    
     if dt=="" || len(dt)==0  {
        fmt.Println("No data........")
        Rating.Date   = np
        
        // Возврат ошибки в случае пустой даты
        http.Error(w, "Bad or incorrect date", 400)
        //w.Write([]byte("С пустой датой нельзя заполнять "))
        fmt.Println("нельзя retings")
        return 
     }
        Rating.Date   = r.FormValue("date")            
        db.Create(&Rating)

        // AddNews (title, body, header, note, grp)
        w.Write([]byte("Добавлен рейтинг "))
        fmt.Println("Add new retings")
}
```

## Database SqLite

```sql
CREATE TABLE "ratings" ( `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, 
                         `alexa` TEXT, `ukraine` TEXT, `date` TEXT, `descr` TEXT )
```














