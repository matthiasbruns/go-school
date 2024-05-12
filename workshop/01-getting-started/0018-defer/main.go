package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("dude")
	defer fmt.Println("world")

	fmt.Println("hello")
}
