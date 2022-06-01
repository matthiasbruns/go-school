package main

import (
	"fmt"
	"log"
)

func intro() {
	//START_1 OMIT
	var a [10]int
	//END_1 OMIT

	log.Println(a)
}

func main() {
	//START_2 OMIT
	var a [2]string // HL2
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13} // HL2
	fmt.Println(primes)
	//END_2 OMIT
}
