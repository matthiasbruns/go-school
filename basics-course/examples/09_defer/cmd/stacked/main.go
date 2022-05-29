package main

import "fmt"

//START_1 OMIT
func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i) // HL1
	}

	fmt.Println("done")
}

//END_1 OMIT
