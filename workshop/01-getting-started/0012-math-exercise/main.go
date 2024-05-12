package main

import (
	"fmt"
	"math"
)

func main() {
	// calculate the Euclidean distance between two points
	// (x1, y1) and (x2, y2)
	x1, y1 := 0.0, 0.0
	x2, y2 := 3.0, 4.0

	dist := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	fmt.Println(dist)
}
