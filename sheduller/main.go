// Пеердача функции в качестве параметра
// Шедулер по времени с вызовом процедур

package main

import "fmt"
import "time"


type fn func(string) 

// Vfin procedure
func main(){
{
	var d="Start one"
     fmt.Println(d) 
}

{
	var d = "Start two"
    fmt.Println(d) 
}


   go DurationProc(10,  f1, "10 сек")        // Третий
   go DurationProc(5,   f1, "5 сек")         // Втрой  
      DurationProc(3,   f2, "3 сек")         // Первый 
      fmt.Println("ok")
}

func f1(s string) {
	  fmt.Printf("------ %s \n", s)
}

func f2(s string) {
     fmt.Printf("========== %s \n", s)
}

//Задержка
func DurationProc(dur time.Duration, fnc fn, t string){
	for {
        time.Sleep(time.Second * dur)
	    fnc(t)
	    fmt.Print(dur)
	    fmt.Print(" Worked")

     }
}

