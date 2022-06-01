package main

import "fmt"

//START_1 OMIT
// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int { // HL1
	for i, v := range s {
		// v and x are type T, which has the comparable constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15)) // HL1

	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello")) // HL1
}

//END_1 OMIT
