package main

import "fmt"

type car struct {
	maker, color string
}

func main() {
	s := []car{
		{maker: "Ford", color: "Red"},
	}
	s[0].maker = "Chevy"  // go set maker to Chevy , chevy red
	fmt.Println(s)

	m := map[string]car{ // key is string and value is car
		"fred": {maker: "Ford", color: "Red"},
	}

	// Tho...
	m["blee"] = car{maker: "Porsche", color: "Green"}

	// But Ruroh!
	//m["fred"].maker = "Chevy" // cannot assign to struct field m["fred"].maker in map ||| not allow to mutate the sturct, to do that, first pull out of map, update and put it back
	// or use pointer semantic if not want to mutate 
	fmt.Println(m)

	// Replace
	e := m["fred"]
	e.maker = "Chevy"
	m["fred"] = e
	fmt.Println(m)

	// Or...
	m1 := map[string]*car{
		"fred": &car{maker: "Ford", color: "Red"},
	}
	m1["fred"].maker = "Chevy"
	fmt.Println(m1["fred"].maker)
}
