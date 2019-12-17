// In go we can create an variable that can hold address of another variable, this variable called pointer variable
// Pointer variable contain address of memory location that hold data of a given type
// Pointer indirection (*p) gives the value stored at this address
// zero value of a pointer of any type is nil
// attempting to acces the value of a nil pointer cause runtime panic
// pointer can be compared to each other and to nil
// Pointers allow sharing data across functions without excessive copying.
// & (an ampersand), Go’s “address of” operator (&amount - Retrieve the varibles address)
// *int - pointer to int
// so *myIntPointer is “value at myIntPointer.”
package main

import "fmt"

func main() {

	var a int = 1
	var p *int // variable p of type "pointer to int" // *int = "pointer to int" // pointer type zero value is nil
	p = &a     // assign the "address of variable a" to p or ( p holds the address of a)// &a = address of a
	fmt.Println("p's values is a's address:", p)
	fmt.Println("*p yields a's value:", *p)

	/*  literal values and contatns are not addressable, you will get error */

	// p1 := &123 // ./p1.go:15:8: cannot take the address of 123
	// const c = 123
	// p1 = &c // ./p1.go:17:7: cannot take the address of c

	/* Retriving a value through pointer */

	// *p = pointer indirection

	// var p2 *int
	// fmt.Println(*p2) // panic: runtime error: invalid memory address or nil pointer dereference

	/* Built in function new() - creates a variable and returns a pointer */

	p3 := new(int)
	fmt.Println("p points to an unnamed int:", p3, *p3)

	*p = 6
	fmt.Println("*p is now: ", *p) // *p is now:  6

}
