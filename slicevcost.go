package main

import "fmt"

func main() {
	// Iteration
	s := []int{0: 1, 3: 3} // len 4 | cap 4

	// This is ok
	// NOTE: More control then range since always iterate from start of slice
	for i := 1; i < len(s); i += 2 {
		fmt.Printf("%d ", s[i])
	}
	fmt.Println()

	// Or...
	// !!!! NOTE: Range makes a copy of each element in the slice!
	for i, v := range s {
		fmt.Printf("[%d] %p -> %p(%d)\n", i, &v, &s[i], v)
	}
	fmt.Println()

	// Or skip index
	for _, v := range s {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	// Or skip value
	for i := range s {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// Slice as arguments
	s1 := make([]int, 2e6)
	process(s1)
	fmt.Println(s1[0])
}

// Always passing a constant 3 words
func process(s []int) {  // much better even though it has cost of 3 words being copied  // most of the time this is that we want to see. don't use *[]int here
    s[0] = 100
}

