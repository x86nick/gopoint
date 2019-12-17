package main

import (
	"fmt"
)

func main() {
	f := 10.5
	fmt.Printf("Main: %v -- %p\n", f, &f)
	inc(&f) //this is still a copy
	fmt.Printf("Main: %f -- %p\n", f, &f)
}

func inc(f *float64) {
	*f += 0.5
	fmt.Printf("Inc: %f -- %p-%p\n", *f, &f, f)
}

// every thing is pass by  value 
