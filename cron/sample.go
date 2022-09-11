// https://godoc.org/github.com/robfig/cron
// go get github.com/mileusna/crontab

// Работающий вариант
// https://github.com/carlescere/scheduler

// Опросник
// https://app.delighted.com/dashboard

// import "github.com/go-co-op/gocron"   // No work with AT!!!! - не работает в библиотеке установка времени ежедневно!!!!!!!
// https://golangexample.com/gocron-a-golang-job-scheduling-package/

package main
import "github.com/jasonlvhit/gocron"       // WORK!!! with AT
import "time"
import "fmt"
import "log"
import "runtime"
import "github.com/carlescere/scheduler"
import "net/http"
import "github.com/labstack/echo/v4"


// ******************************************
// Main process
// ******************************************
func main() {
    
    e := echo.New()
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })
    
    e.GET("/clear",  Clear)
    e.GET("/jobs",   Jobs)
    e.GET("/start",  Starts) 
    e.GET("/info",   Info) 


    // Запуск задач в фоне
    go crons()
    
    // Start serves
    e.Logger.Fatal(e.Start(":1323"))
}

// ******************************************
// Очистить удалить все запланированные задания
// ******************************************
func Clear (c echo.Context) error {

     fmt.Println("All task clear")
     gocron.Clear()
     return c.String(http.StatusOK, "Clear!")
}

// ******************************************
// Очистить удалить все запланированные задания
// ******************************************
func Info (c echo.Context) error {
     fmt.Println(gocron)
     return c.String(http.StatusOK, "Info")
}


// ******************************************
// Jobs возвращает список заданий из планировщика
// ******************************************
func Jobs (c echo.Context) error {
     fmt.Println(gocron.Jobs())
     return c.String(http.StatusOK, "All Jobs!")
}

// ******************************************
// Start tasks
// ******************************************
func Starts (c echo.Context) error {
     go crons()
     return c.String(http.StatusOK, "Start all tasks by sheduler.")
}


// **************************************************
// Shedulers cron
// **************************************************
func crons(){
    log.Println("Запуск всех кронов ....")

    // Запуск по времени
    sheduler:=[]string{"20:10:30","20:30:10","20:30:17"}

    for _, sr := range sheduler {
        log.Println("Set date :", sr)
        gocron.Every(1).Day().At(sr).Do(task16)
    }

    // Запуск каждые 10 секунд
    gocron.Every(10).Second().Do(task)
    gocron.Every(1).Minutes().Do(task)

    // gocron.Every(2).Second().Do(Task1sec)
    // gocron.Every(10).Seconds().Do(Task3sec)

    <- gocron.Start()
}

   
// *********************************************
// Otehr samples
// *********************************************
func ShedulerSample() {
    job := func() {
        t := time.Now()
        fmt.Println("Time's up! @", t.UTC())
    }

   jobs := func() {
        t := time.Now()
        fmt.Println("Time 16 47", t.UTC())
    }

    // Run every 2 seconds but not now.
    scheduler.Every(2).Seconds().NotImmediately().Run(job)
      
    // Run now and every X.
    scheduler.Every(5).Minutes().Run(job)
    scheduler.Every().Day().Run(job)
    scheduler.Every().Monday().At("08:30").Run(job)
    scheduler.Every().Day().At("17:48").Run(jobs)
      
    // Keep the program from not exiting.
    runtime.Goexit()
}



func crons1() {
    // customLocation, _ := time.LoadLocation("Europe/Kiev")
    
    // Инициализируем объект планировщика
    // s := gocron.NewScheduler(time.UTC)
    // s := gocron.NewScheduler(customLocation)
    // fmt.Println(s.Location())

    // //  Каждые 3 минуты
    // s.Cron("*/3 * * * *").Do(taskmint)

    // //добавляем одну задачу на каждую минуту
    // s.Cron("* * * * *").Do(task)

    // // Часы
    // s.Cron("* 3 * * *").Do(taskmint)
    // s.Every(5).Seconds().Do(task5sec)
    // s.Every(1).Day().At("10:30").Do(task)
    // s.Every(1).Day().At("18:28").Do(task)
    // s.Every(1).Day().At("10:30").Do(task16)
    // _, _ = s.Every(1).Days().At("18:19").Do(task)
    // _, _ = s.Cron("10 18 * * *").Do(task16)   // every day at 1 am
    // s.Every(1).Day().At("16:47").Do(task16)

    // запускаем планировщик 
    // с блокировкой текущего потока
    // s.StartBlocking()

    // s.StartAsync()
}



func Preview(txt string){
    log.Println(txt)
}

func task16() {
    Preview("1618 cek ")
}


func task5sec() {
    Preview("5 cek ")
}


func task() {
    Preview("Каждая минута !")
}

func taskmint() {
    Preview("Три минуты прошло !")
}


func Task1sec() {
    Preview("Каждая секунда !")
}

func Task3sec() {
    Preview("Каждая 3 секунда !")
}





// http://www.cronmaker.com/;jsessionid=node0mkh5arr6qlktpipgo241qv3f286021.node0?0
// 
// 

/*

*     *     *     *     *        

^     ^     ^     ^     ^
|     |     |     |     |
|     |     |     |     +----- day of week (0-6) (Sunday=0)
|     |     |     +------- month (1-12)
|     |     +--------- day of month (1-31)
|     +----------- hour (0-23)
+------------- min (0-59)

Examples

* * * * * run on every minute
10 * * * * run at 0:10, 1:10 etc
10 15 * * * run at 15:10 every day
* * 1 * * run on every minute on 1st day of month
0 0 1 1 * Happy new year schedule
0 0 * * 1 Run at midnight on every Monday

Lists

* 10,15,19 * * * run at 10:00, 15:00 and 19:00
1-15 * * * * run at 1, 2, 3...15 minute of each hour
0 0-5,10 * * * run on every hour from 0-5 and in 10 oclock


IF YOU NEED ANY HELP HAVE ANY PROBLEMS THAT CANNOT BE SOLVED BY READING THE FAQ

/*
