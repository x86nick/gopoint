package main

import "fmt"

type multiplier func(i int) int // this is multiplier function

func fred(i int, m multiplier) int { // fred takes int and multiplier, allows to do just that..
	return m(i)
}

func thousands(i int) int {
	return i * 10_000
}

func main() {
	k := fred(10, func(i int) int {
		return i * 100
	})
	fmt.Println("Calling hundreds", k)

	t := fred(10, thousands)
	fmt.Println("Calling thousands", t)
}
// way to name usertypes in go
