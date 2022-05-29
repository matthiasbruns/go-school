package main

import "fmt"

//START_1 OMIT
// i & j are ints
var i, j int = 1, 2 // HL1

func main() {
	// c & python are booleans, java is a string
	var c, python, java = true, false, "no!" // HL1

	fmt.Println(i, j, c, python, java)
}

//END_1 OMIT

//START_2 OMIT
func short() {
	var i, j int = 1, 2
	k := 3                                // HL2
	c, python, java := true, false, "no!" // HL2

	fmt.Println(i, j, k, c, python, java)
}

//END_2 OMIT
