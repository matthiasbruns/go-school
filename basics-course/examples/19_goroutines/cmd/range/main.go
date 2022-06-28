package main

import (
	"fmt"
	"time"
)

//START_1 OMIT
func turnSignal(n int, c chan string) {
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			c <- "on" // HL1
		} else {
			c <- "off" // HL1
		}
		time.Sleep(250 * time.Millisecond)
	}
	close(c) // HL1
}

func main() {
	c := make(chan string, 10)
	go turnSignal(cap(c), c)
	for i := range c { // HL1
		fmt.Println(i)
	}
}

//END_1 OMIT
