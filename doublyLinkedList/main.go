package main

import "fmt"

// Doubly linked list data structure
// e.g. [*nil <- :pre, value: pineapple, next -> : *apples] <===> [*pineapple <- :pre, value: apples, next -> : *nil]

type Node struct {
	value    string
	preNode  *Node
	nextNode *Node
}

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func main() {
	l := NewLinkedList("apples")
	fmt.Println(l.toString())

	l.prepend("pineapple") // start of list  - O(1)
	fmt.Println(l.toString())

	l.prepend("coconut") // start of list  - O(1)
	fmt.Println(l.toString())

	l.append("berries") // end of list  - O(1)
	fmt.Println(l.toString())

	l.append("grapes") // end of list  - O(1)
	fmt.Println(l.toString())

	l.insert("pear", 1) // middle list  - O(n)
	fmt.Println(l.toString())

	l.insert("oranges", 4) // middle list  - O(n)
	fmt.Println(l.toString())

	l.delete(2) // middle list  - O(n)
	fmt.Println(l.toString())

	l.delete(4) // last list - O(1)
	fmt.Println(l.toString())

	l.delete(0) // start of list - O(1)
	fmt.Println(l.toString())
}

// NewNode makes a new node
func NewNode(value string) *Node {
	return &Node{
		value:    value,
		preNode:  nil,
		nextNode: nil,
	}
}

// NewLinkedList make a new linked list with one node
func NewLinkedList(value string) *LinkedList {
	// create a new node
	n := NewNode(value)

	return &LinkedList{
		head:   n,
		tail:   n,
		length: 1,
	}
}

// prepend adds a new node to the start of the linked list  - O(1)
func (l *LinkedList) prepend(value string) {
	n := NewNode(value)

	// link new node to current head
	n.nextNode = l.head

	// link heat node to new node
	l.head.preNode = n

	// make new node the head
	l.head = n

	// increment list size
	l.length++
}

// append adds a new node the end of the linked list - O(1)
func (l *LinkedList) append(value string) {
	n := NewNode(value)

	// point previous node (current tail) to new node
	l.tail.nextNode = n

	// point new node to tail node
	n.preNode = l.tail

	// make new node the tail
	l.tail = n

	// increment list size
	l.length++
}

// insert adds a new node in a linked list - O(n)
func (l *LinkedList) insert(value string, index int) {
	// check node is being added at the start
	if index == 0 {
		l.prepend(value)
		return
	}

	// check node is being added at the end of the linked list
	if index >= l.length {
		l.append(value)
		return
	}

	// create a new node
	n := NewNode(value)

	// find node at given index
	nodeAtIndex := l.getNodeAtIndex(index)

	// point new node previous pointer to previous node of current index node
	n.preNode = nodeAtIndex.preNode

	// point new node next pointer to node at current index
	n.nextNode = nodeAtIndex

	// update current node index next pointer to point to new node
	nodeAtIndex.preNode.nextNode = n

	// update node at index pre pointer to point to new node
	nodeAtIndex.preNode = n

	// increment list size
	l.length++
}

// delete unlinks a node from the linked list - O(n)
func (l *LinkedList) delete(index int) {
	// check if index out of range
	if index > l.length {
		return
	}

	// check if node being deleted is the head
	if index == 0 {
		// next node after head
		nextNode := l.head.nextNode
		// unlink head next node
		l.head.nextNode = nil
		// update heads next node to not point back to head node
		l.head.preNode = nil
		// update next node previous pointer to not point to node to be deleted
		nextNode.preNode = nil
		// make heads next node the head
		l.head = nextNode
		// update linked list size
		l.length--

		return
	}

	// check if node being deleted is the tail
	if index == l.length {
		// get tails previous node
		previousNode := l.tail.preNode

		// unlink previous node to tail
		previousNode.nextNode = nil

		// unlink current tail from linking to previous node
		l.tail.preNode = nil

		// make previous node the tail
		l.tail = previousNode

		// update linked list size
		l.length--

		return
	}

	// find the node in middle to be deleted
	nodeToDelete := l.getNodeAtIndex(index)
	nextNode := nodeToDelete.nextNode
	previousNode := nodeToDelete.preNode

	// point next node to previous node
	nextNode.preNode = previousNode

	// point previous node to next node
	previousNode.nextNode = nextNode

	// unlink index node
	nodeToDelete.nextNode = nil
	nodeToDelete.preNode = nil

	// update linked list size
	l.length--
}

// toString generates a string representation of the linked list
func (l *LinkedList) toString() string {
	list := ""
	arrow := ""

	for pointer := l.head; pointer != nil; pointer = pointer.nextNode {
		if len(list) > 0 {
			arrow = " <===> "
		}

		list = list + arrow + pointer.toString()
	}

	return list
}

// toString generates a string representation of a Node
func (n *Node) toString() string {
	prePointer := "nil"
	nextPointer := "nil"

	if n.nextNode != nil {
		nextPointer = n.nextNode.value
	}

	if n.preNode != nil {
		prePointer = n.preNode.value
	}

	return "[*" + prePointer + " <- :pre, value: " + n.value + ", next -> : *" + nextPointer + "]"
}

// getNodeAtIndex returns the node at the given index - O(n)
func (l *LinkedList) getNodeAtIndex(index int) *Node {
	// start pointer from the head
	pointer := l.head

	// traverse linked list until you reach the index
	for i := 0; i < index; i++ {
		// move pointer to next node
		pointer = pointer.nextNode
	}

	return pointer
}
