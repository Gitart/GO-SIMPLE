## Work with structure and slice



```golang
// 
// Title   : Status production for customer
// Created : 23/02/2018 11:00
// Change  : 23/02/2018 13:19
// Author  : Savchenko Arthur
// Usage   : n:=Rls_status(2)["Id"] n:=Rls_status(2); -> n[id], n["Title"]
// Company : 
// 
func Rls_status(Num int) Mst  {
	var data []Mst
	n:=Num-1
	
	d:=`[
		      {"Id":1, "Title":"Установлено",         "Note":"Установлен в банке протестированный релиз",                            "St":["a","b","c"]},
		      {"Id":2, "Title":"Не установлено",      "Note":"Любая причина не установки   у разработчика..",                                  "St":["d","e","f"]},
		      {"Id":3, "Title":"Тестирование разработчиком",  "Note":"Ожидает установку  у разработчика релиз. Тестируется у разработчика.",         "St":["g","h","i"]},
		      {"Id":4, "Title":"Успешный тест разработчиком", "Note":"Ожидает установку  у разработчика релиз. Разработчик протестировал успешно.",  "St":["k","l","m"]},
		      {"Id":5, "Title":"Тестирование" ,        "Note":"Тестируется заказчиком и ожидается решение от заказчика.",             "St":["o","p","r"]},
		      {"Id":6, "Title":"Успешный тест " ,       "Note":"Тест заказчиком пройден успешно. Подготовка к установке ан продуктивный сервер","St":["s","t","f"]},
		      {"Id":7, "Title":"Установлено  на прод", "Note":"Релиз успешно установлен у заказчика.",                                "St":["h","c","m"]},
		      {"Id":8, "Title":"Отменен",                     "Note":"Релиз отменен заказчиком.",                                            "St":["x","y","z"]},
		      {"Id":9, "Title":"Плановый",                    "Note":"Релиз плановый. Сроки ориентировочные и могут меняться.",              "St":["i","o","p"]}
	    ]`

    er := json.Unmarshal([]byte(d), &data)
	Err(er, "Error unmarshaling.")

    // Test
     fmt.Println(data[1]["Title"])
    // fmt.Println(data[1])


			/*
			    for _, t:=range(data) {
			    	// fmt.Println(t["Title"])

			    	if t["Id"]==n{
			    		fmt.Println(t["Title"])
			    		break
			    	}

			    }
			*/

    return data[n]
}

```


## Usage
```golang


func Test_status(w http.ResponseWriter, req *http.Request){
	d:=Rls_status(9)
	fmt.Println(d["Id"], d["Title"], d["St"].([]interface{})[1:])
    
    s:=Rls_status(6)["Title"]
    fmt.Println(s)
}
```
