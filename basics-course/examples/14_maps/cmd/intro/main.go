package main

import "fmt"

//START_1 OMIT
type Car struct {
	Doors, HP int
}

var inventory map[string]Car // HL1

func main() {
	inventory = make(map[string]Car) // HL1

	inventory["Fiat Punto"] = Car{3, 50}

	fmt.Println(inventory["Fiat Punto"])
}

//END_1 OMIT
