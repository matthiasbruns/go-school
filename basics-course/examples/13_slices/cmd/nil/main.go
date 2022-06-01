package main

import "fmt"

//START_1 OMIT
func main() {
	var s []int // HL1
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

//END_1 OMIT
