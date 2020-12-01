## Формирование объекта из массива

```json
 "statuses": {
        "1": {
            "id": 1,
            "code": "order",
            "title": "Заказан"
        },
        "10": {
            "id": 10,
            "code": "payment",
            "title": "Оплачено"
        },
        "11": {
            "id": 11,
            "code": "close",
            "title": "Закрыт"
        },
        "12": {
            "id": 12,
            "code": "reject",
            "title": "Отказ"
        },
        "13": {
            "id": 13,
            "code": "return",
            "title": "Возврат"
        },
        "14": {
            "id": 14,
            "code": "archive",
            "title": "Архив"
        },
        "2": {
            "id": 2,
            "code": "prepea",
            "title": "Подготовка"
        },
        "3": {
            "id": 3,
            "code": "waiting",
            "title": "Ожидание"
        },
        "4": {
            "id": 4,
            "code": "documents",
            "title": "Формирование"
        },
        "5": {
            "id": 5,
            "code": "package",
            "title": "Упаковка"
        },
        "6": {
            "id": 6,
            "code": "to post",
            "title": "Отправка"
        },
        "7": {
            "id": 7,
            "code": "sending",
            "title": "Отправлено"
        },
        "8": {
            "id": 8,
            "code": "arrival",
            "title": "Прибыло"
        },
        "9": {
            "id": 9,
            "code": "getting",
            "title": "Получено"
        }
    }
    ```

## Convert

```go
// Data 
func statusesRec() models.MapIntStatus {

    Dat:= make(models.MapIntStatus)
    // var St models.Status

	 var statuses []models.Status

    var db = mysql.Connect()
	defer db.Close()
    db.Model(&statuses).Scan(&statuses)     

    // ln:=len(statuses)

	// statusesRec := make([]models.StatusRec, 100)
	// StatusesRec := make([]models.StatusRec, ln)

    // for id, status:=range statuses {
    // 	 statusesRec[id].Id     = status.Id
    // 	 statusesRec[id].Status = status
    // }

    


    // for id, status := range statuses {
    // 	  StatusesRec[id].Id     = status.Id
    // 	  StatusesRec[id].Status = status
    // }


     for _, status := range statuses {
    	  
       //    St.Id    = status.Id
       //    St.Code  = status.Code
       //    St.Title = status.Title

    	  // Dat[status.Id] = St
    	  Dat[status.Id] = status
    }

   // fmt.Println(Dat)

    //fmt.Println(StatusesRec)
return Dat
    //return StatusesRec
}
```
