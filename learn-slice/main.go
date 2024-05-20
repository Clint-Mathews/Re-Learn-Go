package main

import (
	"fmt"
	"sort"
)

var users = []string{}

func addUsers(users []string) {
	for _, user := range users {
		fmt.Println("user", user)
	}
}

func addUsers1(users ...string) {
	for _, user := range users {
		fmt.Println("user", user)
	}
}

func addUsers2(user string) {
	users = append(users, user)
}

func addUsers3(userList ...string) {
	users = append(users, userList...)
}

func removeFromSliceUnOrdered(numbers []int, index int) []int {
	numbers[index] = numbers[len(numbers)-1]
	numbers = numbers[:len(numbers)-1]
	return numbers
}

func removeFromSliceOrdered(numbers []int, index int) []int {
	numbers = append(numbers[:index], numbers[index+1:]...)
	return numbers
}

type Numbers []int

func (n Numbers) Remove(index int) []int {
	return append(n[:index], n[index+1:]...)
}

func (n Numbers) Len() int {
	return len(n)
}
func (n Numbers) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}
func (n Numbers) Less(i, j int) bool {
	return n[i] < n[j]
}

// func (n Numbers) Sort() []int {
// }

func main() {
	// data := []string{"Clint", "Clins"}
	fmt.Println("addUsers")
	addUsers([]string{"Clint", "Clins"})
	fmt.Println("addUsers1")
	addUsers1("Clint", "Clins")

	fmt.Println("addUsers2", users)
	addUsers2("Clint")
	addUsers2("Clins")
	fmt.Println("addUsers2", users)

	fmt.Println("addUsers3", users)
	addUsers3("Clint1", "Clins2")
	fmt.Println("addUsers3", users)

	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("removeFromSliceUnOrdered numbers", numbers)
	numbers = removeFromSliceUnOrdered(numbers, 1)
	fmt.Println("removeFromSliceUnOrdered numbers", numbers)

	numbers = []int{1, 2, 3, 4, 5}
	fmt.Println("removeFromSliceOrdered numbers", numbers)
	numbers = removeFromSliceOrdered(numbers, 1)
	fmt.Println("removeFromSliceOrdered numbers", numbers)

	num := Numbers{2, 3, 4, 6}
	fmt.Println(num)
	num = num.Remove(1)
	fmt.Println(num)

	nums := Numbers{1, 45, 2, 9, 23}
	fmt.Println(nums)
	sort.Sort(nums)
	fmt.Println(nums)

}
