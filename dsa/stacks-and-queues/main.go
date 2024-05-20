package main

import "fmt"

// Stack represents stack that holds a slice
type Stack struct {
	items []int
}

func (s *Stack) push(val int) {
	s.items = append(s.items, val)
}
func (s *Stack) pop() int {
	length := len(s.items) - 1
	if length == -1 {
		return 0
	}
	val := s.items[length]
	s.items = s.items[:length]
	return val
}

// Queue represents a queue that holds a slice
type Queue struct {
	items []int
}

func (q *Queue) enqueue(val int) {
	q.items = append(q.items, val)
}
func (q *Queue) dequeue() int {
	if len(q.items) == 0 {
		return 0
	}
	elem := q.items[0]
	q.items = q.items[1:]
	return elem
}

func main() {
	s := &Stack{}
	s.push(1)
	s.push(2)
	s.push(3)
	s.push(4)
	fmt.Println(s.items)
	l := s.pop()
	fmt.Println(l)
	fmt.Println(s.items)

	q := &Queue{}
	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)
	q.enqueue(4)
	fmt.Println(q.items)
	f := q.dequeue()
	fmt.Println(f)
	fmt.Println(q.items)
}
