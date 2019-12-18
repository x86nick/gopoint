package main

import (
	"fmt"
	"time"
)

func main() {
	c, d := make(chan int), make(chan struct{}) // create 2 channel, with data 2nd
	go consumer("G1", 1000, 1000, c, d) // consumer channel, data channel 
	go consumer("G2", 500, 1000, c, d)
	producer(500, c, d)
	time.Sleep(2 * time.Second)
}

func consumer(ID string, speed, timeout int, in <-chan int, done <-chan struct{}) {  // done channel has no data, it is signal channel
	defer fmt.Println(ID, "Exited!")
	for {
		select { // looping forever, but using select here, figure out case that will be fired when one of the channel has data that time. select blocks by defult
		case <-done: // the caller, can say, i am done, i am forcing exit.
			fmt.Printf("[%s] I am done!\n", ID)
			return // 
		case <-time.After(time.Duration(timeout) * time.Millisecond): // time.After will return a channel that has time object. wait for 15ms and then time out
			fmt.Printf("[%s] Timed out!!\n", ID)
			return
		case v, ok := <-in: //  order of cases does not matter. if you forget to use ok, it will be big impact on channel.
			if !ok {
				return
			}
			fmt.Printf("[%s] Got %v\n", ID, v)
			time.Sleep(time.Duration(speed) * time.Millisecond)
		}
	}
}
// in this use select wil do scan more than one time.
// select also has default clause
func producer(rate int, out chan<- int, done chan<- struct{}) {
	for i := 0; i < 5; i++ {
		if i == 3 {
			close(done)
			break
		}
		out <- i
		time.Sleep(time.Duration(rate) * time.Millisecond)
	}
}

// when i have function f1 calling to f2, making call, i have lost control.
// i am untrusting f2 to be well behave.
// F1 --sending payload--> F2, and F2 is dump, F1 will keep blocked, idea is F1 should diconnect
// A channel is a ref time, and it can be Nil chan. 
// 
