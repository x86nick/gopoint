// Value Semantic
package main

import (
	"fmt"
	"time"
)

type car struct {
	year         int
	maker, model string
	color        string
}

// func isNew(cCar) bool

func (c car) isNew() bool { // this is method :
	return c.year == time.Now().Year() // i need c here to call c.year...thats why i need method
}
func (c car) paint(color string) { // receiver in go // but it does not returning anyting 
	c.color = color // will take the copy of the car...not mutated value
}

func main() {
	fredsCar := car{
		year:  time.Now().Year(),
		maker: "Ford",
		model: "Pinto",
		color: "Puke Green",
	}
	fmt.Println(fredsCar.isNew())
	fmt.Println(fredsCar)
	// Ruroh?
	fredsCar.paint("Bronco Orange") // paint function has no 'return', the refered object (fredsCar) won't get changed.
	fmt.Println(fredsCar)
}
