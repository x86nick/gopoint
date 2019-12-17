package main

import (
	"fmt"
)

func main() {
	// Simple
	fmt.Println("Hello World")

	// Multiple
	fmt.Println("Hello World", "Blee", "Duh")

	fmt.Println()
	// Printing a string
	s := "Hello World!"
	fmt.Printf("%s\n", s)
	fmt.Printf("%s\n", s)
	// Right justified
	fmt.Printf("%20s Yo\n", s)
	// Left justified
	fmt.Printf("%-20s Yo\n", s)

	fmt.Println()
	// Printing a number
	i := 20
	fmt.Printf("%d\n", i)
	fmt.Printf("%10d %-20s\n", i, "Gold")
	fmt.Printf("%10d %-20s\n", 5, "Silver")
	fmt.Printf("%10d %-20s\n", 2, "Bronze")

	fmt.Println()
	// Printing a float
	f := 124.456789
	fmt.Printf("%f - %6.2f - %#v\n", f, f, f)

	fmt.Println()
	// Creating formatted strings
	s = fmt.Sprintf("I've watched %d commercials during the %s ;(\n", 1000, "olympics")
	fmt.Println(s)
}
