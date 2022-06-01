package main

import "fmt"

//START_1 OMIT
func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	ch <- 3 // will block and not allow the next lines to execute
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

//END_1 OMIT
