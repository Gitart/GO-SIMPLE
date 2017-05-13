
## Create form from json


```golang
/*********************************************************************************************************************
 *
 *   Чтение JSON файла и вывод в страницу напрямую
 *   В данном случае чтение сonfig.json
 *   /tst/configread/
 *
 *********************************************************************************************************************/
func Test_GetJson(w http.ResponseWriter, req *http.Request) {

	// Открытие файла настройки  (in Unix ./config.json)
	file, e := ioutil.ReadFile("config.json")

	// Error
	if e != nil {
	   fmt.Printf("File error: %v\n", e)
	   os.Exit(1)
	}

	// Initialization
	// fmt.Printf("%s\n", string(file))
	// m := new(Dispatch)
	// var m interface{}

	// Автоматически подходит для всех форматов Json
	var m Mst

	// Формирование для одного документа
	json.Unmarshal([]byte(file), &m)

	// Header
	Headname := m["Headname"].(string)
	Namedes  := m["Namedes"].(string)
	Descript := m["Descript"].(string)

	fmt.Fprintln(w, `<html>
				     <Head>
					    <title>Head Office</title>
						<meta charset="utf-8">
						<meta name="viewport" content="width=device-width, initial-scale=1">
					    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/css/bootstrap.min.css">
						<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/css/bootstrap-theme.min.css">
						<script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/js/bootstrap.min.js"></script>
					    <script src="https://ajax.googleapis.com/ajax/libs/jquery/1.11.3/jquery.min.js"></script>
						<script src="http://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/js/bootstrap.min.js"></script>
						<script src="http://code.jquery.com/ui/1.10.4/jquery-ui.js"></script>

						<link rel="stylesheet" type="text/css" href="/static/U2.css" />

                        <!--Hack for Bootstrup-->
						<style>
						     .table-condensed>tbody>tr>td,
						     .table-condensed>tbody>tr>th,
						     .table-condensed>tfoot>tr>td,
						     .table-condensed>tfoot>tr>th,
						     .table-condensed>thead>tr>td,
						     .table-condensed>thead>tr>th   {padding: 2px; font-size:14px; margin:2px; vertical-align:middle; padding-left: 10px;}
						</style>
				    </Head>
				  	</body>
				  	<div class="container">
				  	<h3>` + Headname + `</h3><hr>
				  	<div class='panel panel-default'> <div class='panel-heading'> <b>` + Descript + ` </b></div>
                    <small>` + Namedes + `</small>
				  	</div>`)

	fmt.Fprintln(w, "<table class='table table-bordered table-condensed'>")
	fmt.Fprintln(w, "<tr><td>Наименование </td> <td>" + m["port"].(string) + "</td></tr>")
	fmt.Fprintln(w, "<tr><td>Адрес        </td> <td>" + m["ip"].(string) + "</td></tr>")
	fmt.Fprintln(w, "<tr><td>IP Adress    </td> <td>" + m["names"].(string) + "</td></tr>")
	fmt.Fprintln(w, "<tr><td>Name         </td> <td>" + FloatToString(m["id"].(float64)) + "</td></tr>")
	fmt.Fprintln(w, "<tr><td>Fields       </td> <td>" + m["kol"].(string) + "</td></tr>")
	fmt.Fprintln(w, "<tr><td>Keys         </td> <td>" + m["SecKey"].(string) + "</td></tr>")
	fmt.Fprintln(w, "</table>")

	// NN ******************************************************************************************************
	// nn:[{},{},{}]
	ll := m["nn"].([]interface{})
	// ls:=ll[0]
	// fmt.Println(ls)
	// ld:=ls.(map[string]interface{})
	// fmt.Println(ld["name"], ld["id"].(float64))
	// Количество элементов в массиве
	Cnt := len(ll)

	// Прокрутка элементов в массиве nn
	fmt.Fprintln(w, "<div class='panel panel-default'> <div class='panel-heading'> <b>Прочие</b></div></div>")
	fmt.Fprintln(w, "<table class='table table-bordered table-condensed'>")

	for i := 0; i < Cnt; i++ {
		ly := ll[i].(map[string]interface{})
		lz := FloatToStr(ly["id"].(float64))
		lk := ly["name"].(string)
		// lr:=ll[i].(map[string]interface{})["name"]
		// fmt.Println(ly, lz, lk, " Имя : ", lr)
		fmt.Fprintln(w, "<tr><td style='width:40px;'>"+lz+"</td>  <td><a href='"+lk+"'> "+lk+"</a></td></tr>")
	}
	fmt.Fprintln(w, "</table>")

	// POSTS ***********************************************************************************************
	tyl := m["Posts"].([]interface{})
	Cnt = len(tyl)

	// Обход элементов в массиве c link
	for i := 0; i < Cnt; i++ {

		// TITLE **************************************************
		ly := tyl[i].(map[string]interface{}) // Инициализация ветки
		lt := ly["Title"].(string)            // Огловление
		ll := ly["Links"].([]interface{})     // Инициализация ветки - потомка

		// Lnkstr=Lnkstr + "      <li><a href='"+ll+"'>Уроки : "+lt+"</a></li>\n"
		fmt.Fprintln(w, "<br><div class='panel panel-default'> <div class='panel-heading'> <b>", i, ".   ", lt, "</b></div></div>")

		//Подсчет линков в ветке
		Cnts := len(ll)


		//  LINKS **************************************************
		//  Обход линков и наименований
		//  Третий уровень

		fmt.Fprintln(w, "<table class='table table-condensed'>")

		for y := 0; y < Cnts; y++ {
			llz := ll[y].(map[string]interface{})

			fmt.Fprintln(w, "<tr><td style='width:60px;' >"+
				InttoStr(i)+"."+
				InttoStr(y)+"</td>"+
				"<td style='width:800px;'><a href='"+llz["Lnk"].(string)+"'>"+
				llz["Title"].(string)+"</a></td> <td>"+
				llz["Notes"].(string)+"</td></tr>")
		}

		fmt.Fprintln(w, "</table>")
	}

	// LINKS 
	fmt.Fprintln(w, "<div class='panel panel-default'> <div class='panel-heading'> <b>Дополнительные линки и сноски на документы</b></div></div>")

	ll  = m["Links"].([]interface{})
	Cnt = len(ll) // Количество элементов в массиве

	fmt.Fprintln(w, "<ul>")

	// Прокрутка элементов в массиве Links
	for i := 0; i < Cnt; i++ {
		ly := ll[i].(map[string]interface{})
		lz := ly["Title"].(string)
		lk := ly["Lnk"].(string)
		// lr:=ll[i].(map[string]interface{})["Lnk"]
		// fmt.Fprintln(w, ly, lz, lk, " Имя : ", lr)
		fmt.Fprintln(w, "<li><a href='"+lk+"'>"+lz+"</a></li>")
	}

	fmt.Fprintln(w, "</ul>")
	fmt.Fprintln(w, "<div class='bs_02'> <p >All right resirved.<br> Киев "+CTM()+" г. </p></div>")
	fmt.Fprintln(w, "</div></body></html>")
}
```

