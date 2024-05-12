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
	const numStoves = 3
	// allow for 10 orders to be queued up per stove -> 30 orders
	orders := make(chan []Recipe, numStoves*10)

	// start goroutines for each stove
	stoves := make(chan Recipe, numStoves)

	for i := 0; i < numStoves; i++ {
		// start a goroutine for each stove

		go func() {
			for {
				// receive recipes from the stove channel
				recipe := <-stoves
				cook(recipe)
			}
		}()
	}

	go func() {
		for {
			randomOrderCount := rand.Intn(3) + 1
			o := make([]Recipe, randomOrderCount)
			for i := 0; i < randomOrderCount; i++ {
				o[i] = recipes[rand.Intn(len(recipes))]
			}
			fmt.Printf("Guest ordered %d recipes\n", len(o))
			// send orders to kitchen
			orders <- o

			time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		}
	}()

	// this is the kitchen - it takes orders and sends them to the stove queue
	for order := range orders {
		fmt.Printf("Kitchen received order for %d recipes\n", len(order))

		// enqueue the order to the stoves, if there is a stove available
		for _, recipe := range order {
			stoves <- recipe
		}
	}
}
