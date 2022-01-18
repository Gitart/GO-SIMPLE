package main

import (
       "fmt"	
	"time"
	"github.com/fatih/color"
)

// Connect and view
func FmtGreeen(text string) {
	Appcolor := color.New(color.FgGreen, color.Bold)
       Appcolor.Println(text)
}

func FmtRed(text string) {
     Appcolor := color.New(color.FgRed, color.Bold)
     Appcolor.Println(text)
}

func FmtBlue(text string) {
     Appcolor := color.New(color.FgBlue, color.Bold)
     Appcolor.Println(text)
}

func FmtApp(app, text string) {
     Appcolor := color.New(color.FgRed, color.Bold)
     Appcolor.Print(" " + app + ": ")
     Appcolor = color.New(color.FgWhite )
     Appcolor.Println(text)    
}

func FmtSys(err string){
     red := color.New(color.FgRed, color.Bold).PrintfFunc()
     red("Warning")
     red("Error: %s", err)
}


 func InfoSys(text, clr string ){
      cl:=map[string] int {
            "white"   : 0,
            "red"     : 31,
            "green"   : 32,
            "yellow"  : 33,
            "blue"    : 34,
      }  

      fmt.Println(cl[clr])

       atrb:=color.Attribute(cl[clr])
      red := color.New(atrb, color.Bold).PrintfFunc()
      red("-------------------Warning")
      red("NOTE : %s",text )
 }



func main(){

        InfoSys("Красный",      "white")
        InfoSys("Красный",      "red")
        InfoSys("Голубой цвет", "blue")
        InfoSys("Зеленый цвет", "green")
        InfoSys("Зеленый цвет", "yellow")


      fmt.Printf("%+v", color.FgRed) 
      fmt.Printf("%+v", color.FgGreen) 

       FmtSys("Start ghjwtlehss")
       FmtApp("APP","📔 Start......")
       time.Sleep(time.Second * 2)
       FmtApp("APP","FINISH BASE PROCEDURE")

       FmtGreeen(" APP: 📔 Start......")
       time.Sleep(time.Second * 2)
       FmtGreeen(" APP: Начало выполнения загрузки файлов")
       time.Sleep(time.Second * 1)
       FmtGreeen(" Все файлы были успешно загружены")
       time.Sleep(time.Second * 1)
       FmtRed(" AAPP: Все файлы были успешно загружены")
       time.Sleep(time.Second * 1)
       FmtRed(" APP: Все таблицы были созданы")
       time.Sleep(time.Second * 1)
       FmtBlue(" SYS: Начало создания индексов")
       time.Sleep(time.Second * 2)
       FmtGreeen(" IDX:  Индексы были успешно созданы" )
       time.Sleep(time.Second * 2)
       FmtGreeen(" APP: ИНДКЕКСЫ БЫЛИ УСПЕШНО СОЗДАНЫ" )
       time.Sleep(time.Second * 2)
       FmtBlue("⏰ ПРИМЕР ИСПОЛЬЗОВАНИЯ ПРОСТОГО ПРИЛОЖЕНИЯ ДЛЯ СОЗДАНИЯ ПРОЦЕДУРЫ")
       time.Sleep(time.Second * 2)
       FmtRed("🧑‍⚖️ ПРИМЕР ИСПОЛЬЗОВАНИЯ ПРОСТОГО ПРИЛОЖЕНИЯ ДЛЯ СОЗДАНИЯ ПРОЦЕДУРЫ")
}

