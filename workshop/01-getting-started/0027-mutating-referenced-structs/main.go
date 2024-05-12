package main

import "fmt"

func main() {
	a := 5
	aPtr := &a
	aPtrPtr := &aPtr

	fmt.Printf("%+v\n", a)
	fmt.Printf("%+v\n", aPtr)
	fmt.Printf("%+v\n", aPtrPtr)
}
