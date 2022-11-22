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





func Add_days(Dt,Dts string){
	var Cl Calendar
	 db, _ := gorm.Open("sqlite3", "/home/airpc/WORK/RESUME/resume.db")
	 defer db.Close()
	 Cl.Date    = Dt 
	 Cl.Datestr = Dts 
	 db.Create(&Cl)

}



var examples = []string{
    "May 8, 2009 5:57:51 PM",
    "Mon Jan  2 15:04:05 2006",
    "Mon Jan  2 15:04:05 MST 2006",
    "Mon Jan 02 15:04:05 -0700 2006",
    "Monday, 02-Jan-06 15:04:05 MST",
    "Mon, 02 Jan 2006 15:04:05 MST",
    "Tue, 11 Jul 2017 16:28:13 +0200 (CEST)",
    "Mon, 02 Jan 2006 15:04:05 -0700",
    "Thu, 4 Jan 2018 17:53:36 +0000",
    "Mon Aug 10 15:44:11 UTC+0100 2015",
    "Fri Jul 03 2015 18:04:07 GMT+0100 (GMT Daylight Time)",
    "12 Feb 2006, 19:17",
    "12 Feb 2006 19:17",
    "03 February 2013",
    "2013-Feb-03",
    //   mm/dd/yy
    "3/31/2014",
    "03/31/2014",
    "08/21/71",
    "8/1/71",
    "4/8/2014 22:05",
    "04/08/2014 22:05",
    "4/8/14 22:05",
    "04/2/2014 03:00:51",
    "8/8/1965 12:00:00 AM",
    "8/8/1965 01:00:01 PM",
    "8/8/1965 01:00 PM",
    "8/8/1965 1:00 PM",
    "8/8/1965 12:00 AM",
    "4/02/2014 03:00:51",
    "03/19/2012 10:11:59",
    "03/19/2012 10:11:59.3186369",
    // yyyy/mm/dd
    "2014/3/31",
    "2014/03/31",
    "2014/4/8 22:05",
    "2014/04/08 22:05",
    "2014/04/2 03:00:51",
    "2014/4/02 03:00:51",
    "2012/03/19 10:11:59",
    "2012/03/19 10:11:59.3186369",
    // Chinese
    "2014年04月08日",
    //   yyyy-mm-ddThh
    "2006-01-02T15:04:05+0000",
    "2009-08-12T22:15:09-07:00",
    "2009-08-12T22:15:09",
    "2009-08-12T22:15:09Z",
    //   yyyy-mm-dd hh:mm:ss
    "2014-04-26 17:24:37.3186369",
    "2012-08-03 18:31:59.257000000",
    "2014-04-26 17:24:37.123",
    "2013-04-01 22:43",
    "2013-04-01 22:43:22",
    "2014-12-16 06:20:00 UTC",
    "2014-12-16 06:20:00 GMT",
    "2014-04-26 05:24:37 PM",
    "2014-04-26 13:13:43 +0800",
    "2014-04-26 13:13:44 +09:00",
    "2012-08-03 18:31:59.257000000 +0000 UTC",
    "2015-09-30 18:48:56.35272715 +0000 UTC",
    "2015-02-18 00:12:00 +0000 GMT",
    "2015-02-18 00:12:00 +0000 UTC",
    "2017-07-19 03:21:51+00:00",
    "2014-04-26",
    "2014-04",
    "2014",
    "2014-05-11 08:20:13,787",
    // mm.dd.yy
    "3.31.2014",
    "03.31.2014",
    "08.21.71",
    //  yyyymmdd and similar
    "20140601",
    // unix seconds, ms
    "1332151919",
    "1384216367189",
}

var (
    timezone = ""
)

// func Date_samples() {
//     flag.StringVar(&timezone, "timezone", "UTC", "Timezone aka `America/Los_Angeles` formatted time-zone")
//     flag.Parse()

//     if timezone != "" {
//         // NOTE:  This is very, very important to understand
//         // time-parsing in go
//         loc, err := time.LoadLocation(timezone)
//         if err != nil {
//             panic(err.Error())
//         }
//         time.Local = loc
//     }

//     table := termtables.CreateTable()

//     table.AddHeaders("Input", "Parsed, and Output as %v")
//     for _, dateExample := range examples {
//         t, err := dateparse.ParseLocal(dateExample)
//         if err != nil {
//             panic(err.Error())
//         }
//         table.AddRow(dateExample, fmt.Sprintf("%v", t))
//     }
//     fmt.Println(table.Render())
// }
