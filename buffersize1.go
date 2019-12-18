package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	c := make(chan string)
	// 	c := make(chan string, 1) , line 19 is not blocked in this case. go routine in main, will put buffer of size 1, buffer size 1 != unbuffered channel
 
	wg.Add(1)
	consumer(c)

	c <- "Yo!" // write to channel, it never get consumed hence blocked. [in this case process was done]
// line 14, this will always live in memory and GC will take care 
	wg.Wait()
	fmt.Println("All done!")
}

func consumer(c chan string) {
	defer func() {
		fmt.Println("Consumer bailed!")
		wg.Done()
	}()

	go func() {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond) // hangout for few mintue
		c <- "Get This!"
		fmt.Println("GR bailed!")
	}()

	select {
	case v := <-c:
		fmt.Printf("Consumer recv: %v\n", v)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("Timed out!")
	}
}
// in close channel you can always read,and fire (, ok =). you can not write to close channel
