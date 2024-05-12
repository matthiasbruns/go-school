package main

import (
	"fmt"
)

type Recipe struct {
	Name     string
	Calories int
}

func main() {
	var anything interface{} = Recipe{
		Name:     "Spaghetti Carbonara",
		Calories: 1200,
	}

	recipe := anything.(Recipe)
	fmt.Println("recipe := anything.(Recipe)\n", recipe)

	recipe, ok := anything.(Recipe)
	fmt.Println("recipe, ok := anything.(Recipe)\n", recipe, ok)

	f, ok := anything.(float64)
	fmt.Println("f, ok := anything.(float64)", f, ok)

	f = anything.(float64)
	fmt.Println("f := anything.(float64)", f)
}
