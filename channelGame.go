// The Unicorn Game. Create a program where 2 players get to exchange unicorns
// over an un-buffered channel. The game ends once either of the parties gets
// `randomly sick and tired of playing this wonderful game!
// Assume unicorns are integers
// Tip!
// You can make random number more or less predictable by setting a seed
// Fix seed
// rand.Seed(42)
// Rand seed
// rand.Seed(time.Now().UnixNano())
// Sample Run...
// <- Norbert    Got Unicorn     [00]
// -> Norbert    Send Unicorn    [01]
// <- Faranouch  Got Unicorn     [01]
// -> Faranouch  Send Unicorn    [02]
// <- Norbert    Got Unicorn     [02]
// -> Norbert    Send Unicorn    [03]
// <- Faranouch  Got Unicorn     [03]
// ;( Faranouch  Got B.O.R.E.D!
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	defDuration = 400
	defModulo   = 15
)

var wg sync.WaitGroup

func init() {
	rand.Seed(42)
}

func main() {
	c := make(chan int)
	wg.Add(2)
	go player("Faranouch", c)
	go player("Norbert", c)
	// Send out the first unicorn...
	c <- 0
	wg.Wait()
}

func player(ID string, c chan int) {
	defer wg.Done()
	for unicorn := range c {
		fmt.Printf("<- %-10s %-15s [%02d]\n", ID, "Got Unicorn", unicorn)
		if rand.Intn(100)%defModulo == 0 {
			fmt.Printf(";( %-10s %-15s\n", ID, "Got B.O.R.E.D!")
			close(c)
			return
		}
		unicorn++
		fmt.Printf("-> %-10s %-15s [%02d]\n", ID, "Send Unicorn", unicorn)
		c <- unicorn
		time.Sleep(defDuration * time.Millisecond)
	}
}
