// This package implements heterogeneous stack which allows you to store any type of value
// as compared to the built-in slice and maps which will only allow specific types.
// Its a thread safe so it works pretty well with multiple goroutines.
package stack

import (
	"errors"
	"sync"
)

// Stack will accept any value type including int, float, string, slice all together.
type Stack []interface{}

// Declares new error empty Stack
var emptyStack = errors.New("stack.go : Stack is empty")

// Mutex will synchronize access to stack
var mutex = &sync.RWMutex{}

// It will return number of element in stack.
func (s Stack) Len() int {
	mutex.RLock()
	defer mutex.RUnlock()
	return len(s)
}

// It will return capacity of stack.
func (s Stack) Cap() int {
	mutex.RLock()
	defer mutex.RUnlock()
	return cap(s)
}

// It will verify whether the stack is empty or not.
func (s Stack) IsEmpty() bool {
	mutex.RLock()
	defer mutex.RUnlock()

	if len(s) == 0 {
		return true
	}
	return false
}

// It will push elements in stack.
func (s *Stack) Push(v interface{}) {
	mutex.Lock()
	defer mutex.Unlock()
	*s = append(*s, v)
}

// It gives first element from stack.
func (s Stack) Top() (interface{}, error) {
	mutex.RLock()
	defer mutex.RUnlock()

	if len(s) == 0 {
		return nil, emptyStack
	}
	return s[len(s)-1], nil
}

// It will remove element from stack.
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
