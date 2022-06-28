package main

import (
	"fmt"
	"time"
)

//START_1 OMIT
func turnSignal(blinker chan string, quit chan struct{}) {
	var x, y = "on", "off"
	for { // HL1
		select {
		case blinker <- x: // HL1
			x, y = y, x
			time.Sleep(500 * time.Millisecond)
		case <-quit: // HL1
			fmt.Println("quit")
			return
		}
	}
}
func main() {
	blinker := make(chan string)
	quit := make(chan struct{})
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-blinker) // HL1
		}
		quit <- struct{}{} // HL1
	}()
	turnSignal(blinker, quit) // HL1
}

//END_1 OMIT
