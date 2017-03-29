## Client get

Получение информации с сервиса в виде JSON файла

```golang
package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "http://kparser.pp.ua/json/film/7988?="
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)


	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)


    // Чтение заголовка 
	// fmt.Println(res)
	
    // Чтение тела ответа
	fmt.Println(string(body))
}
```
