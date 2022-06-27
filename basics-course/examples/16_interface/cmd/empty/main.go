package main

import "fmt"

//START_1 OMIT
func main() {
	var i interface{} // HL1
	fmt.Printf("(%v)\n", i)

	i = 42 // HL1
	fmt.Printf("(%v)\n", i)

	i = "hello" // HL1
	fmt.Printf("(%v)\n", i)
}

//END_1 OMIT
func describe(i interface{}) {
	fmt.Printf("(%v)\n", i)
}
