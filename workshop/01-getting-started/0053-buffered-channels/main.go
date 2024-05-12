package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	println(<-ch)
	ch <- 3
	ch <- 4 // will block since we already have 2 elements in the channel
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
