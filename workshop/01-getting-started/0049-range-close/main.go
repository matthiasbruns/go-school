package main

import "time"

func cookSteak(n int, c chan string) {
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			c <- "Medium Rare"
		} else {
			c <- "Well Done"
		}
		time.Sleep(1 * time.Second)
	}
	close(c)
}

func main() {
	c := make(chan string)

	go cookSteak(10, c)

	for steak := range c {
		println(steak)
	}

	println("Finished cooking steaks")
}
