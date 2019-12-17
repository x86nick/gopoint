package main

import "fmt"

var a = 10 // package global scope // anybody can mutate it

func main() {
	a := 20 // redeclaring a, just overshadow package scope
	dump("MainA", a)
	if true {
		a := 30 // over shadow a from line 8
		dump("IfA", a)
	}
	{ // you can define scope like this
		a -= 5
		dump("BlockA", a)
		f1()
	}
	dump("MainA", a)

	f1()
	f2(a)
	dump("F3", f3())
}

func f1()      { dump("F1", a) }
func f2(a int) { dump("F2", a) }
func f3() (a int) {  // this is like var a retrun
	return
}

func dump(msg string, a int) {
	fmt.Printf("%-10s %d\n", msg, a)
}

// use vet tool to deal with overshadow 
// or use golintci tool
