package main

import "fmt"

// linear data structure

type Stack struct {
	top    *Node
	bottom *Node
	length int
}

type Node struct {
	value string
	next  *Node
}

func main() {
	s := NewStack("berries")
	fmt.Println(s.toString())

	s.push("coconut")
	fmt.Println(s.toString())

	s.push("apple")
	fmt.Println(s.toString())

	fmt.Println("Peeking: ", s.peek())

	s.pop()
	fmt.Println("poping top element from stackWithLinkedList..: ")
	fmt.Println(s.toString())

	if !s.isEmpty() {
		fmt.Println("Stack is not empty")
	}
	fmt.Println(s.toString())

	s.pop()
	fmt.Println(s.toString())

	s.pop()

	if s.isEmpty() {
		fmt.Println("Stack is empty")
	}

	fmt.Println(s.toString())

	s.pop()
}

// peek returns the first element in a stackWithLinkedList
func (s *Stack) peek() *Node {
	return s.top
}

// push adds an element on top of the stackWithLinkedList
func (s *Stack) push(value string) {
	n := NewNode(value)

	// point new node to top of stackWithLinkedList
	n.next = s.top

	// make new node the top of the stackWithLinkedList now
	s.top = n

	// increase length of stackWithLinkedList
	s.length++
}

// pop deletes the first element in the stackWithLinkedList
func (s *Stack) pop() {
	// check stackWithLinkedList no empty
	if s.length == 0 {
		return
	}

	// check if only one node is left in stackWithLinkedList
	if s.length == 1 {
		s.top = nil
		s.bottom = nil
		s.length--
		return
	}

	// get node after top
	nextNode := s.top.next

	// unlink top node
	s.top.next = nil

	// make next node the top
	s.top = nextNode

	// update size of stackWithLinkedList
	s.length--
}

// isEmpty checks if stackWithLinkedList is empty
func (s *Stack) isEmpty() bool {
	if s.length == 0 {
		return true
	}

	return false
}

// NewStack creates and return a new stackWithLinkedList
func NewStack(value string) *Stack {
	n := NewNode(value)

	return &Stack{
		top:    n,
		bottom: n,
		length: 1,
	}
}

// NewNode creates and returns a new node
func NewNode(value string) *Node {
	return &Node{
		value: value,
		next:  nil,
	}
}

// toString generates a string representation of a Node
func (n *Node) toString() string {
	pointer := "nil"

	if n.next != nil {
		pointer = n.next.value
	}

	return "[value: " + n.value + ", pointer: *" + pointer + "]"
}

// getNodeAtIndex returns the node at the given index - O(n)
func (s *Stack) toString() string {
	stack := ""
	arrow := ""

	for pointer := s.top; pointer != nil; pointer = pointer.next {
		if len(stack) > 0 {
			arrow = " --> "
		}

		stack = stack + arrow + pointer.toString()
	}

	return stack
}
