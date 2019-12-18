package main

import "fmt"

type painter interface { // what is behaviour of your type
	paint(color string)
}

// Non Shared!
type car struct {
	brand, color string
}

func (c car) paint(color string) { // this is using value semantic
	c.color = color
}

// Shared!
type table struct {
	brand, color string
}

func (t *table) paint(color string) { // the reciever is a *table, for pointer semantic, pass the pointer
	t.color = color
}

func main() {
	// 	Note!
	var _ painter = car{}

	booth := []painter{ // bad, you are mixing value semantic, you have to use &table to fix this
		car{brand: "Ford", color: "Puke Green"},
		// table{brand: "Ikea", color: "Red"},
		&table{brand: "Ikea", color: "Red"},
	}

	for i := 0; i < len(booth); i++ {
		paintIt(booth[i], "Caribbean Green")
		fmt.Println(booth[i])
	}
}

func paintIt(p painter, color string) {
	p.paint(color)
}

/*
type -- method set
value  -- instance

don't mix the semantics

*/
