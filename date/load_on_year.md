package main

import (
	"time"
    "fmt"
    "net/http"
     "github.com/jinzhu/gorm"
)



// *********************************************
// Наполнение датами на год вперед календаря
// 
// /api/calendar/
// *********************************************
func Calendar_Of_Year(w http.ResponseWriter, r *http.Request){
     
     StartDate := time.Now()
     layout    := "2006-01-02"                   // Формат даты
     str       := "2019-01-01"                   // Дата старта 
     t,_       := time.Parse(layout , str)
     StartDate  = t

     // Цикл 365
     for i:=0;i<=365;i++{
     	 tm  := StartDate.AddDate(0,0,i).Format("02-01-2006")
     	 tms := StartDate.AddDate(0,0,i).Format("02-January-2006")
     	
     	 // Наполнение базы данных днями с датами на весь год начиная с 01-01-2019
         Add_days(tm,tms)     	 

         // Вывод на  
         fmt.Println(i,tm)
     }
}   
