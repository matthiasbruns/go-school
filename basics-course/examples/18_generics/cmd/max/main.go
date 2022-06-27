package main

import (
	"fmt"
)

//START_1 OMIT
type Ordered interface { // HL1
	~int | ~int8 | ~int16 // ... and so on
}

func Max[E Ordered](input []E) (max E) {
	for _, v := range input { // HL1
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	fmt.Println(Max([]int{1, 2, 3}))
}

//END_1 OMIT
