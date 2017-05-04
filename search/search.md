### Поиск в массиве

Иногда необходимо найти в массиве єлементы по разным условиям.    
Для этого приведена ниже программа.


```golang
package main

import (
	"fmt"
	"strings"
)

// Функция проверки на разные условия для каждого из элементов в массиве
func filter(data []string, callback func(string) bool) []string {
     var result []string

     for _, each := range data {
         if filtered := callback(each); filtered {
            result = append(result, each)
         }
     }
      return result
}


// 
func main() {

	var data = []string{"wick", "jason", "ethan", "Город", "Гора", "Горячий","Ctра"}

	// Проверка на наличие буквы "е"
	var dataContainsO = filter(data, func(each string) bool {return strings.Contains(each, "e")})
	// filter количество букв в каждом слове
	var dataLenght5 = filter(data, func(each string) bool {return len(each) == 5})
	// filter количество букв в каждом слове
	var dataLenghs = filter(data, func(each string) bool {  return strings.HasPrefix(each, "Гор") })
	var dataLenghr = filter(data, func(each string) bool {  return strings.HasSuffix(each, "ра") })

	fmt.Println("data asli \t\t:", data)
	fmt.Println("filter ada huruf \"o\"\t:",        dataContainsO)
	fmt.Println("filter jumlah huruf \"5\"\t:",     dataLenght5)
	fmt.Println("filter jumlah huruf \"город\"\t:", dataLenghs)
	fmt.Println("filter jumlah huruf \"город\"\t:", dataLenghr)
}
```
