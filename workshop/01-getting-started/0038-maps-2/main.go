package main

import "fmt"

type Recipe struct {
	Name                string
	IncludesFamilyStory bool
	Calories            int
}

func main() {
	cookBook := map[string]Recipe{
		"stew": Recipe{
			Name:                "Grandma's Stew",
			IncludesFamilyStory: true,
			Calories:            1000,
		},
		"cake": Recipe{
			Name:                "Chocolate Cake",
			IncludesFamilyStory: false,
			Calories:            2000,
		},
	}

	fmt.Printf("Cookbook: %+v\n", cookBook)
}
