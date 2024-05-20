package main

import (
	"fmt"
	genericlist "generic-list/list/genericList"
)

func main() {
	j := 12
	var k *int
	k = &j
	fmt.Println(k)
	fmt.Println(*k)
	list := genericlist.New[string]()
	list.Insert("Bob1")
	list.Insert("Bob2")
	list.Insert("Bob3")
	list.Insert("Bob4")

	fmt.Printf("%+v \n", list)
	list.Remove(1)
	fmt.Printf("%+v \n", list)
	list.RemoveByValue("Bob4")
	fmt.Printf("%+v \n", list)

}
