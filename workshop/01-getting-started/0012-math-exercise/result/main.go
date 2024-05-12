package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Enter the radius of the circle:")
	var radius float64
	_, err := fmt.Scanf("%f", &radius)
	if err != nil {
		fmt.Println(err)
		return
	}
	diameter := 2 * radius
	circumference := 2 * math.Pi * radius
	area := math.Pi * math.Pow(radius, 2)
	fmt.Printf("Diameter: %v\nCircumference: %v\nArea: %v", diameter, circumference, area)
}
