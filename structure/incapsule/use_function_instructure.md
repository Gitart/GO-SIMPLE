## Использование функции в структуре
Удобно использовать т.к. можно менять тело функции кроме входных параметров
https://www.geeksforgeeks.org/function-as-a-field-in-golang-structure/

Например есть структура с функцией 
```go
type Companies struct {
  Id            int64  `json:"id"`            // Id records 
  Title         string `json:"title"`         // Наименвоание компании
  Pending       func(int, int) int
}
```

Ее использование
```go

func testFunctionInstructure(){
     book     := Companies{}
     book.Pending = func(Ta int, Pa int) int {
            return Ta - Pa
     }

     // Повторное использование функции
     fmt.Println("Pending articles: ", book.Pending(145,2))
           book.Pending = func(Ta int, Pa int) int {
            return Ta + Pa
           }

       fmt.Println("Pending articles: ", book.Pending(145,2))
  }
```

