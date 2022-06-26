package main

import "fmt"

func main() {
	//START_1 OMIT
	features := [6]string{"Engine", "Steering-Wheel", "Driver-Seat", "GPS", "Radio", "AC"} // HL1

	var bestFeatures []string = features[1:4] // HL1
	fmt.Println(bestFeatures)
	//END_1 OMIT
}
