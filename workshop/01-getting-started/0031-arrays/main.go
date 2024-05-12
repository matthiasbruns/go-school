package main

import "fmt"

type Recipe struct {
	Name                string
	IncludesFamilyStory bool
	Calories            int
}

func main() {
	var recipes [2]Recipe
	recipes[0] = Recipe{
		Name:                "Grandma's Stew",
		IncludesFamilyStory: true,
		Calories:            1000,
	}
	recipes[1] = Recipe{
		Name:                "Pumpkin Pie",
		IncludesFamilyStory: true,
		Calories:            1337,
	}

	fmt.Println(recipes[1])

	fmt.Println(recipes)
}
