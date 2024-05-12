package main

import (
	"fmt"
	"sync"
	"time"
)

type Recipe struct {
	Name     string
	BakeTime int
}

func bake(wg *sync.WaitGroup, recipe Recipe) {
	defer wg.Done()

	fmt.Printf("Baking %s for %d minutes\n", recipe.Name, recipe.BakeTime)
	time.Sleep(time.Duration(recipe.BakeTime) * time.Second)
	fmt.Printf("Finished baking %s\n", recipe.Name)
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	fmt.Println("Starting to bake")

	recipe := Recipe{
		Name:     "Cake",
		BakeTime: 30,
	}

	go bake(&wg, recipe)

	wg.Wait()

	fmt.Println("Finished baking")
}
