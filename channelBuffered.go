// Demonstrates the use a a buffered channel of size one.
// Multicast example. Send query to all available databases and gets results
// from the fastest database.
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(42)
}

type response struct {
	db       string
	duration time.Duration
	results  []string
}

// Query Casting...
func main() {
	var wg sync.WaitGroup
	dbPool := [...]string{"db1", "db2", "db3"} 
	castQuery(&wg, dbPool, "select * from dogs")
	wg.Wait()
}

// castQuery dispatches query to all available backends and retrieve the first result
func castQuery(wg *sync.WaitGroup, dbs [3]string, query string) { // making 3 buffered channel 
	qc := make(chan response, len(dbs)) 
	for _, db := range dbs {
		wg.Add(1)
		go queryDB(wg, db, query, qc) // sending request to all 3 dbs
	}
	if r, ok := <-qc; ok { // only care about first one
		fmt.Printf("\nGot my answer! Database `%s, took (%v) -- %v\n", r.db, r.duration, r.results)
	}
}

// queryDB query a given database with the given query
func queryDB(wg *sync.WaitGroup, db, q string, qc chan<- response) {
	defer wg.Done()
	dur := time.Duration(rand.Intn(100)) * time.Millisecond
	fmt.Printf("Querying %s took %v\n", db, dur)
	time.Sleep(dur)
	qc <- response{
		db:       db,
		duration: dur,
		results:  []string{"Rex", "Cujo", "Fiffi"},
	}
}

// put your system under benchmark and try to find out the buffer size... you have to measure it..less is more
// buffer size, how may fan out decision would be depend up benchmark
// concurrency is about streaming - A-->x-->y-->z-->B
// it is not UDP braodcasting, this is http request
