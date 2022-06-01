package main

import "fmt"

//START_1 OMIT
func main() {
	names := [4]string{ // HL1
		"John",   // HL1
		"Paul",   // HL1
		"George", // HL1
		"Ringo",  // HL1
	} // HL1
	fmt.Println(names)

	a := names[0:2] // HL1
	b := names[1:3] // HL1
	fmt.Println(a, b)

	b[0] = "XXX" // HL1
	fmt.Println(a, b)
	fmt.Println(names)

}

//END_1 OMIT
