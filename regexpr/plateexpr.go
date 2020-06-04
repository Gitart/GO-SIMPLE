package main

import (
    "regexp"
    "fmt"
)


func main() {
     s := StrPlateNumber("AAAddaa 1s 23 47 ffdgfg ggf")
     d := StrPlateNumber("ffff1234fasddddddddddddd")
     
     if s==d{
        fmt.Println("YES ! Совпадают цифры в номерах.")	
     }else{
     	fmt.Println("NO  ! ")	
     }
}

// Возвращает цифры из выражения типа "zzzz123354yyyy"
func StrPlateNumber(str string) string {
    ret:=""
    var re = regexp.MustCompile(`\d`)
    re.FindAllString(str, -1)
   
    for _, match := range re.FindAllString(str, -1) {
       ret=ret+match
    }
    return ret
}
