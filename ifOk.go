package main

import "fmt"

func fred() (int, bool) { return 10, true } // boolean allow to check if the first value is okay or not. it allows you to check the value

func main() {
	a := 1
	// not a good way, better us switch statement
	if a > 1 && a < 5 {
		fmt.Println("a > 1", a)
	} else if a == 1 {
		fmt.Println("a == 1", a)
	} else {
		fmt.Println("Anything else", a)
	}

	v, ok := fred()
	if ok {
		fmt.Println("I am happy as a Hipo!", v)
	} else {
		fmt.Println("Boom!")
	}

	// Or... it is better //
	if z, ok := fred(); ok { // ok, is a boolean, it is an ideam, single line if.. comma, ok ideam is common in go
		fmt.Println("I am happy as a Hipo!", z)
	} else {
		fmt.Println("Boom!")
	}
}
