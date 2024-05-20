package main

import (
	"fmt"
)

type Server struct {
	users  map[string]string
	userch chan string
	quitch chan struct{}
}

func NewServer() *Server {
	return &Server{
		users:  make(map[string]string),
		userch: make(chan string),
		quitch: make(chan struct{}),
	}
}

func (s *Server) start() {
	go s.loop()
}

func (s *Server) loop() {
loop:
	for {
		select {
		case msg := <-s.userch:
			fmt.Println("msg", msg)
		case <-s.quitch:
			fmt.Println("Quiting server")
			break loop
		}
		// user := <-s.userch
		// s.users[user] = user
		// fmt.Printf("Adding new user: %s \n", user)
	}
}

// func main() {
// server := NewServer()
// server.start()

// if user, ok := <-server.userch; ok {
// 	fmt.Println(user)
// }

// // for i := 0; i < 10; i++ {
// // 	go func(i int) {
// // 		server.userch <- fmt.Sprintf("User - %d", i)
// // 	}(i)
// // }

// go func() {
// 	// server.quitch <- struct{}{}
// 	close(server.quitch)
// }()

// 	// select {}

// 	userch := make(chan string, 1)

// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		// close(userch)
// 		userch <- "Clint"
// 	}()

// 	for {
// 		select {
// 		case uname, ok := <-userch:
// 			fmt.Println(uname)
// 			fmt.Println(ok)
// 		}
// 	}

// 	username, ok := <-userch

// 	if !ok {
// 		fmt.Println("User not found!")
// 	}

// 	fmt.Println(username)

// }

func (s *Server) addUser(user string) {

	s.users[user] = user
}

func CloseTest1(ch <-chan struct{}) {
	for {
		chanClose := <-ch
		fmt.Println("chanClose1", chanClose)
	}
}

func CloseTest2(ch <-chan struct{}) {
	for {
		chanClose := <-ch
		fmt.Println("chanClose2", chanClose)
	}
}

// func main() {
// 	fmt.Println("Hello")

// 	userch := make(chan string, 1)
// 	// fmt.Println(fmt.Sprintf("%+v", u))
// 	// go func() {
// 	// time.Sleep(1 * time.Second)
// 	// userch <- "Clint"
// 	sendMessage(userch)
// 	readMessage(userch)
// 	// userch <- "Clins"
// 	// }()
// 	// u <- "Clint1"
// 	// go func() {
// 	// 	time.Sleep(1 * time.Second)
// 	// loop:
// 	// 	for {
// 	// 		select {
// 	// 		case <-userch:
// 	// 			fmt.Println("Channel clsoe")
// 	// 			break loop
// 	// 		}
// 	// 	}
// 	// 	fmt.Println("Hasasa")
// 	// }()
// 	// k := <-u
// 	// p := <-u
// 	// close(userch)
// 	// user := <-userch // Blocking awaiting
// 	// fmt.Println(user)
// 	// user = <-userch // Blocking awaiting
// 	// fmt.Println(user)
// 	// fmt.Println(p)
// 	// time.Sleep(2 * time.Second)
// 	chch := make(chan struct{})
// 	go CloseTest1(chch)
// 	go CloseTest2(chch)
// 	close(chch)
// }

func sendMessage(msgch chan<- string) {
	msgch <- "Haha"
}

func readMessage(msgch <-chan string) {
	msg := <-msgch
	fmt.Println(msg)
}

func main() {
	userch := make(chan string, 10000)
	for i := 0; i < 10; i++ {
		go func(i int) {
			userch <- fmt.Sprintf("User - %d", i)
		}(i)
	}
	// close(userch)
	for user := range userch {
		fmt.Println(user)
	}
}
