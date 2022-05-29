package main

import "fmt"

type Vertex struct {
	X, Y int
}

//START_1 OMIT
var (
	v1 = Vertex{1, 2}  // has type Vertex // HL1
	v2 = Vertex{X: 1}  // Y:0 is implicit // HL1
	v3 = Vertex{}      // X:0 and Y:0 // HL1
	p  = &Vertex{1, 2} // has type *Vertex // HL1
)

func main() {
	fmt.Println(v1, p, v2, v3)
}

//END_1 OMIT
