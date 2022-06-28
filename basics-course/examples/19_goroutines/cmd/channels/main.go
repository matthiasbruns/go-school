package main

import (
	"fmt"
	"time"
)

//START_1 OMIT
func refuel(c chan bool) {
	fmt.Println("started refuel")
	time.Sleep(2 * time.Second)
	c <- true // HL1
	fmt.Println("finished refuel")
}

func wipeWindows(c chan bool) {
	fmt.Println("started wipeWindows")
	time.Sleep(1 * time.Second)
	c <- true // HL1
	fmt.Println("finished wipeWindows")
}

func main() {
	c := make(chan bool) // HL1
	go refuel(c)
	go wipeWindows(c)

	// receive from c
	x, y := <-c, <-c // HL1

	fmt.Println("refueled:", x, "wiped:", y)
}

//END_1 OMIT
