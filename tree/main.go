package main

import "fmt"

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

type BinaryTree struct {
	Head *Node
}

var count int

// Insert

func (tree *BinaryTree) Insert(val int) {
	if tree.Head == nil {
		tree.Head = &Node{Key: val}
	} else {
		tree.Head.InsertBT(val)
	}
}

func (node *Node) InsertBT(val int) {
	if node.Key > val {
		if node.Left != nil {
			node.Left.InsertBT(val)
		} else {
			node.Left = &Node{Key: val}
		}
	} else if node.Key < val {
		if node.Right != nil {
			node.Right.InsertBT(val)
		} else {
			node.Right = &Node{Key: val}
		}
	}
}

func (node *Node) SearchBT(k int) bool {
	count += 1
	if node == nil {
		return false
	}

	if node.Key < k {
		return node.Right.SearchBT(k)
	} else if node.Key > k {
		return node.Left.SearchBT(k)
	}

	return true
}

func (node *Node) PrintInOrder() {
	if node.Left != nil {
		node.Left.PrintInOrder()
	}
	fmt.Println(node.Key)
	if node.Right != nil {
		node.Right.PrintInOrder()
	}
}

func (node *Node) PrintPreOrder() {
	fmt.Println(node.Key)
	if node.Left != nil {
		node.Left.PrintPreOrder()
	}
	if node.Right != nil {
		node.Right.PrintPreOrder()
	}
}

func (node *Node) PrintPostOrder() {

	if node.Left != nil {
		node.Left.PrintPostOrder()
	}
	if node.Right != nil {
		node.Right.PrintPostOrder()
	}
	fmt.Println(node.Key)
}

func main() {
	tree := &BinaryTree{}
	tree.Insert(2)
	tree.Insert(1)
	tree.Insert(7)
	tree.Insert(4)
	tree.Insert(9)
	fmt.Println(tree.Head)
	fmt.Println("Search", tree.Head.SearchBT(37), count)

	fmt.Println("Inorder")
	tree.Head.PrintInOrder()
	fmt.Println("PreOrder")
	tree.Head.PrintPreOrder()
	fmt.Println("PostOrder")
	tree.Head.PrintPostOrder()
}
