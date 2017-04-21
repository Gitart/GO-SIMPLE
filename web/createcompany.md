

```golang
/******************************************************************************************************************************************************
 *
 *   Просмотр одного альянса
 *   http.HandleFunc("/api/system/aliance/viewone/",      Sys_aliance_add)               // Добавление нового альянса
 *
 *   Пример использования :
 *   http://10.10.10.24:5555/api/system/aliance/viewone/3
 *
 ******************************************************************************************************************************************************/
func Sys_aliance_viewID(wr http.ResponseWriter, rq *http.Request) {

	// Инициализация переменных
	var ndd []Aliance
	resp, err := r.DB("HO").Table("Aliance").OrderBy("NAME").Run(sessionArray[0])

	// Error
	if err != nil {
		log.Println(err)
	}

	resp.All(&ndd)

	Ptest := `<!DOCTYPE html>
	        <htmL>
	                <head>
				          <title>Head Office</title>
					      <meta http-equiv="Content-Language" content="en-us" />
					      <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
                          <meta name="viewport" content="width=device-width, initial-scale=1">
                          <link rel="stylesheet" href="http://maxcdn.bootstrapcdn.com/bootstrap/3.2.0/css/bootstrap.min.css">
                          <script src="https://ajax.googleapis.com/ajax/libs/jquery/1.11.1/jquery.min.js"></script>
                          <script src="http://maxcdn.bootstrapcdn.com/bootstrap/3.2.0/js/bootstrap.min.js"></script>

                          <style type="text/css">
                                 body  {marging:10px;}
                                 html  {color: #3B3B3B; font-size:10px; font-family:Calibri;}
                          </style>
  				    </head>
	        <body>

	        <div style="left:50%;width:100%; position:absolute;border:1px silver solid; margin-left: -50%; padding:7px; marging-top:10px;">
		    	 <a href="#" class="list-group-item active" > Справочник альянсов</a> <br>
                                  <table class="table table-condensed">
		                                 <thead> <tr class="active">
		                                             <td>ID</td>
			                                         <td>Наименование</td>
			                                         <td>Город</td>
			                                         <td>EMAIL</td>
			                                         <td>SEQ</td>
			                                         <td>Поточный статус</td>
			                                         <td>Дата создания</td>
			                                         <td>Контактные телефоны</td>
			                                         <td>Дата исправления</td>
			                                      </tr>
		                                 </thead>

                                      {{range.}}
								            <tbody>
								               <tr>
								                   <td><b>{{printf "%v" .ID}}</b></td>
								                   <td><b>{{printf "%s" .FULLNAME}}</b></td>
								                   <td>   {{printf "%s" .CITY}}</td>
								                   <td>   {{printf "%s" .EMAIL}}</td>
								                   <td>   {{printf "%v" .SEQ}}</td>
								                   <td>   {{printf "%s" .STATUS}}</td>
								                   <td>   {{printf "%s" .CREATED}}</td>
								                   <td>   {{printf "%s" .TELEPHONE}}</td>
								                   <td>2014-02-03</td>
								                </tr>
								            </tbody>
								       {{end}}
								            </table>
            </div>
            </body>
            </htmL> `

	tmpl, err := template.New("test").Parse(Ptest)

	// Error
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(wr, ndd)
}
```
