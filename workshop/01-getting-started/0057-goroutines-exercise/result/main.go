package main

import (
	"fmt"
	"math/rand"
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
	recipeChan := make(chan Recipe)

	wg := sync.WaitGroup{}
	// spawn 10 goroutines
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			// random amount of recipes
			rnd := rand.Intn(10)
			for j := 0; j < rnd; j++ {
				recipe := recipes[rand.Intn(len(recipes))]

				// add recipe to channel
				recipeChan <- recipe
			}
		}(i)
	}

	go func() {
		for recipe := range recipeChan {
			fmt.Println(recipe)
		}
	}()

	wg.Wait()
	close(recipeChan)
}

// https://go.dev/play/p/zKZchd06dbf
