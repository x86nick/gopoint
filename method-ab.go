// Pointer Semantic
package main

import (
	"fmt"
	"time"
)

type (
	car struct {
		year         int
		maker, model string
		color        string
	}
	posCar car  // any behaviour will not pass to parallel type // posCar is not a car, you have to convert it
)

func (c *car) isNew() bool {
	return c.year == time.Now().Year()
}

func main() {
	c := posCar{
		year:  time.Now().Year(),
		maker: "Ford",
		model: "Pinto",
		color: "Puke Green",
	}
	// RuRoh?
	fmt.Println(c.isNew())
}
