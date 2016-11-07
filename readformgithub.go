/ 
// For read settings file 
// Для удаленного управления программной средой
// для установки нужных настроеек
// для блокировки процессов
// для централизованного управления и т.д.
// для закачки справочников 
// Для вычитки из далека настроеек
// 



/*
{ 
  "Id":"S-0001",
  "Name":"Setting for services",
  "Other":"Description for settings and view",
  "Date":"27-10-2016T12:00:02",
  "DateDue":"27-10-2020T12:01:02",
  "Flag":false
  }
*/


package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
)

type Mst map[string]interface{}

func main() {
	
	url:= "https://raw.githubusercontent.com/Gitart/work/master/setting.json"
	resp, err := http.Get(url)
	
	if err != nil {
		panic(err)
	}
	
	defer resp.Body.Close()
	
	// reads document of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// show the HTML code as a string %s
	// fmt.Printf("%s\n", html)

    var m Mst    
	errj := json.Unmarshal([]byte(html), &m)

   if errj != nil {
		panic(err)
	}


    // Обязательное обработка 
    // Или проверять через val,ок:=m["Name"]
    // или через зараннее опредленную структуру

    t:=m["Name"].(string)
    t1:=m["Other"].(string)
    t2:=m["Date"].(string)
    t3:=m["Flag"].(bool)

    if t3{
	   fmt.Println("Yes")
    }else{
       fmt.Println("No")
    }

	// show the HTML code as a string %s
	fmt.Printf("%s\n", t)
	fmt.Printf("%s\n", t1)
	fmt.Printf("%s\n", t2)
}
