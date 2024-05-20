package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	root   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	// Add element infront of the root node
	n.next = l.root
	l.root = n
	l.length++
}

func (l *linkedList) deleteNode(val int) {
	// 1. Delete node from front
	// 2. Delete node from middle
	// 3. Delete node from last
	if l.root == nil {
		fmt.Println("Linked list empty!")
		return
	}

	// Delete from front
	if l.root.data == val {
		l.root = l.root.next
		l.length--
		return
	}

	// Check if there is single node
	if l.root.next == nil {
		fmt.Println("Node not found!")
		return
	}

	cur := l.root
	for cur.next.data != val {
		if cur.next.next == nil {
			fmt.Println("Node not found!")
			return
		}
		cur = cur.next
	}

	// Delete from end
	if cur.next.next == nil {
		cur.next = nil
		l.length--
		return
	}

	cur.next = cur.next.next
	l.length--
}

func (l linkedList) print() {
	n := l.root
	fmt.Println("Print linked list:")
	fmt.Println("Length of linkedList: ", l.length)
	for n != nil {
		fmt.Println("node: ", n.data)
		n = n.next
	}
}

func main() {
	l := &linkedList{}
	n1 := &node{
		data: 122,
	}
	// n2 := &node{
	// 	data: 21,
	// }
	// n3 := &node{
	// 	data: 221,
	// }
	// n4 := &node{
	// 	data: 28,
	// }
	// n5 := &node{
	// 	data: 3,
	// }
	l.prepend(n1)
	// l.prepend(n2)
	// l.prepend(n3)
	// l.prepend(n4)
	// l.prepend(n5)
	l.print()
	fmt.Println(l)
	l.deleteNode(2211)
	l.print()
	fmt.Println(l)
}
