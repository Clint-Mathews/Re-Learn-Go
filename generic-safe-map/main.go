package main

import (
	"fmt"

	safemap "github.com/Clint-Mathews/GenericSafeMap/safe-map"
)

func main() {
	m := safemap.New[int, int]()
	// m.Insert(1, 1)
	// m.Insert(2, 2)
	// m.Insert(3, 3)
	// m.Insert(4, 4)

	// fmt.Printf("%d exists - %t \n", 1, m.HasKey(1))
	// m.Update(1, 10)
	// val, _ := m.Get(1)
	// fmt.Printf("%d value - %v \n", 1, val)

	for i := 1; i < 5; i++ {
		go func(i int) {
			m.Insert(i, i+10)
			val, err := m.Get(i)
			if err != nil {
				fmt.Printf("Error on %d - %+v \n", i, err)
			}

			if val != i+10 {
				fmt.Println("Update did not work!")
			}

		}(i)
	}
}
