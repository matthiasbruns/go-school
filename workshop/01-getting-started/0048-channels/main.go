package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Recipe struct {
	Name    string
	Minutes int
}

var recipes = []Recipe{
	{
		Name:    "Spaghetti Carbonara",
		Minutes: 20,
	},
	{
		Name:    "Spaghetti Bolognese",
		Minutes: 30,
	},
	{
		Name:    "Spaghetti al Pesto",
		Minutes: 15,
	},
}

func cook(recipe Recipe) {
	fmt.Printf("Cooking %s for %d minutes\n", recipe.Name, recipe.Minutes)
	time.Sleep(time.Duration(recipe.Minutes) * time.Second)
	fmt.Printf("Finished cooking %s\n", recipe.Name)
}

func main() {
	orders := make(chan Recipe)

	// this needs to be non-blocking, otherwise it will block the main goroutine
	go func() {
		// server in your restaurant accepts orders
		// and sends them to the kitchen
		for {
			// get random recipe from recipes
			r := recipes[rand.Intn(len(recipes))]
			// send recipe to kitchen
			orders <- r
		}
	}()

	// kitchen will cook the recipes
	for recipe := range orders {
		// cook the recipe
		cook(recipe)
	}
}
