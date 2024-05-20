package main

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestAddUser(t *testing.T) {
	b := struct{}{}
	fmt.Println(unsafe.Sizeof(b))

	// server := NewServer()
	// server.start()
	//
	//	for i := 0; i < 10; i++ {
	//		// go server.addUser(fmt.Sprintf("user - %d", i))
	//		// go func() {
	//		server.userch <- fmt.Sprintf("user - %d", i)
	//		// }(i)
	//	}
	//
	// fmt.Println("Completed adding users!")
}
