package main

import (
	"fmt"
)

type vehicle struct {
	year                int
	maker, model, color string
}

type car struct {  // anonymous composition
    vehicle
    coupe bool
}

type truck struct {
    vehicle
    payLoad int
}

func main() {
    s := []vehicle { // build slice of vehical
        car{vehicle{year:2019, maker: "Ford", model: "POS"}, true},
        truck{vehicle{year:2000, maker: "Ford", model: "FindOnRoadDead"}, 5_000},
    }
	fmt.Printf("%#v\n", s)
}

// car and truck is composed of vehicals, thats why it allowing me to do this.
// this will not work 
//./prog.go:24:12: cannot use car literal (type car) as type vehicle in array or slice literal
//./prog.go:25:14: cannot use truck literal (type truck) as type vehicle in array or slice literal
