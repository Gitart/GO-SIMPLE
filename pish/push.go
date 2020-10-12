package main

import "fmt"

//Stack represents a stack that hold a slice
type Stack struct {
	items []int
}

// Push will add a value at the end (top)
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// Pop will remove the value at the end (top)
func (s *Stack) Pop() int {
	lastIndex := len(s.items) - 1
	toRemove  := s.items[lastIndex]
	s.items   = s.items[:lastIndex]
	return toRemove
}



func main() {
	myStack := Stack{}
	Prn("Check Stack")
	Prn(myStack)
	
	// Push element to array
	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)
	myStack.Push(500)
	myStack.Push(600)
	myStack.Push(700)	
	
	Prn(myStack)
	
	Prn("\n Выдавить последний  элемент")
	myStack.Pop()
	Prn(myStack)
	
	Prn("\n Выдавить еще один последний елемент")
	myStack.Pop()
	Prn(myStack)
}



func Prn(i interface{}){
	fmt.Println(i)
}
