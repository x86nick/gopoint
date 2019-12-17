package main

import (
	"fmt"
)

func main() {
  // Main use case
  defer fmt.Println("Yo Mama!")  // defer is a function, will be call when it will hit the bracket then it hit line 20 '}'
  // Function call  // use case, trying to open, you will always close the file, rather than closing at ever exit point, you just defer
  defer done() // defer must be a callable function

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
  defer func() { i++ }()
  return i
}

// Named returns. Side effect!!!
func fred() (i int) {
  i = 10
  defer func() { i++ }()
  return
}
