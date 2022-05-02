
/***************************************************************************
Stage to project
***************************************************************************/
func Stage_project(w http.ResponseWriter, r *http.Request)   {
    log.Println("Start stage ")
    var Dt []Mst
    St:=`[ 
         {"id":1,  "Num":"1.",   "Stage":"Проект",                      "Descript": "Потенційний проект"},
         {"id":2,  "Num":"2.",   "Stage":"Передпродаж",                 "Descript": "Передпродаж"},
         {"id":3,  "Num":"3.1.", "Stage":"Початок. Збір вимог",         "Descript": "Вимоги"},
         {"id":4,  "Num":"3.2.", "Stage":"Початок. Узгодження КП",      "Descript": "Узгодження КП"},
         {"id":5,  "Num":"3.3.", "Stage":"Початок. Узгодження угоди",   "Descript": "Договір"},
         {"id":6,  "Num":"4.1.", "Stage":"Уточнення. Архітектура",      "Descript": "Архитектура"},
         {"id":7,  "Num":"4.2.", "Stage":"Уточнення. Узгодження ТЗ",    "Descript": "Техническое задание"},
         {"id":8,  "Num":"5.",   "Stage":"Конструювання",               "Descript": "Констрування"},
         {"id":9,  "Num":"6.",   "Stage":"Впровадження",                "Descript": "Впровадження"},
         {"id":10, "Num":"7.",   "Stage":"Супровід",                    "Descript": "Супровід", "Tag":["Отделпідтримки","Відділрозробки","Консультація"]}
       ]`

  erj:=json.Unmarshal([]byte(St), &Dt)
  if erj!=nil{
     Ltf("Error decode JSON format.") 
  }
    for i, t:=range(Dt){
       l,ok:=t["Tag"]
         
        if ok {
           fmt.Println(i,t["id"], t["Stage"],t["Descript"], "Підтримка :", l.([]interface{})[1] )
        }
           fmt.Println(t["id"], t["Stage"],t["Descript"],l )
  }
}
