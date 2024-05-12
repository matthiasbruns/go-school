package main

import "fmt"

type Recipe struct {
	Name                string
	IncludesFamilyStory bool
	Calories            int
}

func main() {
	recipes := make([]Recipe, 0, 2)
	recipes = append(recipes, Recipe{
		Name:                "Grandma's Stew",
		IncludesFamilyStory: true,
		Calories:            1000,
	})

	recipes = append(recipes, Recipe{
		Name:                "Pumpkin Pie",
		IncludesFamilyStory: true,
		Calories:            1337,
	})

	fmt.Println(recipes[1])

	fmt.Println(recipes)
}
