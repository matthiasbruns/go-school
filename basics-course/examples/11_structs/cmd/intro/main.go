package main

import "fmt"

//START_1 OMIT
type Car struct { // HL1
	Doors int
	HP    int
} // HL1

func main() {
	fmt.Println(Car{5, 200})
}

//END_1 OMIT

func someeStructs() {
	//START_2 OMIT
	cars := []struct { // HL2
		Doors int // HL2
		HP    int // HL2
	}{
		{2, 100}, // HL2
		{3, 125},
		{5, 220},
	}
	fmt.Println(cars)
	//END_2 OMIT
}
