package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s := make([]int, 2) // len=2, cap=2
	dump("s", s)
	// BOOM!
//	s[2] = 40 // panic: runtime error: index out of range [2] with length 2

	s = append(s, 30, 40)
	dump("s", s)

	s = append(s, 50)
	dump("s", s)

	s = append(s, 60)
	dump("s", s)

	s = append(s, 70, 80, 90)
	dump("s", s)
}

func dump(m string, s []int) {
	fmt.Printf("%s - %p %5v: %v\n", m, &s, *(*reflect.SliceHeader)(unsafe.Pointer(&s)), s)
}

/*
slice can keep doubling with every append... at one point double array putting pressure on GC 
slice keep doubling until 1024 if the capacity is not equal of lenght

s - 0x40a0e0 {4251680     2     2}: [0 0]
s - 0x40a100 {4251744     4     4}: [0 0 30 40]
s - 0x40a120 {4579360     5     8}: [0 0 30 40 50]
s - 0x40a140 {4579360     6     8}: [0 0 30 40 50 60]
s - 0x40a160 {4391104     9    16}: [0 0 30 40 50 60 70 80 90]
*/
