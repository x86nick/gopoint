package main

import (
	"fmt"
)

func main() {
	d := [3]float64{100, 200, 300}

	// Slicing...
	fmt.Printf("d[1:3] = %v\n", d[1:3])  // ranges, don't include last
	fmt.Printf("d[1:] = %v\n", d[1:])
	fmt.Printf("d[:2] = %v\n", d[:2])
	fmt.Printf("d[:] = %v\n %T--%T", d[:],d,d[:])
}
