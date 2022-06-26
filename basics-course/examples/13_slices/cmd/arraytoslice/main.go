package main

import "fmt"

//START_1 OMIT
func main() {
	manufacturers := [4]string{ // HL1
		"Volkswagen", // HL1
		"BMW",        // HL1
		"Opel",       // HL1
		"Seat",       // HL1
	} // HL1
	fmt.Println(manufacturers)

	a := manufacturers[0:2] // HL1
	b := manufacturers[1:3] // HL1
	fmt.Println("a:", a, "\nb:", b)

	fmt.Println("\n----- Update b -----\n")

	b[0] = "Porsche" // HL1
	fmt.Println("a:", a, "\nb:", b)
	fmt.Println(manufacturers)

}

//END_1 OMIT
