# Поля разбивают срез s вокруг каждого экземпляра одного или нескольких последовательных пробельных символов, возвращая срез субликсов s или пустой список, если s содержит только пробел.
FieldsFunc интерпретирует s как последовательность кодированных точек Unicode в кодировке UTF-8. Он разбивает срез s при каждом запуске кодовых точек c, удовлетворяющих f (c), и возвращает срез субликсов s.

```golang
package main
 
import (
    "bytes"
    "fmt"
    "strings"  
)
 
func main() {
    fmt.Println("############ Fields #####################\n")
    listCountry := []byte(" Australia   Canada Japan Germany   India")
    fmt.Printf("%q",listCountry)
     
    country := bytes.Fields(listCountry)    
    for index,element := range country{
        fmt.Printf("\n%d => %q", index, element)
    }   
     
    fmt.Println("\n############ FieldsFunc #####################\n")
    sentence := []byte("The Go language has built-in facilities, as well as library support, for writing concurrent programs.")
    fmt.Printf("%q",sentence)
    vowelsSpace := "aeiouy "
    chop := bytes.FieldsFunc(sentence, func(r rune) bool {
        return strings.ContainsRune(vowelsSpace, r)
    })
    for index,element := range chop{
        fmt.Printf("\n%d => %q", index, element)
    }   
}
```
