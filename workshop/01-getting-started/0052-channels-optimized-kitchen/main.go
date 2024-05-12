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
	orders := make(chan []Recipe)

	go func() {
		for {
			randomOrderCount := rand.Intn(3) + 1
			o := make([]Recipe, randomOrderCount)
			for i := 0; i < randomOrderCount; i++ {
				o[i] = recipes[rand.Intn(len(recipes))]
			}
			// send orders to kitchen
			orders <- o
		}
	}()

	for recipes := range orders {
		fmt.Printf("Received order for %d recipes\n", len(recipes))
		for _, recipe := range recipes {
			cook(recipe)
		}
	}
}
