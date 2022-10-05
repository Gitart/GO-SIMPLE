## Использование пакета container/list

В этом подразделе продемонстрирована работа пакета container/list на примере Go-кода из файла conList.go, который мы разделим на три части.

**Внимание!** В пакете container/list реализован двусвязный список.


```go 
package main
import (
 "container/list"
 "fmt"
 "strconv"
)

func printList(l *list.List) {
 for t := l.Back(); t != nil; t = t.Prev() {
 fmt.Print(t.Value, " ")
 }
 fmt.Println()
 for t := l.Front(); t != nil; t = t.Next() {
 fmt.Print(t.Value, " ")
 }
 fmt.Println()
}

// Здесь вы видите функцию printList(), которая позволяет выводить на экран
// содержимое переменной list.List, переданной в виде указателя. В коде Go показано, как вывести элементы list.List, начиная с первого и заканчивая последним,
// и в обратном порядке. Обычно в программах нужно использовать только один из

//  Куча гарантирует, что нулевой элемент (*myHeap)[0] всегда будет минимальным.
// Для извлечения из кучи элементов в соответствии с приоритетом нужно использовать функцию heap.Pop().
// двух методов. Функции Prev() и Next() позволяют перебирать элементы списка в прямом и обратном порядке.

// Второй фрагмент кода conList.go выглядит так:

func main() {
 values := list.New()
 e1 := values.PushBack("One")
 e2 := values.PushBack("Two")
 values.PushFront("Three")
 values.InsertBefore("Four", e1)
 values.InsertAfter("Five", e2)
 values.Remove(e2)
 values.Remove(e2)
 values.InsertAfter("FiveFive", e2)
 values.PushBackList(values)
 printList(values)
 values.Init()

// Функция list.PushBack() позволяет вставлять объект в конец связного списка,
// а функция list.PushFront() — в начало списка. Обе функции возвращают вставленный в список элемент.
// Если вы хотите вставить новый элемент после определенного элемента, то
// следует использовать функцию list.InsertAfter(). Аналогично, для того чтобы
// вставить элемент перед конкретным элементом, необходимо применить функцию
// list.InsertBefore(). Если такой элемент не существует, то список не изменится.

// Функция list.PushBackList() вставляет копию существующего списка в конец другого списка, а list.PushFrontList() помещает копию существующего списка в начало другого списка. Функция list.Remove() удаляет из списка заданный элемент.
// Обратите внимание на использование функции values.Init(), которая либо
// очищает существующий список, либо инициализирует новый список.
// Последняя часть conList.go содержит следующий код Go:

 fmt.Printf("After Init(): %v\n", values)
 for i := 0; i < 20; i++ {
 values.PushFront(strconv.Itoa(i))
 }
 printList(values)
}

// Здесь мы создаем новый список с помощью цикла for. Функция strconv.Itoa()
// преобразует целочисленное значение в строку.
// Таким образом, функции пакета container/list достаточно просты и их использование не должно вызывать затруднений.
```

Выполнение conList.go приведет к следующим результатам:

```
$ go run conList.go
Five One Four Three Five One Four Three
Three Four One Five Three Four One Five
After Init(): &{{0xc420074180 0xc420074180 <nil> <nil>} 0}
0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19
19 18 17 16 15 14 13 12 11 10 9 8 7 6 5 4 3 2 1 0
```
