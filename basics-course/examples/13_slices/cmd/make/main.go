package main

import "fmt"

//START_1 OMIT
func main() {
	ids := make([]int, 5) // HL1
	printSlice("ids:", ids)

	emptyIds := make([]int, 0, 5) // HL1
	printSlice("emptyIds:", emptyIds)

	emptyIdsSubset := emptyIds[:2] // HL1
	printSlice("emptyIdsSubset:", emptyIdsSubset)

	someIds := emptyIdsSubset[2:5] // HL1
	printSlice("someIds:", someIds)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

//END_1 OMIT
