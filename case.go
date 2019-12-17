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

func play(n int) string {
	// YOUR_CODE...
	div3 := "Fizz"
	div5 := "Buzz"
	div3And5 := "FizzBuzz"

	switch {
	case n%3 == 0 && n%5 == 0:
		return div3And5
	case n%3 == 0:
		return div3
	case n%5 == 0:
		return div5

	default:
		return strconv.Itoa(n)
	}

}

func main() {
	// YOUR_CODE...
	for i := 1; i <= 20; i++ {
		fmt.Println(play(i))
	}
}
