// Implement a single link list delete procedure
//  o Create a user defined type named node with 2 fields: s:string and next of type node
//  o Say you have the following link list: a->b->c->d->e
//  o Implement a delete operation where your are only given the node to be deleted ie node("c")
//  o So after delete(c) the list will be: a->b->d->e
//  o Use the provided dump function to print out your list before and after the delete
package main

import "fmt"

type node struct {
	s    string
	next *node
}

func main() {
	na, nb, nc, nd, ne := node{s: "a"}, node{s: "b"}, node{s: "c"}, node{s: "d"}, node{s: "e"}
	na.next, nb.next, nc.next, nd.next = &nb, &nc, &nd, &ne
	h := &na
	dump(h, "Before")
	// Delete c!
	nc.s = nc.next.s
	nc.next = nc.next.next
	dump(h, "After")
}

// Prints out a list...
func dump(n *node, m string) {
	fmt.Printf("%-6s: ", m)
	for c := n; c != nil; c = c.next {
		fmt.Printf("%s -> ", c.s)
	}
	fmt.Printf("nil\n")
}
