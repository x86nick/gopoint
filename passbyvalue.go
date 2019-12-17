package main

import (
	"fmt"
)

func main() {
	f := 10.5
	fmt.Printf("Main: %f -- %p\n", f, &f)
	inc(f)
	fmt.Printf("Main: %f -- %p\n", f, &f)
}

func inc(f float64) {
	f += 0.5
	fmt.Printf("Inc: %f -- %p\n", f, &f)
}

// in go everything is pass by value 
