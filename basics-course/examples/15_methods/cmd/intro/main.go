package main

import (
	"fmt"
)

//START_1 OMIT
type Car struct {
	Doors, HP int
}

func (c Car) IsSportsCar() bool { // HL1
	return c.HP > 100
}

func main() {
	c := Car{Doors: 3, HP: 120}

	fmt.Println(c.IsSportsCar()) // HL1
}

//END_1 OMIT

//START_2 OMIT
func IsSportsCar(c Car) bool { // HL2
	return c.HP > 100
}

//END_2 OMIT
