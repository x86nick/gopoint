package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go func(c <-chan int) { // receiver 
		fmt.Println("Formatting images...")
		// Drain...
		for v := range c { // no indexing here, i have no meaning here, this loop will process as long as data coming and channel not closed
			fmt.Printf("IMG-%-2d ", v)
			time.Sleep(500 * time.Millisecond)
		}
	}(c)

	for i := 0; i < 10; i++ { // prouducer
		c <- i
	}
	close(c)
}
