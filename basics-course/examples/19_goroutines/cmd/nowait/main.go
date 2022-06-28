package main

import (
	"fmt"
	"time"
)

//START_1 OMIT
func doSomething() {
	time.Sleep(100 * time.Millisecond) // simulate work
	fmt.Println("doSomething done")
}

func main() {
	go doSomething() // HL1

	fmt.Println("main done")
}

//END_1 OMIT
