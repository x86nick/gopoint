package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go func(c <-chan int) { //  receiver of data  // anonymous function, signature, c<- chan is reader here, chan <-- int, it is writer
		for {
			if v, ok := <-c; !ok { // imp, it will exit, when it all 5 are recived, <-c is blocking operation, waiting for data to be received
				break
			} else {
				fmt.Println("Recv", v)
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(c)

	for i := 0; i < 5; i++ { //  producer of data
		fmt.Println("Send", i)
		c <- i
	}
	close(c) // close once 5 are recived by receiver, closing always happen on producer side.because he know how much data will be there
}

// c <-chan int: reader from channel
// c chan <- int: writer to channel
