package main

import "fmt"

func main() {
	var s []int // create nil slice
	hydrate(s)  // call hydrate on that.
	// Ruroh!
	fmt.Println(s) // print the result
}

func hydrate(s []int) { // most of libray "takes in stuff", don't return. // this is unallocate slice
	for i := 0; i < 10; i++ {
		s = append(s, i) // i use apped but never return the append
	}
}
