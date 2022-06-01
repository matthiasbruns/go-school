package main

import "fmt"

//START_1 OMIT
func main() {
	q := []int{2, 3, 5, 7, 11, 13} // HL1
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true} // HL1
	fmt.Println(r)

	s := []struct { // HL1
		i int  // HL1
		b bool // HL1
	}{ // HL1
		{2, true},   // HL1
		{3, false},  // HL1
		{5, true},   // HL1
		{7, true},   // HL1
		{11, false}, // HL1
		{13, true},  // HL1
	} // HL1
	fmt.Println(s)
}

//END_1 OMIT
