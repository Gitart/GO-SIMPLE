// An example of type assertions for empty interface (i.e., interface{})

package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	k string
	v float64
}

func main() {
	{
		l := list.New()
		l.PushBack(Node{"Hello", -222.0})
		l.PushBack(Node{"world", -42.1})

		for e := l.Front(); e != nil; e = e.Next() {
			// Type assertions.
			n := e.Value.(Node)
			fmt.Printf("%T (k,v) = (%s,%g)\n", n, n.k, n.v)
		}
	}

	// reference types
	{
		rl := list.New()
		rl.PushBack(&Node{"Hello", -222.0})
		rl.PushBack(&Node{"world", -42.1})

		for e := rl.Front(); e != nil; e = e.Next() {
			// Type assertions.
			n := e.Value.(*Node) // pointer to Node
			fmt.Printf("%T (k,v) = (%s,%g)\n", n, n.k, n.v)
		}
	}
}
