package main

import "fmt"

func main() {
	var a [5]int
	fmt.Printf("a = %#v\n", a)
	var b [2]bool
	fmt.Printf("b = %#v\n", b)

	// Shorthand...
	c := [3]float64{}
	fmt.Printf("c = %#v\n", c)

	d := [3]float64{100, 200, 300}
	fmt.Printf("d = %#v\n", d)

	// Index assigns...
	e := [3]float64{0: 100.5, 2: 1000.5}
	fmt.Printf("e = %#v\n", e)

	f := [...]float64{0: 100.0, 4: 1000.0} // compiler will fill the number 
	fmt.Printf("f = %v - len = %d / cap = %d\n\n", f, len(f), cap(f))

	// Ruroh!
// 	fmt.Println(f[6])
}
