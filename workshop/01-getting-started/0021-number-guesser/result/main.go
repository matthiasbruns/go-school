package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

// number guessing game
func main() {
	// 1. generate a random number
	rnd := rand.Intn(100)

	for {
		// read the user input
		println("Guess the number")
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			continue
		}

		// convert the user input to an integer
		guess, err := strconv.Atoi(input)
		if err != nil {
			println("Invalid input")
			// skip the rest of the loop and ask for input again
			continue
		}

		// 2. compare the user input with the random number
		if guess == rnd {
			// 3. if the user input is equal to the random number
			println("You win\n🥳\n\n\n")
		} else if guess > rnd {
			println("Too high")
		} else {
			println("Too low")
		}
	}
}

// https: //go.dev/play/p/mHqOLF50vwH
