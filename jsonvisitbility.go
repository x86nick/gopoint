package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	b := []byte(`{"name":"Fernand","metrics":[6.3, 180]}`) // creating byte of stream

	j := struct { // i can producde this anonymous structure..that i am defining
		name         string
		measurements []float32 `json:"metrics",omitempty` // little dsl `` allows you to do simple manipulation, mapping to my attribute name - measurement
	}{} // i passed the address and i have instance
	if err := json.NewDecoder(bytes.NewReader(b)).Decode(&j); err != nil { // do marshaling and unmarshaling of data
		panic(err)
	}
	// Ruroh?...
	fmt.Printf("%#v\n",j)
	fmt.Println(j.name, j.measurements)

	buff := new(bytes.Buffer)
	if err := json.NewEncoder(buff).Encode(j); err != nil {
		panic(err)
	}
	fmt.Println(buff)
}

// []
// {}
// measurements  ==> this should have Measurements for visitibilty outside
