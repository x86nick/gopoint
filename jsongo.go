package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	b := []byte(`{"name":"Bee","metrics":["6.3","180"]}`)
	
	m := map[string]interface{}{} // no schema, empty interface , creating map, kv is string, value of it is empty interface type. loading everying with this trick.
	if err := json.NewDecoder(bytes.NewReader(b)).Decode(&m); err != nil { // leaverage io reader and writer interface, read those bytes and decode it. crossing pacakge boundry
		panic(err)
	}
	fmt.Println(m["name"], m["metrics"]) // create new data that i have read 

	buff := new(bytes.Buffer)
	if err := json.NewEncoder(buff).Encode(m); err != nil { // marshal from data structure onto jason
		panic(err)
	}
	fmt.Println(buff)
}

//interface{}{} ||||  type name| building map
