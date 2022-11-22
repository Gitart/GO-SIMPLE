// ******************************************************************************************************
// 
// Просмотр всей информации
// 
// 
// ******************************************************************************************************
func Strview(w http.ResponseWriter, rr *http.Request) {

   var Recs Ms 
   rk,_:=r.DB("test").Table("Temp").Run(sessionArray[0])

    thtm:=`
          <html>
          <head>
             <meta charset="utf-8">
             <meta http-equiv="X-UA-Compatible" content="IE=edge">
             <meta name="viewport" content="width=device-width, initial-scale=1">
             <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/css/bootstrap.min.css">
             <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/css/bootstrap-theme.min.css">
             <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/js/bootstrap.min.js"></script>
          </head>

          <body>
          <div class="container">
               <h3 style="clor:#CCF;">Список пример для просмотра центрального офиса.</h3>
               <a href="/view/">Возврат</a>
               
               <table  class="table table-hover table-bordered table-condensed"> 
               <tr> <td>Наименование</td>  <td>Описание</td> <td>Точка</td> <td>ID</td>  </tr> 
               `
    thtf:="</table></div></body></html>"

    fmt.Fprintf(w, "%s ",thtm )

    // Заполнение таблицы данным
    for rk.Next(&Recs) {
	      fmt.Fprintf(w, "<tr> <td>%s</td>  <td>%s</td> <td>%s</td> <td>%s</td> </tr>", Recs.Dateinsert, Recs.Title, Recs.Code, Recs.Operation)
	  }
    
    // Вывод информации на страницу
    fmt.Fprintf(w, "%s ", thtf )
}
