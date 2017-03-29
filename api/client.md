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

	url := "http://kparser.pp.ua/json/film/7988?Name=Films&ID=12435"
	req, _ := http.NewRequest("POST", url, nil)
        
	// Header set
	req.Header.Add("key", "key-003939")
	req.Header.Add("user", "User")
	req.Header.Add("authorization", "Basic R29uczoxMjM0")

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
```
