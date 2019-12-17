//used a lot of json marshall and unmarshall
package main

import (
	"fmt"
)

type car struct {
	year                int
	maker, model, color string
}

func main() {
	knockOff := struct { // same shape and same order type, compiler will be relexed, not used frequently, only time it is used...
		year                int
		maker, model, color string
	}{
		year:  2017,
		maker: "Ford",
		model: "Pinto",
		color: "Blue",
	}

	fmt.Printf("%#v\n", paint(knockOff, "Gnarly Orange")) // calling paint
}

func paint(c car, color string) car { // call to paint that returns a car
	c.color = color
	return c
}
