package main

import "fmt"

func main() {
	m := map[string]int{"fred": 1, "blee": 2, "duh": 3}

	// Deleting...
	fmt.Printf("Before %#v\n", m)
	if _, ok := m["fred"]; ok {
		delete(m, "fred")
	}
	fmt.Printf("After: %#v\n", m)
}
