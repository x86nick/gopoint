package main

import (
	"fmt"
)

func main() {
  // Main use case
  defer fmt.Println("Yo Mama!")
  // Function call
  defer done()

  // LIFO
  for i := 0; i < 5; i++ {
    defer fmt.Println(i)
  }

  fmt.Println("Blee", blee())
  fmt.Println("Fred", fred())
}

func done() {
  fmt.Println("That's all folks!")
}

// Can be an anonymous function. No side effect!
func blee() int {
  i := 10
  defer func() { i++ }()  // gets evaluate not run, thats why j is set 
  return i
} // this will call after function return

// Named returns. Side effect!!!=, try to avoid them
func fred() (i int) {
  i = 10
  defer func() { i++ }() // this will print 11
  return
}

// or  // this is better 
func fred() (int) { 
  i := 10
  defer func() { i++ }()  // this will print 10
  return i
}
