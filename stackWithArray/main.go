package main

import "fmt"

// linear data structure

type Stack struct {
	data []string
}

func main() {
	s := NewStack("apples")
	fmt.Println(s.data)

	if !s.isEmpty() {
		fmt.Println("Stack is NOT empty")
	}

	s.push("pineapple")
	fmt.Println(s.data)

	s.push("coconut")
	fmt.Println(s.data)

	s.pop()
	fmt.Println(s.data)

	fmt.Println(s.peek())

	s.pop()
	s.pop()

	fmt.Println(s.peek())

	if s.isEmpty() {
		fmt.Println("Stack is empty")
	}
}

func NewStack(value string) *Stack {
	return &Stack{
		data: []string{
			value,
		},
	}
}

// push appends a value to the slice
func (s *Stack) push(value string) {
	s.data = append(s.data, value)
}

// pop removes last added item from slice
func (s *Stack) pop() {
	l := len(s.data)

	// when empty escape
	if l == 0 {
		return
	}

	s.data = s.data[:l-1]
}

// peek returns the last added element
func (s *Stack) peek() string {
	l := len(s.data)

	if l == 0 {
		return ""
	}

	return s.data[l-1]
}

// isEmpty validates if stack is empty
func (s *Stack) isEmpty() bool {
	if len(s.data) == 0 {
		return true
	}

	return false
}
