package main

import (
	"fmt"
)

func main() {
	caller()
	fmt.Println("This program always works!")
}

func caller() {
	defer func() {
		if val := recover(); val != nil { // when we recover, if not nil, then reover 
			fmt.Printf("Recovering from `%T %v`\n", val, val)
		}
	}()

	killMe() // cause painc
}

func killMe() {
	defer fmt.Println("KillMe completed...")
	panic("And we died!")
}
// if anybody below me did painc, then i need to be able to trigger recover()

//KillMe completed...
//Recovering from `string And we died!`
//This program always works!
