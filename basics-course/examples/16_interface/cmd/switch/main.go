package main

import "fmt"

//START_1 OMIT
func do(i interface{}) {
	switch v := i.(type) { // HL1
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

//END_1 OMIT

func main() {
	do(21)
	do("hello")
	do(true)
}
