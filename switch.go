package main

import (
	"fmt"
)

func main() {
	i, j := 10, 20

	switch i {
	case 10:
		fmt.Println("It is 10")
	case 20:
		fmt.Println("It is 20")
	default:
		fmt.Println("It is, what it is", i)
	}

	// Expression less...
Fini:	
	switch {
	case j == 10:
		fmt.Println("It is 10")
		break
	case j == 20:
		fmt.Println("It is 20")
		break Fini
	}

	// Fallthrough...
	switch i {
	case 10:
		fallthrough
	case 20:
		fmt.Println("All good")
		// break
		// break is implicit in go 
	default:
		fmt.Println("No good")
	}

	// Or... merge cases
	switch i {
	case 10, 20: // this is better than fallthrough
		fmt.Println("All good")
	default:
		fmt.Println("Not good")
	}
}
