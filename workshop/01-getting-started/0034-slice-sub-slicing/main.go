package main

import "fmt"

func main() {
	ingredients := []string{"onions", "carrots", "celery", "potatoes", "beef"}

	subSlice := ingredients[1:4]
	fmt.Println(subSlice)

	fmt.Printf("Length: %d, Capacity: %d\n", len(subSlice), cap(subSlice))

	subSlice[2] = "garlic"

	fmt.Println(subSlice)
	fmt.Println(ingredients)
}
