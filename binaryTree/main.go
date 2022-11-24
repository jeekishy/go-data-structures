package main

import (
	"errors"
	"fmt"
)

// Binary tree - each node can have 0 to 2 nodes
// and each child can only have one parent
// left nodes decreasing and right is increasing in value
// O(log (N)) - complexity

type Node struct {
	value int
	left  *Node
	right *Node
}

type BinaryTree struct {
	parent *Node
	length int
}

func main() {
	// create a binary tree with just one node and amke it the parent
	b := NewBinaryTree(9)

	b.insert(4)
	b.insert(6)
	b.insert(15)
	b.insert(120)
	b.insert(40)
	b.insert(56)
	b.insert(58)
	b.insert(3)
	b.insert(2)
	b.insert(1)

	n, err := b.lookup(40)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Node found:", n)
	}

	n, err = b.lookup(58)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Node found:", n)
	}

	n, err = b.lookup(123)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Node found:", n)
	}

	n, err = b.lookup(120)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Node found:", n)
	}

	b.delete(120)
	n, err = b.lookup(120)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Node found:", n)
	}

}

func NewBinaryTree(value int) *BinaryTree {
	n := NewNode(value)

	return &BinaryTree{
		parent: n,
		length: 1,
	}
}

func NewNode(value int) *Node {
	return &Node{
		value: value,
		left:  nil,
		right: nil,
	}
}

func (b *BinaryTree) insert(value int) {
	// check value already in binary tree
	if b.parent.value == value {
		return
	}

	// create a new node to add to tree
	n := NewNode(value)

	// keep a pointer of the current node
	p := b.parent

	for {
		// dnt insert duplicates
		if value == p.value {
			break
		}

		// check any value on the right of tree
		if value > p.value {
			if p.right == nil {
				p.right = n
				b.length++
				break
			}
			p = p.right
			continue
		}

		// else check value on left of tree
		if p.left == nil {
			p.left = n
			b.length++
			break
		}
		p = p.left
	}
}

func (b *BinaryTree) lookup(value int) (*Node, error) {
	// check if value is head
	if value == b.parent.value {
		return b.parent, nil
	}

	// keep track of node pointer
	p := b.parent

	for {
		fmt.Println("Node: ", p)

		// check if current pointer's value is same as the value wanted
		if value == p.value {
			return p, nil
		}

		// check right node
		if value > p.value {
			// check if next node is empty
			if p.right == nil {
				return nil, errors.New("node not found")
			}
			// update pointer
			p = p.right
			continue
		}

		// else node would be left
		if p.left == nil {
			return nil, errors.New("node not found")
		}

		p = p.left
	}
}

func (b *BinaryTree) delete(value int) bool {
	// keep track of parent and current node
	parentNode := b.parent
	currentNode := b.parent

	for {
		// check if current pointer's value is same as the value wanted
		if value == currentNode.value {
			break
		}

		// check right node
		if value > currentNode.value {
			// check if next node is empty
			if currentNode.right == nil {
				break
			}

			// update node pointers
			parentNode = currentNode
			currentNode = currentNode.right
			continue
		}

		// else node would be left
		if currentNode.left == nil {
			break
		}

		// update node pointers
		parentNode = currentNode
		currentNode = currentNode.left
	}

	// when node has no children
	// just delete it by make parent node nil
	if currentNode.right == nil && currentNode.left == nil {
		// find parent pointing node and delete
		if parentNode.right == currentNode {
			parentNode.right = nil
		}

		if parentNode.left == currentNode {
			parentNode.left = nil
		}

		return true
	}

	// when node has no right child
	// set parent of to be deleted node to left child of to be deleted node
	if currentNode.right == nil {
		// find parent pointing node and delete
		if parentNode.right == currentNode {
			parentNode.right = currentNode.left
		}

		if parentNode.left == currentNode {
			parentNode.left = currentNode.left
		}

		return true
	}

	// when node has no left child
	// set parent of to be deleted node to right child of to be deleted node
	if currentNode.left == nil {
		// find parent pointing node and delete
		if parentNode.right == currentNode {
			parentNode.right = currentNode.right
		}

		if parentNode.left == currentNode {
			parentNode.left = currentNode.right
		}

		return true
	}

	return false
}

func (b *BinaryTree) toString() string {
	return ""
}
