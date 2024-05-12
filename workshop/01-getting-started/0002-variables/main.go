package main

import "fmt"

var thisIsGlobal = "global" // thisIsGlobal is now "global"

func main() {
	thisIsGlobal = "still global" // thisIsGlobal is now "still global"
	fmt.Println(thisIsGlobal)

	// declaration
	var number int // number is 0
	number = 2     // number is now 2
	fmt.Println(number)

	name := "John Doe" // name is "John Doe"
	name = "Jane Doe"  // name is now "Jane Doe"
	fmt.Println(name)
}
