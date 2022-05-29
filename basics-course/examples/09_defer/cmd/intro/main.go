package main

import "fmt"

//START_1 OMIT
func main() {
	defer fmt.Println("world") // HL1

	fmt.Println("hello")
}

//END_1 OMIT
