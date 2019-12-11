## Этот пример программы демонстрирует, как декодировать строку JSON.
Язык Go обладает большей гибкостью для работы с документом JSON. В приведенном ниже примере вы можете декодировать или распаковывать документ JSON в переменную карты. Функция Unmarshal пакета json используется для декодирования значений JSON в значения Go.

```golang
// This sample program demonstrates how to decode a JSON string.
 
package main
 
import (
    "fmt"  
    "encoding/json" // Encoding and Decoding Package
)
 
// JSON Contains a sample String to unmarshal.
 
var JSON = `{
    "name":"Mark Taylor",
    "jobtitle":"Software Developer",
    "phone":{
        "home":"123-466-799",
        "office":"564-987-654"
    },
    "email":"markt@gmail.com"
}`
 
func main() {
    // Unmarshal the JSON string into info map variable.
    var info map[string]interface{}
    json.Unmarshal([]byte(JSON),&info)
 
    // Print the output from info map.
    fmt.Println(info["name"])
    fmt.Println(info["jobtitle"])
    fmt.Println(info["email"])
    fmt.Println(info["phone"].(map[string]interface{})["home"]) 
    fmt.Println(info["phone"].(map[string]interface{})["office"])
}
```


```
// GO language program with an example of Hash Table

package main

import (
"fmt"
)

func main() {
    var country map[int]string
    country = make(map[int] string)
    country[1]="India"
    country[2]="China"
    country[3]="Pakistan"
    country[4]="Germany"
    country[5]="Australia"
    country[6]="Indonesia"
    for i, j := range country {
        fmt.Printf("Key: %d Value: %s\n", i, j)
    }
}
```
