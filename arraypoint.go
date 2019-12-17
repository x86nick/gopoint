package main

import "fmt"

func main() {
	a := [5]*int{}
	b := [5]*int{0: new(int)}
	fmt.Printf("a(%p) = %#v\nb(%p) = %#v\n\n", &a, a, &b, b)

	a = b
	fmt.Printf("a(%p) = %#v\nb(%p) = %#v\n\n", &a, a, &b, b)

	*a[0] = 10
	*b[0] = 20
	fmt.Printf("a(%p) = %#v\nb(%p) = %#v\n\n", &a, *a[0], &b, *b[0]) // pointers are for sharing state
}
