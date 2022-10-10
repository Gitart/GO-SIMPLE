package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

var size = 0
var queue = new(Node)

// Наличие переменной size для хранения количества узлов в очереди удобно, но
// не обязательно. В представленной здесь реализации это сделано для того, чтобы
// упростить работу. Вероятно, вы тоже захотите создать эти поля в своей структуре.
// Вторая часть queue.go содержит следующий код Go:

func Push(t *Node, v int) bool {
	if queue == nil {
		queue = &Node{v, nil}
		size++
		return true
	}
	t = &Node{v, nil}

	t.Next = queue
	queue = t
	size++
	return true
}

// Здесь показана реализация функции Push(), которая сама по себе очень проста.
// Если очередь пуста, то новый узел становится очередью. Если очередь не пуста, то
// создается новый узел, который помещается перед текущей очередью. После этого
// только что созданный узел становится заголовком очереди.
// Третья часть queue.go содержит следующий код Go:

func Pop(t *Node) (int, bool) {
	if size == 0 {
		return 0, false
	}
	if size == 1 {
		queue = nil
		size--
		return t.Value, true
	}
	temp := t
	for (t.Next) != nil {
		temp = t
		t = t.Next
	}
	v := (temp.Next).Value
	temp.Next = nil
	size--
	return v, true
}

// В этом коде показана реализация функции Pop(), которая удаляет из очереди
// самый старый элемент. Если очередь пуста (size == 0), то из нее ничего не извлекается.
// Если в очереди есть только один узел, то извлекается значение этого узла,
// после чего очередь становится пустой. В противном случае извлекается последний
// элемент очереди, удаляется последний узел и изменяются необходимые указатели,
// после чего возвращается нужное значение.
// Четвертая часть queue.go содержит следующий код Go:

func traverse(t *Node) {
	if size == 0 {
		fmt.Println("Empty Queue!")
		return
	}
	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}

// Функция traverse() не обязательна для работы с очередью, но она удобна для
// просмотра всех узлов очереди.
// Последний фрагмент кода queue.go содержит следующий код Go:

func main() {
	queue = nil
	Push(queue, 10)
	fmt.Println("Size:", size)
	traverse(queue)
	v, b := Pop(queue)
	if b {
		fmt.Println("Pop:", v)
	}
	fmt.Println("Size:", size)
	for i := 0; i < 5; i++ {
		Push(queue, i)
	}
	traverse(queue)
	fmt.Println("Size:", size)
	v, b = Pop(queue)
	if b {
		fmt.Println("Pop:", v)
	}
	fmt.Println("Size:", size)
	v, b = Pop(queue)
	if b {
		fmt.Println("Pop:", v)
	}
	fmt.Println("Size:", size)
	traverse(queue)
}
