package main

import "fmt"

type I interface {
	M()
}

//START_1 OMIT
func main() {
	var i I // HL1
	describe(i)
	i.M() // will panic since i is nil
}

//END_1 OMIT

func describe(i I) {
	fmt.Printf("(%v)\n", i)
}
