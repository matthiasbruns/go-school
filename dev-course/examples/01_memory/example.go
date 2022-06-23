package main

import "fmt"

//START_1 OMIT
func main() {
	n := 4
	inc(&n) // HL1
	fmt.Println(n)
}
func inc(x *int) {
	*x++ // HL1
}

//END_1 OMIT
