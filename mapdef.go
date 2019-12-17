package main

import "fmt"

func main() {
	// Constructing...
	m1 := map[string]int{}
	m2 := make(map[string]int)
	m3 := make(map[int]bool, 10)
	m4 := map[string]int{"fred": 1, "blee": 2, "duh": 3}

	fmt.Printf("m1 = %#v\n", m1)
	fmt.Printf("m2 = %#v\n", m2)
	fmt.Printf("m3 = %#v\n", m3)
	fmt.Printf("m4 = %#v\n", m4)
}
