package main

import "fmt"

//START_1 OMIT
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128} // HL1

func main() {
	for i, v := range pow { // HL1
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

//END_1 OMIT
