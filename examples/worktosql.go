

// Useful link for ODBC connect to Database
// https://github.com/alexbrainman/odbc - драйвер для ODBC
// http://www.easysoft.com/support/kb/kb01045.html
// https://github.com/golang/go/wiki/SQLDrivers - примеры драйверов подключения к разным источникам

package main

import (
 // _ "code.google.com/p/odbc"
    _ "github.com/alexbrainman/odbc"
      "database/sql"
      "fmt"
      "os"
    //"github.com/astaxie/beedb"
    //"github.com/weigj/go-odbc"
)



// Gloval variable
var (
      Server = "driver={sql server};server=170.20.60.20;uid=k;pwd=Q123;database=DB"
)
// var db *odbc.Connect

// *************************************************************************************
// Онсновная процедура
// *************************************************************************************
func main() {

    // pid:=os.Getppid()
    Nametable:=os.Args[1]  // Name table
    NameField:=os.Args[2]  // Name fields - имена двух полей через запятую 
    Valuestr :=os.Args[3]  // Value 1
    Descrpit :=os.Args[4]  // Value 2

    fmt.Println(os.Args)
    
    // Добавление в таблицу
    AddToTable(Nametable,NameField,Valuestr,Descrpit)

    // Addtodatabase(pif)
    ExecXp("Adding",Valuestr)
}


// **************************************************************************
// Init
// **************************************************************************
func init(){
    // fmt.Println("Init start")
}

// func init(){
//  db, err: = sql.Open("odbc", "driver={sql server};server=172.25.65.2;uid=Savchenko;pwd=Qwerty123;database=TestDB")

//     if err != nil {
//         fmt.Println(err)
//         return
//     }
// }


// *************************************************************************************
// Вставка одного поля 
// в одну таблицу
// *************************************************************************************
func AddToTable(Table, Field, Values, Descript string){

    // db, err := sql.Open("odbc", "driver={sql server};server=.;uid=sa;pwd=Qwerty123;database=TestDB")
    // Очистка таблиц
    // Deltab(Table)

    db, err := sql.Open("odbc", Server)

    if err != nil {
        fmt.Println(err)
        return
    }

    if err != nil {
       fmt.Println(err)
       return
    }
   
    stms,_ := db.Prepare("INSERT INTO "+Table+"("+Field+") VALUES(?,?)")
    _, Err := stms.Exec(Values, Descript) 

    if Err != nil {
        fmt.Println("Ошибка при вставке", Err)
        return
    }

    defer db.Close()
}



// **************************************************************************
// Добавление в таблицу № 1 вариант
// **************************************************************************
func Addtodatabase(Txt string){

    // db, err := sql.Open("odbc", "driver={sql server};server=.;uid=sa;pwd=Qwerty123;database=TestDB")
    db, err := sql.Open("odbc", Server)

    if err != nil {
       fmt.Println(err)
       return
    }

    /*stms, err := db.Prepare("INSERT INTO dbo.Test (playname, MATCH ) VALUES ( ?, ? )")
    _, err = stms.Exec("zk", 23)
    CheckError(err)*/
    /*stmt, err := db.Prepare("delete from Test where playname=?")
    CheckError(err)

    res, err := stmt.Exec("zk")
    CheckError(err)

    affect, err := res.RowsAffected()
    CheckError(err)

    fmt.Println(affect)*/

    /*stms, err := db.Prepare("update Test set Match=? where playname=?")
    _, err = stms.Exec(12, "p")
    CheckError(err)*/
    
    // rows, err := db.Query("SELECT MATCH FROM Test where playname=?", "p")
    stms,_ := db.Prepare("INSERT INTO Plans(Title, Startdate, Montsname, [User], Description) VALUES(?,?,?,?,?)")
    _, Err := stms.Exec(Txt,"2016-02-03","Jaunary","adm","New") // Количество столько же сколько и переменніх

    if Err != nil {
       fmt.Println("Error ",Err)
       return
    }

    defer db.Close()

    // var match int
    // for rows.Next() {
    //     rows.Scan(&match)
    //     fmt.Println(match)
    // }
}


// **************************************************************************
// Удаление таблицы
// **************************************************************************
func Deltab(Ntab string){
     db, err := sql.Open("odbc", Server)
     if err != nil {
        fmt.Println(err)
        return
     }
    
    stmt, err := db.Prepare("DELETE FROM " + Ntab  )
    _, Err    := stmt.Exec() 

     if Err != nil {
        fmt.Println("Error",Err)
        return
    }
    defer db.Close()
}


// **************************************************************************
//  Хранимая процедура
//  Name - имя процедуры
//  Param - параметр
// **************************************************************************
func ExecXp(Name, Param string){
     db, err := sql.Open("odbc", Server)
     if err != nil {
        return
     }
     defer db.Close()

    // Варианты использования выполнения хранимой процедуры
    // db.Exec("Adding " + Param )
    // db.Exec("Adding ?", Param )
     db.Exec("? ?", Name, Param)
}

// error check
func CheckError(err error) {
    if err != nil {
        panic(err)
    }
}

// Read Json file
func ReadJson(){
   

}

