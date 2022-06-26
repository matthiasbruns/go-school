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
