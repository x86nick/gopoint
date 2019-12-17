package main

import "fmt"

type car struct {
	maker, color string
}

func main() {
	s := []car{
		{maker: "Ford", color: "Red"},
		{maker: "Toyota", color: "Blue"},
	}

	for _, v := range s {
		v.color = "Cool " + v.color  // changing only the copy, local to scope..how to fix this ? use index like line 21
	}
	// Ruroh?...
	fmt.Printf("%#v\n", s)

	for i := range s {
	  s[i].color = "Cool " + s[i].color
	}
	fmt.Printf("%#v\n", s)
}
