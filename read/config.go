/*********************************************************************************************************************
 *
 *   Чтение JSON файла и вывод в страницу напрямую
 *   /tst/configread/
 *
 *********************************************************************************************************************/
func Test_GetJson(w http.ResponseWriter, req *http.Request){
      
    // Открытие файла настройки
    // in Unix ./config.json
    file, e := ioutil.ReadFile("config.json")          

    // Error
    if e != nil {
       fmt.Printf("File error: %v\n", e)
       os.Exit(1)
    }
    
    // fmt.Printf("%s\n", string(file))
    // m := new(Dispatch)
    // var m interface{}
    var m Mst // Автоматически подходит для всех форматов Json

    // Формирование для одного документа
    json.Unmarshal([]byte(file), &m)


    fmt.Fprintln(w,`<html>
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
                    </Head>
                    </body>
                    <div class="container">
                    <h3>Основные настройки сервиса </h3><hr>`)
    
    fmt.Fprintln(w, "1. Level ")
    fmt.Fprintln(w, "2. PORT  : ...................", m["port"].(string)+"<br>")
    fmt.Fprintln(w, "3. IP    : ...................", m["ip"].(string)+"<br>")
    fmt.Fprintln(w, "4. Name  : ...................", m["names"].(string)+"<br>")
    fmt.Fprintln(w, "5. ID    : ...................", FloatToString(m["id"].(float64))+"<br>")
    fmt.Fprintln(w, "6. Count : ...................", m["kol"].(string)+"<br>")
    fmt.Fprintln(w, "7. Key   : ...................", m["SecKey"].(string)+"<br><br>")


    // NN **************************************************
    // nn:[{},{},{}]
    ll:=m["nn"].([]interface{})
    // ls:=ll[0]
    // fmt.Println(ls)
    // ld:=ls.(map[string]interface{})
    // fmt.Println(ld["name"], ld["id"].(float64))
    // Количество элементов в массиве
    Cnt:=len(ll)
    lrs:=""
   

    // Прокрутка элементов в массиве nn
    for i:=0; i<Cnt; i++ { 
        ly:=ll[i].(map[string]interface{})
        lz:=ly["id"]
        lk:=ly["name"].(string)
        lr:=ll[i].(map[string]interface{})["name"]
        fmt.Println(ly, lz, lk, " Имя : ", lr)
        lrs=lrs + "        <li><a href='"+lk+"'>Хозяин собаки : " + lk + "</a></li>\n"
    }
  
     // LINKS **************************************************
     fmt.Fprintln(w, lrs, "<hr><h2>Полезные линки </h2>")
     ll=m["Links"].([]interface{})
     // Количество элементов в массиве
     Cnt=len(ll)
     lrs=""

   
    // Прокрутка элементов в массиве Links
    for i:=0; i<Cnt; i++ { 
         fmt.Println(i,"")
         ly:=ll[i].(map[string]interface{})
         lz:=ly["Title"].(string)
         lk:=ly["Lnk"].(string)
         // lr:=ll[i].(map[string]interface{})["Lnk"]
         // fmt.Fprintln(w, ly, lz, lk, " Имя : ", lr)
         lrs=lrs + "        <li><a href='"+lk+"'>"+lz+"</a></li>\n"
       }


     // POSTS **************************************************
     tyl :=m["Posts"].([]interface{})
     Cnt  =len(tyl)

     // Обход элементов в массиве c link
     for i:=0; i<Cnt; i++ { 
            
              // TITLE **************************************************
              ly:=tyl[i].(map[string]interface{})                       // Инициализация ветки
              lt:=ly["Title"].(string)                                  // Огловление  
              ll:=ly["Links"].([]interface{})                           // Инициализация ветки - потомка   

              // Lnkstr=Lnkstr + "      <li><a href='"+ll+"'>Уроки : "+lt+"</a></li>\n"
              fmt.Fprintln(w,"<h3> ",i,". Tема : ", lt, "</h3> <br>")
            
              //Подсчет линков в ветке 
              Cnts:=len(ll)

              //  LINKS **************************************************
              //  Обход линков и наименований
              //  Третий уровень
              for y:=0; y<Cnts; y++{
                  llz:= ll[y].(map[string]interface{})
                  // fmt.Fprintln(w,"<a href=''     * ",i,".",y,".", llz["Title"], "  ", llz["Lnk"],"/a> ")
                  fmt.Fprintln(w,"<a href='", llz["Lnk"], "'>",  llz["Title"], "</a><br>")
              } 

             fmt.Fprintln(w,"<hr>")
        }
             fmt.Fprintln(w, lrs)
             fmt.Fprintln(w, "</div></body></html>")
}
