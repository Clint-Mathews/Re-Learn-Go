package main

import (
	"fmt"
)

var count int

// Node represents the component in a binary search tree
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Insert will add a node to the tree
func (n *Node) Insert(k int) {
	if n.Key < k {
		// Move Right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		// Move Left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	} else {
		// Key already present
	}
}

// Search

func (n *Node) Search(k int) bool {
	count++
	if n == nil {
		return false
	}
	if n.Key < k {
		// Right
		return n.Right.Search(k)
	} else if n.Key > k {
		// Left
		return n.Left.Search(k)
	}
	return true
}

func main() {
	tree := &Node{
		Key: 10,
	}
	tree.Insert(232)
	tree.Insert(7)
	tree.Insert(4)
	tree.Insert(5)
	tree.Insert(34)
	tree.Insert(78)
	fmt.Println(tree)
	fmt.Println(tree.Search(78), count)
}
