// https://go.dev/play/p/T8WOs9sSw8p

package memo

import (
	"errors"
)

type Stack []interface{}

// Push
func (s *Stack) Push(x interface{}) {
	*s = append(*s, x)
}

// Len
func (s Stack) Len() int {
	return len(s)
}

// Top
func (s Stack) Top() (interface{}, error) {
	if len(s) == 0 {
		return nil, errors.New("ERROR: Can't Top() is empty stack")
	}
	return s[len(s)-1], nil
}

// Pop
func (s *Stack) Pop() (interface{}, error) {
	thestack := *s
	if len(thestack) == 0 {
		return nil, errors.New("ERROR: Can't Pop() an empty stack")
	}
	x := thestack[len(thestack)-1]
	*s = thestack[:thestack.Len()-1]
	return x, nil
}
