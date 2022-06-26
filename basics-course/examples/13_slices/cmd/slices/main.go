package main

import "fmt"

//START_1 OMIT
func main() {
	manufacturers := []string{"Opel", "BMW", "Tesla", "Seat", "Audi", "Porsche"} // HL1
	fmt.Println("manufacturers:", manufacturers)

	years := []int{1985, 2001, 2015, 2022} // HL1
	fmt.Println("years:", years)

	cars := []struct { // HL1
		Doors       int  // HL1
		HadAccident bool // HL1
	}{ // HL1
		{2, true},  // HL1
		{3, false}, // HL1
		{5, true},  // HL1
	} // HL1
	fmt.Println(cars)
}

//END_1 OMIT
