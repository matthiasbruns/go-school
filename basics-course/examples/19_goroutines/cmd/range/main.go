package main

import (
	"fmt"
)

//START_1 OMIT
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) // HL1
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c { // HL1
		fmt.Println(i)
	}
}

//END_1 OMIT
