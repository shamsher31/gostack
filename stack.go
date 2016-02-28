// stack implements heterogeneous stack which allows you to store any type of value
// in stack including int, float, string and slice
package stack

import (
	"errors"
	"sync"
)

// Stack will accept any value type
type Stack []interface{}

// Declares new error empty Stack
var emptyStack = errors.New("stack.go : Stack is empty")

// Mutex will synchronize access to stack
var mutex = &sync.Mutex{}

// Len will return number of element in stack
func (s Stack) Len() int {
	mutex.Lock()
	defer mutex.Unlock()
	return len(s)
}

// Cap will return capacity of stack
func (s Stack) Cap() int {
	mutex.Lock()
	defer mutex.Unlock()
	return cap(s)
}

// IsEmpty will verify whether thestack is empty or not
func (s Stack) IsEmpty() bool {
	mutex.Lock()
	defer mutex.Unlock()

	if len(s) == 0 {
		return true
	}
	return false
}

// Pushes elements in stack
func (s *Stack) Push(v interface{}) {
	mutex.Lock()
	defer mutex.Unlock()
	*s = append(*s, v)
}

// Gives first element from stack
func (s Stack) Top() (interface{}, error) {
	mutex.Lock()
	defer mutex.Unlock()

	if len(s) == 0 {
		return nil, emptyStack
	}
	return s[len(s)-1], nil
}

// Pop removes element from stack
func (s *Stack) Pop() (interface{}, error) {
	mutex.Lock()
	defer mutex.Unlock()

	ls := *s
	if len(ls) == 0 {
		return nil, emptyStack
	}

	v := ls[len(ls)-1]
	*s = ls[:len(ls)-1]

	return v, nil
}
