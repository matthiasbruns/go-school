package main

import "fmt"

//START_1 OMIT
type Vertex struct { // HL1
	X int
	Y int
} // HL1

func main() {
	fmt.Println(Vertex{1, 2})
}

//END_1 OMIT
