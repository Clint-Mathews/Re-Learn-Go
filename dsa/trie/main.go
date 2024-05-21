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

// Insert method will add a word into the trie
func (t *Trie) Insert(s string) {
	currNode := t.root
	for _, v := range s {
		charIndex := v - 'a' // Interseting part where we find index by subtracting ASCII value of 'a' from the current string
		if currNode.children[charIndex] == nil {
			currNode.children[charIndex] = &Node{}
		}
		currNode = currNode.children[charIndex]
	}
	currNode.isEnd = true
}

// Serach method will take a word and return true if it s found in Trie
func (t *Trie) Search(s string) bool {
	currNode := t.root
	for _, v := range s {
		charIndex := v - 'a'
		if currNode.children[charIndex] == nil {
			return false
		}
		currNode = currNode.children[charIndex]
	}
	return currNode.isEnd
}

func main() {
	trie := InitTrie()

	toAdd := []string{
		"test", "kkkksd", "popiko",
	}

	for _, v := range toAdd {
		trie.Insert(v)
	}

	fmt.Println("popiko: ", trie.Search("popiko"))
	fmt.Println("zclint: ", trie.Search("zclint"))

}
