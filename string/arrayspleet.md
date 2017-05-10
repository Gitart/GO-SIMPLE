## Пример работы со строкой

```golang
package main

import (
	"fmt"
)

func main() {

}


// Incorrect
func Incorrect(){
  t:=[]string{"Первый","Второй", "Третий"}
  z:=[]byte("Первоочередной")
  l:="Второстепенный"


  fmt.Println("Ответ:", string(l[1]))
  fmt.Println(t[1],string(z[4]))
}

// Correct
func Nstr(){

  x:="Привет"
  z:=strings.Split(x,"")[2]
  fmt.Println(z)

  s:=makeID("Dsdf")
  fmt.Println(s)
}


func Filter(limit int, predicate func(int) bool, appender func(int)) {
for i := 0; i < limit; i++ {
if predicate(i) {
 appender(i)
 }
 }
}
```
