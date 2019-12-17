// Pointer Semantic
package main

import (
	"fmt"
	"time"
)

type car struct {
	year        int
	maker, model string
	color       string
}

func (c *car) isNew() bool {  // copy of pointer to car // it is ment to be shared, it is essenc of the car.. if esence is not to be shared then use the 'value semantic ' func (c car) isNew() bool
}
func (c *car) paint(color string) {  // it mutate, its pointer //  if you want to mutate without *, then use return statement and no *
	c.color = color
}

func main() {
	fredsCar := car{
		year:  time.Now().Year(),
		maker:  "Ford",
		model: "Pinto",
		color: "Puke Green",
	}
	fmt.Println(fredsCar)
	fmt.Println("New?", fredsCar.isNew())
    fredsCar.paint("Navajo White")
    fmt.Println(fredsCar)
}
