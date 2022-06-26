package main

import "fmt"

type Car struct {
	Doors, HP int
}

//START_1 OMIT
var (
	c1 = Car{5, 100}             // has type Car // HL1
	c2 = Car{Doors: 2}           // HP:0 is implicit // HL1
	c3 = Car{}                   // Doors:0 and HP:0 // HL1
	p  = &Car{Doors: 2, HP: 500} // has type *Car // HL1
)

func main() {
	fmt.Println("c1:", c1, "\nc2:", c2, "\nc3:", c3, "\np:", p)
}

//END_1 OMIT
