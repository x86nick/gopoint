package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("YO!")
	
	panic("Something bad happened!")
}

// if you write library code, if you have to painc, it is not good practice, painc ment to be at top level of program, 
//we can not pass this decision to a user.
// painc in a library is an error...
// anything below main(), don't make painc 
