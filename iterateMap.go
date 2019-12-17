
// How to access key in map as order is not guranteed

package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[string]int{"fred": 1, "blee": 2, "duh": 3}

	// Iterating
	for k, v := range m {
		fmt.Printf("%-4s -> %v\n", k, v)
	}

	// Sorted access...
	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}
	
	sort.Strings(keys)
	fmt.Printf("\nSorted\n")
	for _, k := range keys {
	    fmt.Printf("%-4s -> %d\n", k, m[k])
	}
}
