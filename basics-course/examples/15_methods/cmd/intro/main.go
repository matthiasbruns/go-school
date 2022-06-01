package main

import (
	"fmt"
	"math"
)

//START_1 OMIT
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 { // HL1
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs()) // HL1
}

//END_1 OMIT

//START_2 OMIT
func Abs(v Vertex) float64 { // HL2
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//END_2 OMIT
