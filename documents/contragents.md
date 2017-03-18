### Просмотр справочника контрагентов
```go

/******************************************************************************************************************************************************
 *	Просмотр справочника контрагентов
 *	http.HandleFunc("/api/system/users/viewall/",  Sys_users_viewall)
 *
 *	Пример использования :
 *	http://10.10.3.10:5555/api/system/users/viewall/key
 ******************************************************************************************************************************************************/
func Sys_users_viewall(wr http.ResponseWriter, rq *http.Request) {

	// Пользователи сситемы
	type USR struct {
		Id        string `gorethink:"id"` // Id (GUID)
		ID        string `gorethink:"Id"` // ID
		Ip        string // Имя комп
		Os        string // Os
		Name      string `gorethink:"Name"`  // Полное имя
		Lname     string `gorethink:"Lname"` // Фамилия
		Mname     string // Имя
		Fname     string // Отчество
		Telephone string // телефон рабочий
		Status    string // Должность
		Position  string // Ид структуры
		Structure int64
	}

	// Инициализация переменных
	var Usr USR

	// cnntt := 0
	Prf := `<div class='panel panel-default' style='margin:10px; padding:5px;'>
            <div class='panel-heading' style='color:#3399FF;'>
                 <b style='font-size:20px;'>HEAD OFFICE</b>
            </div>`

	Pre := `<br><div class="panel panel-info">
	                 <div class="panel-heading"> <b>Пользователи системы<b> </div>
	                 <div class="panel-body">
			<p>
			<a href="/static/newuser.html?corp=asdasd2232322344" target="_blank">
			  <button type="button" class="btn btn-default btn-sm">
					  <span class="glyphicon glyphicon-plus-sign" aria-hidden="true"></span>
					  Добавить
			  </button>
			</a>

			<a href="/static/userblock.html?corp=asdasd2232322344" target="_blank">
			  <button type="button" class="btn btn-default btn-sm">
					  <span class="glyphicon glyphicon-ban-circle" aria-hidden="true"></span>
					  Заблокировать
			  </button>
			</a>

			<a href="/static/users.html?corp=asdasd2232322344" target="_blank">
			  <button type="button" class="btn btn-default btn-sm">
					  <span class="glyphicon glyphicon-star" aria-hidden="true"></span>
					  Корректировать
			  </button>
			</a>
			<a href="/static/userdelete.html?corp=asdasd2232322344" target="_blank">
			  <button type="button" class="btn btn-default btn-sm">
					  <span class="glyphicon glyphicon-star" aria-hidden="true"></span>
					  Удалить
			  </button>
			</a>
			</p>`

	Prs := `<table class='table table-condensed table-hover'
		    <thead><tr> <th>...</th>
		           <th>Фамилия</th>
		           <th>Имя</th>
		           <th>Отчество</th>
		           <th>Должность</th>
		           <th>Статус</th>
	               <th>Структура</th>
	         </thead> <tbody> </div>`

	Hrw := `<tr>
	        <td> <input type='checkbox' id=A%s/>  <span class="%s"></span></td>
	        <td> <a target="_blank" href="/api/system/aliance/viewone/%s">%s</a></td>
	        <td>%v</td>
	        <td>%v</td>
	        <td> <span class="label label-success">%v</span></td>
	        <td> <span class="label label-info">%s</span></td>
	        <td> %v</td>
	        </tr>`
	Prh := `</tbody></table></div></body></div></div></html>`

	resp, err := r.DB("HO").Table("Users").OrderBy(r.Desc("ID")).Run(sessionArray[0])

	// Error
	if err != nil {
	   log.Println(err)
	}

	// Формирование страницы
	fmt.Fprintf(wr, HtmlSTR)
	fmt.Fprintf(wr, Prf)
	fmt.Fprintf(wr, Pre+Prs)

	// Заполнение тела формы записями
	for resp.Next(&Usr) {
		fmt.Fprintf(wr, Hrw, Usr.Id, Usr.Name, Usr.Id, Usr.Fname, Usr.Lname, Usr.Mname, Usr.Position, Usr.Status, Usr.Structure)
	}

	// Подножье - оформление
	fmt.Fprintf(wr, Prh)
}
```
