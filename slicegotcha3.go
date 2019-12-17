package main

import "fmt"

func main() {
	s := make([]int, 5)
	fmt.Println(s)
	load(s)
	fmt.Println(s)
}

func load(s []int) {
	copy(s, []int{1, 2, 3, 4, 5, 6, 7})
}

func load1(s []int) int {
	return copy(s, []int{1, 2, 3, 4, 5, 6, 7})
}
