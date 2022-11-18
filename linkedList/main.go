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

	l.delete(5)
	fmt.Println(l.toString())

	l.delete(6)
	fmt.Println(l.toString())

	l.delete(0)
	fmt.Println(l.toString())

	l.reverse()
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

// append add a Node at the end of the linked list - O(1)
func (l *LinkedList) append(value string) {
	// create new Node
	n := NewNode(value)

	// check linked list is empty
	if l.length == 0 {
		l.head = n
		l.tail = n
		l.length++
		return
	}

	// point current tail next pointer to new Node
	l.tail.next = n

	// make new Node the tail
	l.tail = n

	// increment list number
	l.length++
}

// prepend adds a Node at the start of the linked list - O(1)
func (l *LinkedList) prepend(value string) {
	// create new Node
	n := NewNode(value)

	// check linked list is empty
	if l.length == 0 {
		l.head = n
		l.tail = n
		l.length++
		return
	}

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

	// traverse linked list and find the node before the index
	preNode := l.getNodeAtIndex(index - 1)

	// make new node pointer to node at current index
	n := NewNode(value)
	n.next = preNode.next

	// update and assign previous node to newly created node
	preNode.next = n

	// update node counter
	l.length++
}

// delete removes a node at a given index - O(n)
func (l *LinkedList) delete(index int) {
	// check index is out of range
	if index > l.length {
		return
	}

	// check index is the head then make the next node the head
	if index == 0 {
		nodeToDelete := l.head

		// make the next node the head
		l.head = nodeToDelete.next

		// remove link of node to be deleted
		nodeToDelete.next = nil

		// update counter
		l.length--
		return
	}

	// traverse linked list and find the node before the index
	preNode := l.getNodeAtIndex(index - 1)
	nodeToDelete := preNode.next

	// link preNode to node after node to be deleted index
	preNode.next = nodeToDelete.next

	// check pre node is nil then make it the tail
	if preNode.next == nil {
		l.tail = preNode
	}

	// remove node to delete next link to other node
	nodeToDelete.next = nil

	// update counter
	l.length--
}

// reverse re-orders the linked list
func (l *LinkedList) reverse() {
	// check list has one or less nodes
	if l.length <= 1 {
		return
	}

	// get first and second nodes
	firstNode := l.head
	secondNode := firstNode.next

	// loop through list until second node is empty and we reach the tail
	for secondNode != nil {
		// get a copy of third node
		thirdNode := secondNode.next

		// point second node to head
		secondNode.next = firstNode

		// move first node forward and point it to second node
		firstNode = secondNode

		// update second node to next node
		secondNode = thirdNode
	}

	// update head and tail
	l.tail = l.head
	l.head.next = nil
	l.head = firstNode
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

	return "[value: " + n.value + ", pointer: *" + pointer + "]"
}

// getNodeAtIndex returns the node at the given index - O(n)
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
