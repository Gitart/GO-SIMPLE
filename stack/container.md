# Пакет container

Пакет **container** поддерживает три структуры данных: кучу, список и кольцо.
Эти структуры данных реализованы в пакетах **container/heap, container/list и container/ring** соответственно.
На случай, если вы не знаете: кольцо представляет собой циклический список, так что последний элемент кольца указывает на его первый элемент. 
По сути, это означает, что все узлы кольца эквивалентны, кольцо не имеет ни начала, ни конца.
В результате можно пройти через все кольцо, начиная с любого элемента.
В следующих трех подразделах описаны все пакеты, входящие в состав container.

**Совет:** если функциональность стандартного Go-пакета container соответствует вашим потребностям, используйте его; в противном случае лучше реализовать
и использовать собственные структуры данных.

```go
package main
import (
 "container/heap"
 "fmt"
)
type heapFloat32 []float32

//Вторая часть conHeap.go содержит следующий код Go:
func (n *heapFloat32) Pop() interface{} {
 old := *n
 x := old[len(old)-1]
 new := old[0 : len(old)-1]
 *n = new
 return x
}

func (n *heapFloat32) Push(x interface{}) {
 *n = append(*n, x.(float32))
}

// Здесь мы определили две функции с именами Pop() и Push(), которые используются для того, чтобы соответствовать интерфейсу1
// Чтобы добавлять и удалять элементы в куче, необходимо вызывать функции heap.Push() и heap.Pop() соответственно.


// Третий фрагмент кода conHeap.go содержит следующий код Go:

func (n heapFloat32) Len() int {
 return len(n)
}

func (n heapFloat32) Less(a, b int) bool {
 return n[a] < n[b]
}

func (n heapFloat32) Swap(a, b int) {
 n[a], n[b] = n[b], n[a]
}

// В этой части реализованы три функции, необходимые для совместимости с интерфейсом sort.Interface.
// Четвертая часть conHeap.go выглядит так:

func main() {
 myHeap := &heapFloat32{1.2, 2.1, 3.1, -100.1}
 heap.Init(myHeap)
 size := len(*myHeap)
 fmt.Printf("Heap size: %d\n", size)
 fmt.Printf("%v\n", myHeap)
  
  // И последняя часть кода conHeap.go выглядит следующим образом:
 myHeap.Push(float32(-100.2))
 myHeap.Push(float32(0.2))
 fmt.Printf("Heap size: %d\n", len(*myHeap))
 fmt.Printf("%v\n", myHeap)
 heap.Init(myHeap)
 fmt.Printf("%v\n", myHeap)
}

// В этой части conHeap.go мы добавляем в кучу myHeap два новых элемента, используя функцию heap.Push(). Однако, для того чтобы восстановить правильную
// упорядоченность кучи, нужно снова вызвать heap.Init()1
```

Выполнение conHeap.go приведет к следующим результатам:
```
$ go run conHeap.go
Heap size: 4
&[-100.1 1.2 3.1 2.1]
Heap size: 6
&[-100.1 1.2 3.1 2.1 -100.2 0.2]
&[-100.2 -100.1 0.2 2.1 1.2 3.1]
````

В приведенном примере для добавления в кучу использовалась функция myHeap.Push(),
которая не поддерживает упорядоченность кучи. Это и привело к необходимости повторной инициации кучи функцией heap.Init(). Для автоматического поддержания
упорядоченности кучи лучше использовать функцию heap.Push().


