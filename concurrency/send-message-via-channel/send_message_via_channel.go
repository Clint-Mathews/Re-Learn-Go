package sendMessageViaChannel

import (
	"fmt"
	"time"
)

type Message struct {
	From    string
	Paylaod string
}

type Server struct {
	msgch  chan Message
	quitch chan struct{}
}

func (s *Server) StartAndListen() {
loop:
	for {
		select {
		case msg := <-s.msgch:
			fmt.Printf("Message received: %s From %s\n", msg.Paylaod, msg.From)
		case <-s.quitch:
			fmt.Printf("Server shutting doen gracefully! \n")
			break loop
		default:
			// time.Sleep(1 * time.Second)
		}
	}
	fmt.Println("Server shutdown!")
}

func sendMessageToServer(msgch chan<- Message, msg Message) {
	msgch <- msg
	fmt.Println("Sending message to server!")
}

func (s *Server) quitServer() {
	close(s.quitch)
}

func MainFunc() {
	s := &Server{
		msgch:  make(chan Message),
		quitch: make(chan struct{}),
	}
	go s.StartAndListen()

	sendMessageToServer(s.msgch, Message{From: "Clint", Paylaod: "Test"})
	for i := 0; i < 10; i++ {
		go func(i int) {
			time.Sleep(1 * time.Second)
			sendMessageToServer(s.msgch, Message{From: fmt.Sprintf("Clins - %d", i), Paylaod: "Test"})
		}(i)
	}
	go func() {
		time.Sleep(2 * time.Second)
		s.quitServer()
	}()
	select {}
}
