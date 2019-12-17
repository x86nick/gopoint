package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s := []int{1, 2, 3}
	dump("s", s)

	s1 := append(s[:1], s[2:]...) // go grab org upto index 1 , pass the index that you want to delete upto end 
	// len =1 , cap = 3, 
	dump("Delete 2", s1)
	dump("s", s)

	// Delete first
	s2 := s1[1:]
	dump("Delete 1", s2)

	// Delete last
	s3 := s2[:len(s2)-1]
	dump("Delete 3", s3)
}

func dump(m string, s []int) {
	fmt.Printf("%-10s - %p %5v: %v\n", m, &s, *(*reflect.SliceHeader)(unsafe.Pointer(&s)), s)
}



/*
append looks - does len eq to capacity 
append work same way for adding and deleting
s          - 0x40a0e0 {4251680     3     3}: [1 2 3]
Delete 2   - 0x40a100 {4251680     2     3}: [1 3]
s          - 0x40a120 {4251680     3     3}: [1 3 3]
Delete 1   - 0x40a140 {4251684     1     2}: [3]
Delete 3   - 0x40a160 {4251684     0     2}: []

*/
