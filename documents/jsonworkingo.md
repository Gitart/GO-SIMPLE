## Работа с JSON & GO for RETHINKDB

### Необходимо подготовить файл в формате JSON
             [ {},{},{} ]   - Для пакетной вставки
             {}                - для построчной вставке  :

#### В программе необходимы изменения :
```
  // Формирование для нескольких документов в Json файле документа
	// var m []*Mst // Автоматически подходит для всех форматов Json

	// Формирование для одного документа
	// {}
	m := make(map[string]interface{})
	errj := json.Unmarshal([]byte(reads), &m)

	// Обработка ошибок
	if errj != nil {
		log.Fatalln(errj)
	}
```

### Проверить можно с помощью команды :
```bat
CURL -X POST http://10.0.3.24:5555/docs/ -T "test.json" -H "Content-Type: application/json; charset=utf-8;" 
```

#### Готовая реализация
```go
/ *********************************************************************************************************************************************************
// Запись информации в таблицу из JSON формате строки
// с добавлением дополнительного поля InsertTime
// StrJson = "`" + StrJson + "`"
// StrJsons = StrJson
// StrJson = `{"Id":"T0001", "Title":"Samples"}`
// DelFile := os.Args[4]    // A=добавление D=удаление
// WriteJsonString - WJS
// *********************************************************************************************************************************************************
func WriteJsonString(StrJson, NameBase, NameTables string) {

	// Для использования параметров
	// NameBase := os.Args[1]   // База
	// NameTables := os.Args[2] // Таблица
	// StrJson := os.Args[3]    // Строка JSON

	// Текущее время
	var CurTime = time.Now().Format("2006-01-02 15:04:05")

	// Read JSON
	m := make(map[string]interface{})
	byt := []byte(StrJson)
	errj := json.Unmarshal(byt, &m)

	// Control input
	//	fmt.Println(StrJson)

	// Обработка ошибок
	if errj != nil {
		log.Fatalln(errj)
	}

	// Добавление полей в базу автоматом
	// Например дату и время или пользователя
	// r.Db("test").Table("Post").Insert(r.Expr(m).Merge(Mst{"InsertTime": CurTime}).Merge(Mst{"Sequence": MaxID()})).Run(sessionArray[0])
	// ms := r.Expr(m).Merge(Mst{"InsertTime": CurTime}).Run(session)

	// Запись в базу
	r.Db(NameBase).Table(NameTables).Insert(r.Expr(m).Merge(Mst{"InsertTime": CurTime})).Run(session)
}



// Фильтрация документов по условию
// Условие для проверки работоспособности фильтра
// Условия необходимо соблюдать в рамках договоренности по формату
// Формат чисел оговорен в структуре
func FilterDataForDocumnets(w http.ResponseWriter, rr *http.Request) {
	// m := make(map[string]interface{}) // Автоматически подходит для всех форматов Json
	// var outm []*Mst         // Автоматически подходит для всех форматов Json
	var response []interface{} // вывод пачки документов

	var FilterDataCond = Mst{"ID_PARTNER_CORP": 1, "ID_PARTNER_CONTR": "5985645", "ID_PARTNER_STRUCT": "5985647", "ID_TYPE": 14}

	res, err := r.Db("test").Table("Docmove").Filter(FilterDataCond).Run(sessionArray[0])
	// Error
	if err != nil {
		log.Println("No open table for Import ...")
	}

	err = res.All(&response)

	if err != nil {
		fmt.Fprintf(w, "Sorri no Documents for You Conditional%s", strings.ToUpper("404"))
	} else {
		data, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json; carset=utf-8")
		//w.Write(data)
		fmt.Fprintf(w, string(data))
	}

}
```
