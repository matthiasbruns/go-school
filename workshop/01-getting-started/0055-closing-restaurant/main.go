package main

import (
	"fmt"
	"math/rand"
	"sync"
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

func handleOrder(orders []Recipe) {
	fmt.Printf("Received order for %d orders\n", len(orders))
	wg := &sync.WaitGroup{}
	wg.Add(len(orders))

	// cook recipes in parallel per order
	for _, recipe := range orders {
		go func(r Recipe) {
			defer wg.Done()
			cook(r)
		}(recipe)
	}
	wg.Wait()
}

func main() {
	ordersChan := make(chan []Recipe)
	closingTime := time.After(10 * time.Second)

	go func() {
		for {
			randomOrderCount := rand.Intn(3) + 1
			o := make([]Recipe, randomOrderCount)
			for i := 0; i < randomOrderCount; i++ {
				o[i] = recipes[rand.Intn(len(recipes))]
			}
			// send orders to kitchen
			ordersChan <- o
		}
	}()

	for {
		// priority to closing time
		select {
		case <-closingTime:
			fmt.Println("Closing time! Kitchen is closing...")
			close(ordersChan)
			return
		default:
			// do nothing, will skip to next select
		}

		select {
		case dishes := <-ordersChan:
			handleOrder(dishes) // moved to function for readability
		default:
			// do nothing, will restart loop
		}
	}
}
