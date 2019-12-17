package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5}

	// NOTE! Ensure new slice gets its own copy
	s1 := append([]int{}, s[2:]...)
	s1[0] = 100
	fmt.Println(s, s1)
}
