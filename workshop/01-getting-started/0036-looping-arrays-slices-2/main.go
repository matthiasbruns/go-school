package main

import "fmt"

func main() {
	ingredients := []string{"onions", "carrots", "celery", "potatoes", "beef"}

	for _, ingredient := range ingredients {
		fmt.Println(ingredient)
	}

	for i := range ingredients {
		fmt.Println(i, ingredients[i])
	}
}
