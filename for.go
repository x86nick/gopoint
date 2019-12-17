package main

import (
	"fmt"
	"time"
)

func main() {
	s := [4]int{1, 2, 3, 4}

	// Traditional
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}

	// While
	var done bool
	for !done {
		fmt.Println("Sleepy...")
		time.After(1 * time.Second)
		done = true
	}

	// While
	for {
		fmt.Println("Not Forevar!")
		time.Sleep(1 * time.Second)
		break
	}

	i := 10
Blee:
	for {
	Fini:
		for {
			fmt.Println("Breakit down!")
			if i == 10 {
				time.Sleep(1 * time.Second)
				break Blee
			} else {
				break Fini
			}
		}
		break Blee
	}
	fmt.Println("BreakIt Done!")

	// Omit increment
	for done := false; done; {
		fmt.Println("No Inc!!")
		time.After(1 * time.Second)
		done = true
	}

	done = false
	// Omit init
	for ; !done; done = true {
		fmt.Println("No Init!!")
		time.After(1 * time.Second)
	}

	// Multi inits
	for i, j := 0, 10; i < 10; i, j = i+1, j+10 {
		fmt.Println(i, j)
	}
}
