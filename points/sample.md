## Пример работы с Pointer.
[Pointer method](https://tour.golang.org/methods/5)

Показывает как меняется переменная (метод I) в структуре
если использовать для передачи данных через & адресс 
а в функции принимать вызов  через *Vertex. Будет измененно 
начальное значение в структуре и значение будет равно 18.
Если не использовать адрес & и ссылку на значение *, 
значение будет равно 17.


```golang
package main

import (
	"fmt"
	
)

type Vertex struct {
     I  int
}

func Abs(v Vertex) int {
	return v.I+12
}

func Scale(v *Vertex) {
       v.I=v.I+1
}

func main() {
	v := Vertex{5}
	Scale(&v)
	fmt.Println(Abs(v))
}
```

