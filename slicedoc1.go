package main

import "fmt"

func main() {
	// Note! Use this to allocate storage. Init to zero value of type!
	s0 := make([]int, 3)      // len=3, cap=3
	s1 := make([]int, 3, 5)   // len=3, cap=5
	// Note! Use this when have values to init with
	s2 := []int{10, 20}       // len=2, cap=2
	s3 := []int{0: 10, 9: 20} // len=10, cap=10
    s4 := []int{}             // Empty Slice1 len=cap=0
    var s5 []int              // Nil Slice
    
	dump("s0", s0)
	dump("s1", s1)
	dump("s2", s2)
	dump("s3", s3)
	dump("s4", s4)
	dump("s5", s5)
}

func dump(msg string, s []int) {
    if s == nil {
        fmt.Println("Hoy! Nil Slice detected...")
    }
	fmt.Printf("(%-2s) -- [%p] %v -- %d|%d\n", msg, &s, s, len(s), cap(s))
}

/*
The Go Playground    Imports  
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
22
23
24
25
26
27
28
29
30
31
32
33
34
35
36
package main

import "fmt"

func main() {
	// Note! Use this to allocate storage. Init to zero value of type!
	s0 := make([]int, 3)      // len=3, cap=3
	s1 := make([]int, 3, 5)   // len=3, cap=5
	// Note! Use this when have values to init with
	s2 := []int{10, 20}       // len=2, cap=2
	s3 := []int{0: 10, 9: 20} // len=10, cap=10
    s4 := []int{}             // Empty Slice1 len=cap=0
    var s5 []int              // Nil Slice
    
	dump("s0", s0)
	dump("s1", s1)
	dump("s2", s2)
	dump("s3", s3)
	dump("s4", s4)
	dump("s5", s5)
}

func dump(msg string, s []int) {
    if s == nil {
        fmt.Println("Hoy! Nil Slice detected...")
    }
	fmt.Printf("(%-2s) -- [%p] %v -- %d|%d\n", msg, &s, s, len(s), cap(s))
}

/*
(s0) -- [0x40a0e0] [0 0 0] -- 3|3
(s1) -- [0x40a100] [0 0 0] -- 3|5
(s2) -- [0x40a120] [10 20] -- 2|2
(s3) -- [0x40a140] [10 0 0 0 0 0 0 0 0 20] -- 10|10
(s4) -- [0x40a160] [] -- 0|0
Hoy! Nil Slice detected...
(s5) -- [0x40a180] [] -- 0|0
*/
