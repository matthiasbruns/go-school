package main

import (
	"fmt"
	"sync"
	"time"
)

type Recipe struct {
	Name    string
	Minutes int
}

type Oven struct {
	m sync.Mutex
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

func (o *Oven) cook(recipe Recipe) {
	o.m.Lock()
	defer o.m.Unlock()

	fmt.Printf("Cooking %s for %d minutes\n", recipe.Name, recipe.Minutes)
	time.Sleep(time.Duration(recipe.Minutes) * time.Second)
	fmt.Printf("Finished cooking %s\n", recipe.Name)
}

func main() {
	o := Oven{}

	now := time.Now()

	o.cook(recipes[0])
	o.cook(recipes[1])

	fmt.Printf("Cooking took %v\n", time.Since(now))
}
