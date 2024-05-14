package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
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
	{
		Name:    "Spaghetti Aglio e Olio",
		Minutes: 10,
	},
}

func main() {
	// spawn 10 goroutines
	numGoroutines := 10
	c := make(chan []Recipe)

	wg := sync.WaitGroup{}
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		// start goroutine
		go func() {
			defer wg.Done()

			// for each goroutine spawn rnd Recipes and send to channel
			r := rand.IntN(len(recipes)) + 1

			var rndRecipes []Recipe
			for i := range r {
				rndRecipes = append(rndRecipes, recipes[i])
			}

			c <- rndRecipes
		}()
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	// collect data from channel
	for rndRecipes := range c {
		fmt.Println(rndRecipes)
	}
}
