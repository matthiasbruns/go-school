package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

//START_2 OMIT
//START_1 OMIT
func (t *T) M() {
	if t == nil { // HL1
		fmt.Println("M(): <nil>")
		return
	}
	fmt.Println("M(): " + t.S)
}

//END_1 OMIT
func main() {
	var i I // interface

	var t *T // implementing struct
	i = t
	describe(i)
	i.M() // HL1

	i = &T{"hello"}
	describe(i)
	i.M()
}

//END_2 OMIT
func describe(i I) {
	fmt.Printf("T: (%v)\n", i)
}
