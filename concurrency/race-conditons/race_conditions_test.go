package raceconditons

import (
	"sync/atomic"
	"testing"
)

func TestDataRaceCondition(t *testing.T) {
	var state int32
	// var mu sync.RWMutex

	// for i := 0; i < 10; i++ {
	// 	go func(i int) {
	// 		mu.Lock()
	// 		state += int32(i)
	// 		mu.Unlock()
	// 	}(i)
	// }

	for i := 0; i < 10; i++ {
		go func(i int) {
			atomic.AddInt32(&state, int32(i))
		}(i)
	}
}

// type Server struct {
// 	isRunning atomic.Value
// }

// func TestDataRaceConditions(t *testing.T) {
// 	respCh := make(chan int)
// 	s := &Server{
// 		isRunning: atomic.Value{1},
// 	}
// }
