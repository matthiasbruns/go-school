package main

import (
	"fmt"
	"strings"
)

//START_1 OMIT
func main() {
	openAstra := 0
	vwBeetle := 1
	// Create a tic-tac-toe board.
	featureMatrix := [][]string{ // HL1
		{"", "", ""},                     // HL1
		{"DAB", "GPS", "Cruise Control"}, // HL1
	} // HL1

	// The players take turns.
	featureMatrix[openAstra][0] = "Camera"
	featureMatrix[openAstra][1] = "Leather"
	featureMatrix[openAstra][2] = "GPS"
	featureMatrix[vwBeetle][2] = "Isofix"

	for i := 0; i < len(featureMatrix); i++ {
		fmt.Printf("%s\n", strings.Join(featureMatrix[i], " "))
	}
}

//END_1 OMIT
