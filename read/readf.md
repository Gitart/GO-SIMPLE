## Read file

```golang
package main

import(
  "io/ioutil"
 "strings"
 "fmt"
)

func main(){  
   data,_ :=ioutil.ReadFile("readlines.go")
   lines  :=strings.Split(string(data),"\n")

   for i,l:= range lines{   
       fmt.Printf("%d %s",i,l)
       fmt.Println(lines[i])

   }

   fmt.Println(string(data))

}
```

### Парсинг строки с разделителями
```golang
package main

import (
  "fmt"
   "regexp"

)
func main(){

  r,_:=regexp.Compile("p|oo|d")

 matches:=r.FindAllString("oopsdddd", -1)

 // %q quote element to see... instead of %s
 fmt.Printf("%q \n", matches)

 for _, m:=range matches {
        fmt.Println(m)
 }

}
```
