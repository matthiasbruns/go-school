package main

import "fmt"

type Recipe struct {
	Name                string
	IncludesFamilyStory bool
	Calories            int
}

func main() {
	// if no initial value is provided, the map is nil
	cookBook := make(map[string]Recipe)

	cookBook["stew"] = Recipe{
		Name:                "Grandma's Stew",
		IncludesFamilyStory: true,
		Calories:            1000,
	}

	cookBook["cake"] = Recipe{
		Name:                "Chocolate Cake",
		IncludesFamilyStory: false,
		Calories:            2000,
	}

	fmt.Printf("Cookbook: %+v\n", cookBook)
}
