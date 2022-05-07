// Генерация файлов на основе структур данных

package main

import "os"
import "log"
import "fmt"
import "reflect"


// Vfin procedure
func main(){
     StructuresView()
}

// Прокурутка структуры
func StructuresView(){
     filename := "work.txt"
     book := Companies{}
	 e    := reflect.ValueOf(&book).Elem()

	for i := 0; i < e.NumField(); i++ {
		varName  := e.Type().Field(i).Name
		varType  := e.Type().Field(i).Type
		varValue := e.Field(i).Interface()
		fmt.Printf("%v %v %v\n", varName, varType, varValue)
		sfd := fmt.Sprintf(ajxfield, varName)
		Wrfile(filename, sfd)
	}
}

// *********************************************
// main process ....
// *********************************************
func mainі(){
	filename   := "work.txt"
	filenameJs := "work.js"
    sajax      := fmt.Sprintf(ajx, "POST", "/doc/add")

    Wrfile(filenameJs, sajax + "\n")

    for _, rt:= range readtxt {
    	sfd := fmt.Sprintf(ajxfield, rt)
    	Wrfile(filename, sfd)
    }
    
    log.Println("Genertation successful...")
}

// ******************************************************
func Wrfile(filename, text string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
    
    if err != nil {
       panic(err)
    }

    defer f.Close()

    if _, err = f.WriteString(text); err != nil {
       panic(err)
    } 
}
