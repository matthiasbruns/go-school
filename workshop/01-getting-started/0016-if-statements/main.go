package main

import "fmt"

var (
	guess = 10
	rnd   = 5
)

func main() {
	// TODO: get rnd number and guess

	if guess == rnd {
		fmt.Println("You win!")
	} else if guess < rnd {
		fmt.Println("Too low!")
	} else {
		fmt.Println("Too high!")
	}
}
