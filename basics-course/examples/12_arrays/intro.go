package main

import (
	"fmt"
	"log"
)

func intro() {
	//START_1 OMIT
	var features [2]string
	//END_1 OMIT

	log.Println(features)
}

func main() {
	//START_2 OMIT
	var features [2]string // HL2
	features[0] = "Steering Wheel"
	features[1] = "Working Engine"
	fmt.Println(features[1])
	fmt.Println(features)

	primes := [6]int{2, 3, 5, 7, 11, 13} // HL2
	fmt.Println(primes)
	//END_2 OMIT
}
