package main

import (
	"fmt"
	"time"
)

func main() {
	f := func(i int) int {
		time.Sleep(time.Duration(i) * time.Millisecond)
		return i
	}

	timeMe(10, f) // take interger, and take a function as argument
}

func timeMe(i int, f func(int) int) int {  //  this is better way, poor mans timer
	defer func(t time.Time) {  // defer must be a funcation, it is, takes time stamp 
		fmt.Printf("Dude! You Took %v to run...\n", time.Since(t))
	}(time.Now()) // you start the clock here..., then make  call to function
	return f(i)
} // the anonomsuys function is called here, defer stack is unwinded
