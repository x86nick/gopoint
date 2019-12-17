package main

import (
	"fmt"
)

// ---------------------------------------
type car struct {
	year               int
	make, model, color string
}

func paint(c car, color string) car {
	c.color = color
	return c
}

// ---------------------------------------
type posCar car  // paralle type

func maaco(c posCar, color string) posCar { // take c,cllor and return posCar
	c.color = color
	return c
}

func main() {
	var (
		c car
		p posCar
	)
	// Doh! -- cannot use c (type car) as type posCar in assignment
//	p = c

	p = posCar(c) // type promotion
	fmt.Printf("%#v\n\n", p)

	barneys := posCar{
		year:  1971,
		make:  "Ford",
		model: "Pinto",
		color: "Puke Green",
	}
	garys := car{1975, "Chevrolet", "Impala", "Bronco Orange"}
	fmt.Println(paint(garys, "Gnarly Orange"))
	// Ruroh??
	fmt.Println(paint(barneys, "Budweiser Blue"))
	fmt.Println()
	// Ruroh??
	fmt.Println(maaco(garys, "Raging Pink"))
	fmt.Println(maaco(barneys, "Revolting Slob"))
}
