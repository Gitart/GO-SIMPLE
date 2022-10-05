# Реализация двусвязного списка в Go

Программа, реализующая двусвязный список на Go, называется doublelyLList.go.
Рассмотрим ее, разделив на пять частей. Общая идея двусвязного списка такая же,
как и у односвязного, но здесь из-за наличия двух указателей в каждом узле списка
придется выполнять больше операций по обслуживанию списка.

```go
package main
import (
 "fmt"
)
type Node struct {
 Value int
 Previous *Node
 Next *Node
}

// В данной части мы видим определение узла двусвязного списка как структуры Go.
// Однако на этот раз по понятным причинам структура имеет два поля указателей.
// Вторая часть файла duplyLList.go содержит следующий код Go:


func addNode(t *Node, v int) int {
 if root == nil {
 t = &Node{v, nil, nil}
 root = t
 return 0
 }
 if v == t.Value {
 fmt.Println("Node already exists:", v)
 return -1
 }
 if t.Next == nil {
 temp := t
 t.Next = &Node{v, temp, nil}
 return -2
 }
 return addNode(t.Next, v)
}


// Как и в случае односвязного списка, каждый новый узел помещается в конец
// текущего двусвязного списка. Однако так поступать не обязательно, если вы хотите
// построить упорядоченный двусвязный список.
// Третья часть doublyLList.go выглядит так:

func traverse(t *Node) {
 if t == nil {
 fmt.Println("-> Empty list!")
 return
 }
 for t != nil {
 fmt.Printf("%d -> ", t.Value)
 t = t.Next
 }
 fmt.Println()
}

func reverse(t *Node) {
 if t == nil {
 fmt.Println("-> Empty list!")
 return
 }
 temp := t
 for t != nil {
 temp = t
 t = t.Next
 }
 for temp.Previous != nil {
 fmt.Printf("%d -> ", temp.Value)
 temp = temp.Previous
 }
 fmt.Printf("%d -> ", temp.Value)
 fmt.Println()
}

// Здесь вы видите код Go для функций traverse() и reverse(). Реализация
// функции traverse() такая же, как и в программе connectedList.go. Однако логика
// функции reverse() очень интересна. Поскольку мы не храним указатель на конец
// двусвязного списка, то нам нужно перейти к концу такого списка, чтобы получить
// доступ к его узлам в обратном порядке.
// Обратите внимание, что Go позволяет писать такой код, как a, b = b, a, чтобы
// поменять местами значения двух переменных, не создавая временную переменную.
// Четвертая часть dublyLList.go содержит следующий код Go:

func size(t *Node) int {
 if t == nil {
 fmt.Println("-> Empty list!")
 return 0
 }
 n := 0
 for t != nil {
 n++
 t = t.Next
 }
 return n
}

func lookupNode(t *Node, v int) bool {
 if root == nil {
 return false
 }
 if v == t.Value {
 return true
 }
 if t.Next == nil {
 return false
 }
 return lookupNode(t.Next, v)
}


// Последний фрагмент файла duplyLList.go содержит следующий код Go:
var root = new(Node)

func main() {
 fmt.Println(root)
 root = nil
 traverse(root)
 addNode(root, 1)
 addNode(root, 1)
 traverse(root)
 addNode(root, 10)
 addNode(root, 5)
 addNode(root, 0)
 addNode(root, 0)
 traverse(root)
 addNode(root, 100)
 fmt.Println("Size:", size(root))
 traverse(root)
 reverse(root)
}
```
