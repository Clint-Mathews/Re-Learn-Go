package main

import (
	"fmt"
	"sync"
	"time"
)

type Clint []int

func main() {
	start := time.Now()
	fmt.Println("Goroutine aggregate data!")

	wg := &sync.WaitGroup{}
	// Buffered channel would only work as we dont read data from channel untill every single function completes its execution
	// Unbuffered would accept data untill anything added is read
	// Buffered with 3 or more to avoid deadlock
	resch := make(chan any, 3)

	wg.Add(3)
	go getUsername(resch, wg)
	go getUserLikes(resch, wg)
	go getUserDOB(resch, wg)

	wg.Wait()
	close(resch)

	for i := range resch {
		if res, ok := i.(int); ok {
			fmt.Println("Likes: ", res)
		}
		if res, ok := i.(string); ok {
			fmt.Println("Username: ", res)
		}
		if res, ok := i.(Clint); ok {
			fmt.Println("Data Clint: ", res)
		}

	}

	fmt.Println("Took: ", time.Since(start))
}

func getUsername(resch chan<- any, wg *sync.WaitGroup) {
	time.Sleep(100 * time.Millisecond)
	resch <- "Clint"
	wg.Done()
}

func getUserDOB(resch chan<- any, wg *sync.WaitGroup) {
	time.Sleep(300 * time.Millisecond)
	resch <- Clint{1, 2, 3, 4}
	wg.Done()
}

func getUserLikes(resch chan<- any, wg *sync.WaitGroup) {
	time.Sleep(200 * time.Millisecond)
	resch <- 10
	wg.Done()
}
