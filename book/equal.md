## Как сравнить равенство структуры, среза и карты?
Функция DeepEqual из отражающего пакета, используемая для проверки x и y, «глубоко равна». Это применимо для: Значения массива глубоко равны, когда их соответствующие элементы глубоко равны. Значения структур сильно равны, если их соответствующие поля, как экспортированные, так и не экспортированные, глубоко совпадают. Ниже приведена короткая программа для сравнения структуры, среза или карты, равны или нет.

```golang
package main
 
import (
"fmt"
"reflect"
)
 
type testStruct struct {
    A int
    B string
    C []int
}
 
func main() {
    st1 := testStruct{A:100,B:"Australia",C:[]int{1,2,3}}   
    st2 := testStruct{A:100,B:"Australia",C:[]int{1,2,3}}
    fmt.Println("Struct equal: ", reflect.DeepEqual(st1, st2))
     
     
    slice1 := []int{1,2,3,4}
    slice2 := []int{1,2,3,4}
    fmt.Println("Slice equal: ", reflect.DeepEqual(slice1, slice2))
     
    map1 := map[string]int{ 
        "x":1,
        "y":2,
    }
    map2 := map[string]int{ 
        "x":1,
        "y":2,
        "z":3,
    }
    fmt.Println("Map equal: ",reflect.DeepEqual(map1, map2))
}
```
