package main

import "fmt"

//START_1 OMIT
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for { // HL1
		select {
		case c <- x: // HL1
			x, y = y, x+y
		case <-quit: // HL1
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c) // HL1
		}
		quit <- 0 // HL1
	}()
	fibonacci(c, quit) // HL1
}

//END_1 OMIT
