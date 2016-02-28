# gostack

[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/shamsher31/gostack)
[![Build Status](https://travis-ci.org/shamsher31/gostack.svg)](https://travis-ci.org/shamsher31/gostack)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](license)

This package implements heterogeneous stack which allows you to store any type of value
as compared to the built-in slice and maps which will only allow specific types. Its a thread safe so it works pretty well with multiple goroutines.  

### How to install
```go
go get github.com/shamsher31/gostack
```

### How to use
```go
package main

import (
  "fmt"
  "github.com/shamsher31/gostack"
)

func main() {
  // Declare new stack
  var myStack stack.Stack

  fmt.Println("Stack is empty : ", myStack.IsEmpty())

  //Push some elements of different types in Stack
  myStack.Push(10)
  myStack.Push([]string{"Shamsher", "Ansari"})
  myStack.Push(13.5)
  myStack.Push("My Awesome stack")

  fmt.Println("Elements of stack : ", myStack)

  fmt.Println("Stack is empty : ", myStack.IsEmpty())

  if myStack.IsEmpty() == false {
    fmt.Println("Total elements in stack : ", myStack.Len())
  }

  top, err := myStack.Top()
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println("Top of the stack : ", top)

  elem, err := myStack.Pop()
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println("Poped element : ", elem)
  fmt.Println("Elements of stack : ", myStack)
}

### Why
This package is born while learning and experimenting with golang to understand how pointers, mutual exclusion locks and interface works together. 

### License
MIT © [Shamsher Ansari](https://github.com/shamsher31)
