package main

import "fmt"

const ArraySize = 7

// HashTable Structure
type HashTable struct {
	array [ArraySize]*bucket
}

// Hash Function
func hash(key string) int {
	sum := 0
	for _, s := range key {
		sum += int(s)
	}
	return sum % ArraySize
}

// Insert will take in a key and add it in to the hash table
func (h *HashTable) Insert(key string) {
	// find index for insert using hash function
	index := hash(key)
	h.array[index].insert(key)
}

// Search will take in a key and search it in to the hash table
func (h *HashTable) Search(key string) bool {
	// find index for insert using hash function
	index := hash(key)
	return h.array[index].search(key)
}

// Delete will take in a key and delete it from the hash table
func (h *HashTable) Delete(key string) {
	// find index for insert using hash function
	index := hash(key)
	h.array[index].delete(key)
}

// bucket Structure
type bucket struct {
	head *bucketNode
}

// bucketNode Structure
type bucketNode struct {
	key  string
	next *bucketNode
}

// insert
func (b *bucket) insert(key string) {
	if b.search(key) {
		fmt.Println("Key already exists!")
		return
	}
	node := &bucketNode{key: key}
	node.next = b.head
	b.head = node
}

// search
func (b *bucket) search(key string) bool {
	cur := b.head
	for cur != nil {
		if cur.key == key {
			return true
		}
		cur = cur.next
	}
	return false
}

// delete
func (b *bucket) delete(key string) {

	if b.head.key == key {
		b.head = b.head.next
		return
	}

	prevNode := b.head
	for prevNode.next != nil {
		if prevNode.next.key == key {
			prevNode.next = prevNode.next.next
			return
		}
		prevNode = prevNode.next
	}
}

// Init will create a bucket in each slot of each hash table
func InitHashTable() *HashTable {
	result := HashTable{}
	for i := range ArraySize {
		result.array[i] = &bucket{}
	}
	return &result
}

func main() {
	hashTable := InitHashTable()
	list := []string{
		"ERRIC", "CLINT", "CLINS", "AKS", "AYY",
	}
	for _, v := range list {
		hashTable.Insert(v)
	}

	fmt.Println("CLINT exists: ", hashTable.Search("CLINT"))
	hashTable.Delete("CLINT")
	fmt.Println("CLINT exists: ", hashTable.Search("CLINT"))
}
