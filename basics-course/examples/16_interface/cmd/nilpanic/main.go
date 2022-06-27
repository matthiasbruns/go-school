package main

import "fmt"

type Vehicle interface {
	Drive() string
}

//START_1 OMIT
func main() {
	var v Vehicle // HL1
	fmt.Printf("(%v)\n", v)
	v.Drive() // will panic because v is nil
}

//END_1 OMIT
