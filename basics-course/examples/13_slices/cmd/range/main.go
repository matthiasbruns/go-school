package main

import "fmt"

//START_1 OMIT
var manufacturers = []string{"BMW", "VW", "Opel", "Audi"} // HL1

func main() {
	for i, m := range manufacturers { // HL1
		fmt.Printf("%d %s\n", i, m)
	}
}

//END_1 OMIT
