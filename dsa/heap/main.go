package main

import "fmt"

// MaxHeap struct has a slice that holds the array

type MaxHeap struct {
	array []int
}

// Insert adds an element to the heap

func (m *MaxHeap) Insert(val int) {
	// Add element to array
	m.array = append(m.array, val)
	// Heapify up
	m.maxHeapifyUp(len(m.array) - 1)
}

// Extract returns the largest key, and removes it from the heap

func (m *MaxHeap) Extract() int {
	val := m.array[0]

	endIndex := len(m.array) - 1

	// Empty array
	if endIndex == -1 {
		fmt.Println("Empty array!")
		return endIndex
	}

	// Update last index as oarent index
	m.swap(0, endIndex)
	m.array = m.array[:endIndex]

	m.maxHeapifyDown(0)

	return val
}

// maxHeapifyUp will heapify from bottom to top
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[findparent(index)] < h.array[index] {
		// Swap the indexes
		h.swap(findparent(index), index)
		// Update index to the parent position to compare the above indexes
		index = findparent(index)
	}
}

// maxHeapifyDown will heapify from top to bottom
func (m *MaxHeap) maxHeapifyDown(index int) {
	// Larger child, compare to current index and swap
	lastIndex := len(m.array) - 1
	l, r := findLeftChild(index), findRightChild(index)
	childToCompare := 0

	// Loop while index has atleast one child
	for l <= lastIndex {
		// Define child to compare
		if l == lastIndex {
			// When left child is only child, to avoid panic
			childToCompare = l
		} else if m.array[l] > m.array[r] {
			// When left child is larger
			childToCompare = l
		} else {
			// When Right child is larger
			childToCompare = r
		}

		// Compare array value of current index to larger child and swap if smaller
		if m.array[index] < m.array[childToCompare] {
			m.swap(index, childToCompare)
			index = childToCompare
			l, r = findLeftChild(index), findRightChild(index)
		} else {
			// Found it right place!
			return
		}
	}

}

// get the parent index
func findparent(index int) int {
	return (index - 1) / 2
}

// get the left-right child index
func findLeftChild(index int) int {
	return index*2 + 1
}

func findRightChild(index int) int {
	return index*2 + 2
}

// Swap keys in an array
func (m *MaxHeap) swap(i, j int) {
	m.array[i], m.array[j] = m.array[j], m.array[i]
}

func main() {
	buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}
	heap := &MaxHeap{}
	for _, v := range buildHeap {
		heap.Insert(v)
		fmt.Println(heap)
	}
	for i := 0; i < 5; i++ {
		heap.Extract()
		fmt.Println(heap)
	}

}
