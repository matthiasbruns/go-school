package main

import "fmt"

//START_1 OMIT
func main() {
	m := make(map[string]int) // HL1

	m["Answer"] = 42 // HL1
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48 // HL1
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer") // HL1
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"] // HL1
	fmt.Println("The value:", v, "Present?", ok)
}

//END_1 OMIT
