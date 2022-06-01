package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

//START_1 OMIT
func (v Vertex) Abs() float64 { // HL1
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vertex) Scale(f float64) { // HL1
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}

//END_1 OMIT
