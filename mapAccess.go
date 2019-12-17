package main

import "fmt"

func main() {
	m := map[string]int{}

	// Inserting...
	m["yo"] = 10
	fmt.Printf("m     = %#v\n", m)
	fmt.Printf("m[yo] = %#v\n", m["yo"])
	fmt.Printf("m[ye] = %#v\n", m["ye"]) // NOTE!!

	// Member?
	if v, ok := m["boom"]; ok {  // scope of v is here
		fmt.Println(v, ok)
	} else {
		fmt.Println("Not found!")
	}

	// Update...
	if v, ok := m["yo"]; ok {
		m["yo"] = v + 1
	}
	fmt.Printf("m = %#v\n", m)
}

/*

every time you access a map, don't get lazy and use v, ok idem 
v, ok = m['yo']   

True -- v is good
False -- v is not good

if v is  not in map, it will return zero value...

if key is not in map, you get zero value, if no one set, you will get false value

*/
