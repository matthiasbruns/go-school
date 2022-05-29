package main

import "fmt"

//START_2 OMIT
func addOmitted(x, y int) int { // HL2
	return x + y
}

//END_2 OMIT

//START_1 OMIT
func add(x int, y int) int { // HL1
	return x + y
}

func main() {
	fmt.Println(add(42, 13)) // HL1
}

//END_1 OMIT

//START_3 OMIT
func swap(x, y string) (string, string) {
	return y, x // HL3
}

func callSwap() {
	var x, y = "hello", "world"
	x, y = swap(x, y)
	fmt.Printf("x:%s y:%s", x, y)
}

//END_3 OMIT
