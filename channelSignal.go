// just want to signal, like in unix kill signal
package main

import (
	"fmt"
	"time"
)

func main() {
	var c = make(chan struct{})  // empty struct, holds no memory,
	go func(c chan<- struct{}) { // producer, it will take empty stuct channel and do sometning
		defer func() {
			fmt.Println("Backup successful! (Bucket=BumblebeeTuna)")
		}()
		fmt.Println("DB Backup to S3...")
		time.Sleep(2 * time.Second)
		c <- struct{}{} // when 11 done, pusing the empty struct to channel, signal that no data is coming
		close(c)        // close channel, it has anythign to do the resourcs
	}(c)

	_ = <-c //holding/blocking main() here for backup to finish // <-c will also work,
	// 	_, ok := <-c
	// 	fmt.Println(ok)
	//--if you send twice, the 26 will be false, if you comment 18, then it will be deadlock,and there never ever will any receiver on 25
	// 	_, ok1 := <-c
	// 	fmt.Println(ok1)
	fmt.Println("\nWe're done!")
}
