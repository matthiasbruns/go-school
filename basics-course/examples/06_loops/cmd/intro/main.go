package main

import "fmt"

//START_1 OMIT
func main() {
	sum := 0

	//   init ;  cond ; post
	for i := 0; i < 10; i++ { // HL1
		// loop body
		sum += i
	} // HL1
	fmt.Println(sum)
}

//END_1 OMIT

func while() {
	//START_2 OMIT
	sum := 1
	for sum < 1000 { // HL2
		sum += sum
	}
	fmt.Println(sum)

	//END_2 OMIT
}

//END_2 OMIT

func endless() {
	//START_3 OMIT
	for { // HL3
		// called endlessly
	}
	//END_3 OMIT
}
