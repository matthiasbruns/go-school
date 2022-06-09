package main

import (
	"fmt"
	"strconv"
)

func t(ch chan string, i int) {
	for j := 0; j < 10; j++ {
		ch <- "Goroutine : " + strconv.Itoa(i) //Integer TO ASCII-String
	}
}
func main() {
	ch := make(chan string)

	for i := 0; i < 10; i++ {
		go t(ch, i)
	}

	for k := 1; k <= 100; k++ {
		fmt.Println(k, <-ch)
	}
}
