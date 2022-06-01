package main

import (
	"fmt"
	"time"
)

func doSomething() {
	time.Sleep(100 * time.Millisecond) // simulate work
	fmt.Println("doSomething done")
}

func main() {
	go doSomething()

	fmt.Println("main done")

	// Delay real exit to illustrate the problem
	time.Sleep(1 * time.Second)
}
