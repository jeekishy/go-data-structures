package main

import "fmt"

type Node struct {
	value string
	next  *Node
}

type Queue struct {
	first  *Node
	last   *Node
	length int
}

func main() {
	q := NewQueue("apples")
	fmt.Println(q.toString())

	fmt.Println(q.peek())

	q.enqueue("coconut")
	fmt.Println(q.toString())

	q.enqueue("berries")
	fmt.Println(q.toString())

	q.dequeue()
	fmt.Println(q.toString())

	q.dequeue()
	q.dequeue()
	q.dequeue()
	fmt.Println("Queue:", q.toString())

}

func NewNode(value string) *Node {
	return &Node{
		value: value,
		next:  nil,
	}
}

func NewQueue(value string) *Queue {
	n := NewNode(value)

	return &Queue{
		first:  n,
		last:   n,
		length: 1,
	}
}

// returns the first value in the queue
func (q *Queue) peek() string {
	return q.first.value
}

// enqueue adds a node to the start/front of the queue
//
//	 ("apples") ---> ("bananas")
//	     ^				  ^
//		 |				  |
//		new node		first node
func (q *Queue) enqueue(value string) {
	// create a new node
	n := NewNode(value)

	if q.length == 0 {
		q.last = n
		q.first = n
		q.length = 1
		return
	}

	// pointer new node to back of queue
	n.next = q.first

	// make new node the first element
	q.first = n

	// increase length
	q.length++
}

// dequeue removes the last node from the end of the queue
func (q *Queue) dequeue() {
	if q.isEmpty() {
		return
	}

	// when there is only one node in queue just delete it
	if q.length == 1 {
		q.first = nil
		q.last = nil
		q.length = 0
		return
	}

	// get node before last
	nodeBeforeLast := q.getNodeAtIndex(q.length - 2)

	// unlink node before last node
	//  ("apples") ---> ("bananas") ---> ("coconut") -X-> ("berries")
	//      ^								 				  ^
	//		|								 				  |
	//	  first 											last
	nodeBeforeLast.next = nil

	// make node the last node in the queue
	// unlink node before last node
	//  ("apples") ---> ("bananas") ---> ("coconut")
	//      ^								  ^
	//		|								  |
	//	  first 							 last
	q.last = nodeBeforeLast

	// update length
	q.length--
}

// isEmpty checks if queue is empty
func (q *Queue) isEmpty() bool {
	if q.length == 0 {
		return true
	}
	return false
}

// toString generates a string representation of a Node
func (n *Node) toString() string {
	pointer := "nil"

	if n.next != nil {
		pointer = n.next.value
	}

	return "[value: " + n.value + ", pointer: *" + pointer + "]"
}

// toString generates a string representation of the queue
func (q *Queue) toString() string {
	queue := ""
	arrow := ""

	for pointer := q.first; pointer != nil; pointer = pointer.next {
		if len(queue) > 0 {
			arrow = " --> "
		}

		queue = queue + arrow + pointer.toString()
	}

	return queue
}

// getNodeAtIndex returns the node at the given index - O(n)
func (q *Queue) getNodeAtIndex(index int) *Node {
	// start from head of linked list
	pointer := q.first

	// move pointer to index we want
	for i := 0; i < index; i++ {
		// move pointer to next node
		pointer = pointer.next
	}

	return pointer
}
