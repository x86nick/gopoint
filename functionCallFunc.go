package main

import (
	"fmt"
)

func main() {
	loveMe := func(i int) func() string {
		return func() string { return fmt.Sprintf("Love me `%d times, I'm gone away...", i) }
	}
	fmt.Println(loveMe(2)())
	fmt.Println(loveMe(3)())
}
