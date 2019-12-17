package main

import (
	"fmt"
)

func main() {
	var i int
	defer fmt.Println("Booya!", func() int { i *= 5; return i }()) // fmt.P.. is argument call, it must be evaulesd at time it see defer, and evalute later.., func() must be called 
	// , at that time i = 0 
	
	i++	
	fmt.Println("i", i)
}
