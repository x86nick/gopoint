// Write FizzBuzz game!
//   o FizzBuzz Game
//     o If the number is a multiple of 3, they say **Fizz**
//     o If the number is a multiple of 5, they say **Buzz**
//     o For multiple of both 3 and 5 they say **FizzBuzz**
//     o Otherwise they say the number.
//   o Write a program to compute FizzBuzz numbers up to 20
package main

import (
	"fmt"
	"strconv"
)

const (
	divBy3     = "Fizz"
	divBy5     = "Buzz"
	divBy3And5 = divBy3 + divBy5
)

func play(n int) string {
	switch {
	case n%3 == 0 && n%5 == 0:
		return divBy3And5
	case n%3 == 0:
		return divBy3
	case n%5 == 0:
		return divBy5
	default:
		return strconv.Itoa(n)
	}
}

func main() {
	for i := 1; i <= 20; i++ {
		fmt.Printf("%v ", play(i))
	}
}
