package main

import (
	"sync"
	"time"

	sendMessageViaChannel "github.com/Clint-Mathews/Concurrency/send-message-via-channel"
)

func main() {
	// now := time.Now()
	// userId := 10
	// responseChan := make(chan string, 100)

	// wg := &sync.WaitGroup{}

	// go fetchUserData(userId, responseChan, wg)
	// go fetchUserRecommendations(userId, responseChan, wg)
	// go fetchUserLikes(userId, responseChan, wg)
	// wg.Add(3)

	// wg.Wait()
	// close(responseChan)

	// for data := range responseChan {
	// 	fmt.Println(data)
	// }
	// fmt.Println(time.Since(now))

	sendMessageViaChannel.MainFunc()
}

func fetchUserData(userId int, responseChan chan<- string, wg *sync.WaitGroup) {
	time.Sleep(60 * time.Millisecond)
	responseChan <- "user data"
	wg.Done()
}

func fetchUserRecommendations(userId int, responseChan chan<- string, wg *sync.WaitGroup) {
	time.Sleep(120 * time.Millisecond)
	responseChan <- "user recommendations"
	wg.Done()
	// close(responseChan)
}

func fetchUserLikes(userId int, responseChan chan<- string, wg *sync.WaitGroup) {
	time.Sleep(30 * time.Millisecond)
	responseChan <- "user likes"
	wg.Done()
}
