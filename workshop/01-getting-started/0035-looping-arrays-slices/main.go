package main

import "fmt"

func main() {
	ingredients := []string{"onions", "carrots", "celery", "potatoes", "beef"}

	for i, ingredient := range ingredients {
		fmt.Printf("Ingredient %d: %s\n", i+1, ingredient)
	}
}
