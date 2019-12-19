// Lucy and Ethel are chocolate factory workers trying to make a living
// packaging chocolates coming from a rolling machine.
//
// Write a program to simulate Lucy and Ethel wrapping the chocolates coming
// from a chocolate roller using a buffered channels.
// The program should output how many chocolates have been wrapped vs dropped
// at the end of the run.
// A chocolate is considered dropped if it's been on the roller and not picked
// up by a worker for more than 1 second.
//
// * BONUS! Figure out the optimum worker speed for Lucy/Ethel to either wrap
// all chocolates, drop them all or process 50% of the chocolates.
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Choco struct {
	// ID is the unique id of the chocolate on the roller
	ID int
	// Timestamp is the time when the chocolate was put on the roller
	TimeStamp time.Time
}

// Expired checks whether a chocolate has been on the roller for more than 1 sec.
func (c Choco) Expired() bool {
	return time.Since(c.TimeStamp) > time.Second
}

const (
	// RunSize tracks the size of the buffer channel
	runSize = 5
	// ChocolateCount tracks the total number of chocolates to process in the run
	chocolateCount = 10
)

func main() {
	var (
		wg sync.WaitGroup
		// Wrapped tracks the number of wrapped chocolates
		wrapped int64
		// Dropped tracks the number of dropped chocolates
		dropped int64
		// Buffered channel representing chocos on roller
		coco    = make(chan Choco, runSize)
	)

	wg.Add(2)
	go factoryGirl(&wg, "Lucy", 4, coco, &wrapped, &dropped)
	go factoryGirl(&wg, "Ethel", 2, coco, &wrapped, &dropped)
	roller(coco, 1)
	wg.Wait()
	printTally(wrapped, dropped)
}

// FactoryGirl represents a chocolate factory worker
// It has the name of the worker, the speed of the worker and a channel
// to consume chocolates from the roller
func factoryGirl(wg *sync.WaitGroup, name string, speed float64, coco <-chan Choco, w, d *int64) {
	defer wg.Done()
	for c := range coco {
		if c.Expired() {
			fmt.Printf("[%-6s] ;( missed (%d)\n", name, c.ID)
			atomic.AddInt64(d, 1)
			continue
		}
		time.Sleep(time.Duration(speed) * time.Second)
		fmt.Printf("[%-6s] <- wraps (%d)\n", name, c.ID)
		atomic.AddInt64(w, 1)
	}
}

// Roller represent the chocolate roller machine
// It sends chocolates to a buffer channel at the given speed 500ms or faster
// to signal workers chocolates a ready to be wrapped.
func roller(coco chan<- Choco, speed float64) {
	for i := 0; i < chocolateCount; i++ {
		time.Sleep(time.Duration(500/speed) * time.Millisecond)
		fmt.Printf("[%-6s] -> choco (%d)\n", "Roller", i)
		coco <- Choco{ID: i, TimeStamp: time.Now()}
	}
	close(coco)
}

// PrintTally prints out chocolate batch results
func printTally(wrapped, dropped int64) {
	fmt.Println("\n--------------------------------------")
	fmt.Printf("%-10s : %02d\n", "Total", chocolateCount)
	fmt.Printf("%-10s : %02d\n", "Wrapped", wrapped)
	fmt.Printf("%-10s : %02d\n", "Dropped", dropped)
	fmt.Println("--------------------------------------")
}
