package main


// Пакеты
import (

	r "github.com/dancannon/gorethink"
	// "encoding/json"
	"fmt"
)



/***************************************************************************************************************************************
 *
 *   Title        : Connection to DB
 *   Initialisation Service
 * 	 Date         : 2015-03-11
 *	 Description  : Initialization DB Connect
 *   Author       : SAVCHENKO ARTHUR 
 *   ☎           : 8-097-5547468
 *
 ****************************************************************************************************************************************/
func Dbini() {
	// Инициализация подключения к базе
	// на той машине где расположен и стартует сервис
	// Для переключения на тестовую машину
    IpPort := "localhost:28015"                       // Локальный ресурс 
	session, err := r.Connect(r.ConnectOpts{Address: IpPort, Database: "test",  Username:"admin", Password: ""})
	// Обработка ошибок
	if err != nil {
	   fmt.Println("NO CONNECTION.")
	   return
	}

	// Максимальное количество подключений
	// По умолчанию 200
	session.SetMaxOpenConns(200)
	session.SetMaxIdleConns(200)
	sessionArray = append(sessionArray, session)
}


// Описание структуры
type Trr struct{
     Id    string
     Tags  []string
     Title string
}

// Инициализация
func main() {
    Dbini()
    start()
}


// *************************************************************************
// http://stackoverflow.com/questions/25025409/delete-element-in-a-slice
// Пример работы с массивом 
// Добавление
// Удаление
//  и запись в базу данных
// *************************************************************************
func start() {
	var T Trr
	var U USER

    go Dels()  // Удаление таблиц

    T.Id   = "ID-00" 
    T.Tags = []string{"Nom","Ters","News"}

    fmt.Println("Новоая ветка - первое значение",T.Tags[0])
    Inst(T)

    T.Tags[0] = "Новое значение 0"
    T.Tags[1] = "Новое значение 2"
    T.Tags[2] = "Новое значение 3"

    // y:=[]string{"ddde",",fkrjy","Балкон"}
    l:=make([]string,3)
    l=append(l, "fff")

    fmt.Println(l)

    // Добавление нового элемента в структуру
    T.Tags =append(T.Tags, "eerrrrrrrrrrr", "Второй элемент")
    
    // T.Tags[3] = "Новое значение 4"
    // T.Tags[4] = "Новое значение 5"

    fmt.Println("Новоая ветка - первое значение")
    Inst(T)


    U.Password = "Pass-00"
    Inst(U)

    U.Password = "Pass-030"
    U.Os       = "Os WIndows"
    Inst(U)

    U.Password="Newpass"
    Inst(U)    

    
    Dels()  // Удаление таблиц
    
    // Добавление несколько элементов в маасив
    T.Tags =append(T.Tags, "Новая секция", "Второй секция")
    Inst(T)    

    // T.Tags =delete(T.Tags, "Новая секция", "Второй секция")

    RemoveIndex(T.Tags,2)

    T.Title="Без второго элемента"
    Inst(T)    

    fmt.Println(T)	
}




// Универсальная вствака
func Inst(Ms interface{}){
 	 r.DB("test").Table("wrk").Insert(Ms).Run(sessionArray[0])
}

// Удаление записей из таблицы
func Dels(){
 	 r.DB("test").Table("wrk").Delete().Run(sessionArray[0])
}



// Удаление элемента из массива
func RemoveIndex(s []string, index int) []string {
    return append(s[:index], s[index+1:]...)
}




// Пример удаление элемента из массива
// http://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-array-in-golang
// http://stackoverflow.com/questions/28699485/remove-elements-in-slice
// func RemoveIndex(s []int, index int) []int {
//     return append(s[:index], s[index+1:]...)
// }

// func main() {
//     all := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
//     fmt.Println(all) //[0 1 2 3 4 5 6 7 8 9]
//     n := RemoveIndex(all, 5)
//     fmt.Println(n) //[0 1 2 3 4 6 7 8 9]
// }
