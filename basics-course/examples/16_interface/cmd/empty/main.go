package main

import "fmt"

//START_1 OMIT
func main() {
	var i interface{} // HL1
	describe(i)

	i = 42 // HL1
	describe(i)

	i = "hello" // HL1
	describe(i)
}

//END_1 OMIT
func describe(i interface{}) {
	fmt.Printf("(%v)\n", i)
}
