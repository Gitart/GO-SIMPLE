## Пример работы с Pointer.
[This sample](https://play.golang.org/p/ETouyXB_nZ)   
[Pointer method](https://tour.golang.org/methods/5)   
[Point sample](https://www.golang-book.com/books/intro/8)   
[Point sample](https://nathanleclaire.com/blog/2014/08/09/dont-get-bitten-by-pointer-vs-non-pointer-method-receivers-in-golang/)   




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

Полезно когда необходимо, что-бы после вызова метода
начальное значение ситруктуры было измененно то-же.
Вторя функция использует уже результат вызова первого метода.

