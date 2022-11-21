# Stack 



```go
package main

import (
	"fmt"
	"memos/memo"
)

// Main
func main() {
	// Push Pop
	PushPopup()
}  

func PushPopup() {
	var mStack memo.Stack
	mStack.Push("Say")
	mStack.Push("Other World")
	mStack.Push("Hi All People")
	mStack.Push("I Glad to see you")
	mStack.Push(92)

	itm, _ := mStack.Pop()
	fmt.Println("-->", itm)
	itm, _ = mStack.Pop()
	fmt.Println("-->", itm)

	for {
		item, err := mStack.Pop()
		if err != nil {
			//fmt.Println(err.Error())
			break
		}
		fmt.Println(item)
	}
}
```
