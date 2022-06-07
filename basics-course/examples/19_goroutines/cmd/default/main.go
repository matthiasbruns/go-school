package main

import (
	"fmt"
	"time"
)

//START_1 OMIT
func main() {
	tick := time.Tick(100 * time.Millisecond)  // triggers after x repeatedly
	boom := time.After(500 * time.Millisecond) // triggers after x once

	for {
		select {
		case <-tick: // HL1
			fmt.Println("tick.")
		case <-boom: // HL1
			fmt.Println("BOOM!")
			return
		default: // HL1
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

//END_1 OMIT
