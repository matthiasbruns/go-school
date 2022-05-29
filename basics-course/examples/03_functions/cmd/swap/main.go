package main

import "fmt"

//START_1 OMIT
func swap(x, y string) (string, string) {
	return y, x // HL1
}

func main() {
	var x, y = "hello", "world"
	x, y = swap(x, y) // HL1
	fmt.Printf("x:%s y:%s", x, y)
}

//END_1 OMIT
