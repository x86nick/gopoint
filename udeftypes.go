import "fmt"

type (
	car struct {
		year                int
		maker, model, color string
	}
	monsterTruck struct {  // annonymous struct bbased on car
		car
		liftHeight int
	}
)

func main() {
	bubbas := monsterTruck{
		liftHeight: 40,
		car: car{ // Note! Default to type name
			year:  1971,
			maker: "Ford",
			model: "Pinto",
			color: "Puke Green",
		},
	}
	fmt.Printf("Make:   %s\n", bubbas.maker) // Note! , it is not bubbas.car.maker, but both will work, you can assume it is part of struct
	fmt.Printf("Color:  %s\n", bubbas.color)
	fmt.Printf("Height: %din lift Babby!\n", bubbas.liftHeight)
}
// a monster truck - it is composition and not inheritance... 
//
