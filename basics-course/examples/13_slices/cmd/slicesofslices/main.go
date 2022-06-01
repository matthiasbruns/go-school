package main

import (
	"fmt"
	"strings"
)

//START_1 OMIT
func main() {
	// Create a tic-tac-toe board.
	board := [][]string{ // HL1
		[]string{"_", "_", "_"}, // HL1
		[]string{"_", "_", "_"}, // HL1
		[]string{"_", "_", "_"}, // HL1
	} // HL1

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

//END_1 OMIT
