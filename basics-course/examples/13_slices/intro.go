package main

import "fmt"

func main() {
	//START_1 OMIT
	primes := [6]int{2, 3, 5, 7, 11, 13} // HL1

	var s []int = primes[1:4] // HL1
	fmt.Println(s)
	//END_1 OMIT
}
