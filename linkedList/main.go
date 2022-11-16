package main

import (
	"fmt"
)

// Node linked list data structure
// e.g. [v: pears, p: *apple] --> [v: apple, p: *oranges] --> [v: oranges, p: *nil]
// Node is a single element of linked list structure
type Node struct {
	value string
	next  *Node
}

// LinkedList holds all the nodes associated with the linked list
type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func main() {
	l := NewLinkedList("apple")
	fmt.Println(l.toString())

	l.append("oranges")
	fmt.Println(l.toString())

	l.prepend("pears")
	fmt.Println(l.toString())

	l.append("avocado")
	fmt.Println(l.toString())

	l.insert("coconut", 2)
	fmt.Println(l.toString())

	l.insert("berries", 0)
	fmt.Println(l.toString())
}

func NewLinkedList(value string) *LinkedList {
	// create first Node
	n := NewNode(value)

	// creates linked list
	return &LinkedList{
		head:   n,
		tail:   n,
		length: 1,
	}
}

func NewNode(value string) *Node {
	return &Node{
		value: value,
		next:  nil,
	}
}

// append add a Node at the end of the linked list
func (l *LinkedList) append(value string) {
	// create new Node
	n := NewNode(value)

	// pointer current tail next pointer to new Node
	l.tail.next = n

	// make new Node the tail
	l.tail = n

	// increment list number
	l.length++
}

// prepend adds a Node at the start of the linked list
func (l *LinkedList) prepend(value string) {
	// create new Node
	n := NewNode(value)

	// update node next pointer
	n.next = l.head

	// make new Node the head
	l.head = n

	// increment list number
	l.length++
}

// insert adds a node at a given index - O(n) / O(1) if adding to start or end
func (l *LinkedList) insert(value string, index int) {
	// when first index perform prepend
	if index == 0 {
		l.prepend(value)
		return
	}

	// check index out of range and add to end of list
	if index >= l.length {
		l.append(value)
		return
	}

	// traverse linked list and find starting node
	preNode := l.getNodeAtIndex(index - 1)

	// make new node pointer to node at current index
	n := NewNode(value)
	n.next = preNode.next

	// update and assign previous node to newly created node
	preNode.next = n
}

// toString generates a string representation of the linked list
func (l *LinkedList) toString() string {
	list := ""
	arrow := ""

	for pointer := l.head; pointer != nil; pointer = pointer.next {
		if len(list) > 0 {
			arrow = " --> "
		}

		list = list + arrow + pointer.toString()
	}

	return list
}

// toString generates a string representation of a Node
func (n *Node) toString() string {
	pointer := "nil"

	if n.next != nil {
		pointer = n.next.value
	}

	return "[v: " + n.value + ", p: *" + pointer + "]"
}

// getNodeAtIndex returns the node at the given index
func (l *LinkedList) getNodeAtIndex(index int) *Node {
	// start from head of linked list
	pointer := l.head

	// move pointer to index we want
	for i := 0; i < index; i++ {
		// move pointer to next node
		pointer = pointer.next
	}

	return pointer
}
