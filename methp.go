// Implement a factory floor supervisor/worker model.
// Define two user defined types *worker* and *supervisor*.
//  o Both types must have an ID(int) and a name(string).
//  o A supervisor must have a level(int) to indicate her supervisor level type ie 1, 2, 3, etc...
//  o Worker support 2 functions:
//   o **badge** to generate a string with the worker ID and name
//   o **setName** to update the worker name
// o Similarly a supervisor must support badge and setName, but **badge** must also generate the supervisor level!
// o Additionally users of the system must be able to update the supervisor level
package main

import "fmt"

type (
	worker struct {
        name string	    
		id   int
	}

	supervisor struct {
		worker
		level int
	}
)

func (w *worker) badge() string {
	return fmt.Sprintf("[%d] %s", w.id, w.name)
}

func (w *worker) setName(n string) {
	w.name = n
}

func (s *supervisor) badge() string {
	return fmt.Sprintf("%s (%d)", s.worker.badge(), s.level)
}

func (s *supervisor) setLevel(l int) {
	s.level = l
}

func main() {
	w := worker{id: 100, name: "Fred"}
	s := supervisor{worker: worker{"Frank", 200}, level: 10}
	fmt.Printf("%s\n%s\n", w.badge(), s.badge())

	w.setName("Freddy")
	s.setName("Franky")
	s.setLevel(11)
	fmt.Printf("\n%s\n%s\n", w.badge(), s.badge())
}
