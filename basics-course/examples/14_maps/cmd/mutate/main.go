package main

import "fmt"

//START_1 OMIT
func main() {
	inventory := make(map[string]int) // HL1

	inventory["Fiat Punto"] = 2 // HL1
	fmt.Println("The value:", inventory["Fiat Punto"])

	inventory["Fiat Punto"] = 1 // HL1
	fmt.Println("The value:", inventory["Fiat Punto"])

	delete(inventory, "Fiat Punto") // HL1
	fmt.Println("The value:", inventory["Fiat Punto"])

	v, ok := inventory["Fiat Punto"] // HL1
	fmt.Println("The value:", v, "Present?", ok)
}

//END_1 OMIT
