package main

import (
	"fmt"
	"math"
)

//START_1 OMIT
func sqrt(x float64) complex128 {
	if x < 0 { // HL1
		s := math.Sqrt(-x)
		return complex(0.0, s)
	} // HL1
	return complex(math.Sqrt(x), 0)
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}

//END_1 OMIT

//START_2 OMIT
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim { // HL2
		return v
	}
	return lim
}

//END_2 OMIT

func ifElse(x, n, lim float64) float64 {
	//START_3 OMIT
	if v := math.Pow(x, n); v < lim { // HL3
		return v
	} else { // HL3
		fmt.Printf("%g >= %g\n", v, lim)
	}
	//END_3 OMIT
	return lim
}
