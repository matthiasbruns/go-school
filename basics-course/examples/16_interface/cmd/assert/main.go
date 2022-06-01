package main

import "fmt"

//START_1 OMIT
func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64) // what is f?
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}

//END_1 OMIT
