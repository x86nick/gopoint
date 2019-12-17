//  write a program to compute 3x3 matrix addition
// and multiplication using slices.
//   How is this different than using arrays?
package main

import "fmt"

// add 2 matrices of size 3x3
func add(a, b [][]int) [][]int {
	res := make([][]int, len(a))
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			res[i] = append(res[i], a[i][j] + b[i][j])
		}
	}
	return res
}

// multiply 2 matrices of size 3x3
func multiply(a, b [][]int) [][]int {
	res := make([][]int, len(a))
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
		    res[i] = append(res[i], 0)
			for k := 0; k < len(a[i]); k++ {
				res[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return res
}

func main() {
	a1 := [][]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}}
	a2 := [][]int{{3, 3, 3}, {4, 4, 4}, {5, 5, 5}}

	dump(add(a1, a2))
	dump(multiply(a1, a2))
}

func dump(a [][]int) {
    for i := 0; i < len(a); i++ {
        for j := 0; j < len(a[i]); j++ {
            fmt.Printf("%2d ", a[i][j])
        }
        fmt.Println()
    }
    fmt.Println()
}

