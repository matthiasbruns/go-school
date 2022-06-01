package main

import "fmt"

//START_1 OMIT
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}

	// send sum to c
	c <- sum // HL1
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int) // HL1
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	// receive from c
	x, y := <-c, <-c // HL1

	fmt.Println(x, y)
}

//END_1 OMIT
