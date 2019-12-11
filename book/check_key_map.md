Как проверить, содержит ли карта ключ в Go?
Операторы if в Go могут включать как условие, так и оператор инициализации.
Первая инициализация двух переменных - «value», которая получит либо значение «china» с карты, и «ok» получит значение bool, которое будет установлено в true, если «china» действительно присутствовал на карте.

Во-вторых, оценивается нормально, что будет правдой, если «фарфор» был на карте. На этой основе выведите результат.

Используя len, вы можете проверить, что карта погоды пуста или нет.

Нужно работать над своим кодом с различными размещенными приложениями из любого места на нескольких устройствах? Пройдите бесплатную пробную версию у поставщика Desktop-as-a-Service (DaaS), такого как CloudDesktopOnline . Для размещения на хостинге SharePoint и Exchange посетите сайт Apps4Rent.com сегодня.

```golang
package main
 
import "fmt"
 
func main() {
     
    fmt.Println("\n##############################\n")
    strDict := map[string]int {"japan" : 1, "china" : 2, "canada" : 3}
    value, ok := strDict["china"]
    if ok {
            fmt.Println("Key found value is: ", value)
    } else {
            fmt.Println("Key not found")
    }
     
    fmt.Println("\n##############################\n")
    if value, exist := strDict["china"]; exist {
        fmt.Println("Key found value is: ", value)
    } else {
        fmt.Println("Key not found")
    }
     
    fmt.Println("\n##############################\n")
    intMap := map[int]string{
        0: "zero",
        1: "one",
    }
    fmt.Printf("Key 0 exists: %t\nKey 1 exists: %t\nKey 2 exists: %t",
    intMap[0] != "", intMap[1] != "", intMap[2] != "")
     
    fmt.Println("\n##############################\n")   
    t := map[int]string{}
    if len(t) == 0 {
        fmt.Println("\nEmpty Map")
    }
    if len(intMap) == 0 {
        fmt.Println("\nEmpty Map")
    }else{
        fmt.Println("\nNot Empty Map")
    }
     
}
```
