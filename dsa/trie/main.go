package main

import "fmt"

// AlphabetSize is the number of possible characters
const AlphabetSize = 26

// Create struct for Node, Node represents each node in the trie
type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

// Create struct for trie
type Trie struct {
	root *Node
}

//  InitTrie creates a new Trie
func InitTrie() *Trie {
	return &Trie{
		root: &Node{},
	}
}

// Insert method

// Serach method

func main() {
	trie := InitTrie()
	fmt.Println(trie)
}
