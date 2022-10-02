package main

import "fmt"
import "os"
import "time"
import "flag"

// ************************************************
// Main func
// ************************************************
func main(){
     filename := flag.String("fl", "utl", "Name output go file")
     funcname := flag.String("fn", "fn",  "Name function")
	 flag.Parse()
     
     FileGenerator(*filename, *funcname)

 	 fmt.Printf("GENERATOR: FUNCTION: %s FILE: %s \n", *funcname, *filename)
}

// ************************************************
// Creator file
// ************************************************
func FileGenerator(namefile, namefunc string){
  	 fl,_:=os.Create(namefile+".go")
	 defer fl.Close()
 	 fl.Write(Assembly(namefile, namefunc))
}

// ***********************************
// Assembly function body
// ***********************************
func Assembly(fl, fn string) []byte {

	cd:=map[string]string {
		           "pr1": process1,
		           "pr2": process2,
		           "pr3": process3,
		           "filter:pr": process3,
		           "ut": utl,
		          } 

     d:=time.Now().Format("02/01/2006 15:04:05")
     h:=Hdl(fl,fn,d)
    return []byte(h + cd[fn] )
}


func Hdl(fl,fn,dt string) string{
	return fmt.Sprintf(Hendl,fl,fn,dt)
}



var (
Hendl =`
// ***********************************************************
// WARNNIG !
// DATE GENERATE  : %s
// FUNCTIONS NAME : %s
// FILE NAME      : %s
// THIS CODE GENERATE AUTOMATICALLY GENERATOR PROGRAMM
// PLEASE DO NOT TOUCH THIS WITH YOUR HANDS
// ***********************************************************
`

process1 = `

package main
import "fmt"

// *************************************
// Функция для создания компании
// *************************************
func Func01_Ut01(inp string){
	fmt.Println(inp)
}

// *************************************
// Функция для создания имени
// *************************************
func Func01_Ut03(){
	fmt.Println("Process1 Func2")
}

`	
process2 = `
package main
import "fmt"

// *************************************
// Функция для создания имени
// *************************************
func Ut01(){
	fmt.Println("Process 2 Func1")
}

// *************************************
// Функция для создания имени
// *************************************
func Ut03(){
	fmt.Println("Process2 Func 2")
}`	

process3 = `
package main
import "fmt"

// *************************************
// Функция для создания имени
// *************************************
func Process03_CreatedNewCompany(){
	fmt.Println("Company is created")
}

// *************************************
// Функция для создания имени
// *************************************
func Process03_Ut03(){
	fmt.Println("Name is created")
}`	


utl = `
package main
import "fmt"
// *************************************
// Функция для создания имени
// *************************************
func Utl_Process03_Ut03(){
	fmt.Println("Name is created")
}`	

)

func DerefString(s *string) string {
    if s != nil {
        return *s
    }
    return ""
}
