## Пример работы с возвратом ошибок в GO

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














