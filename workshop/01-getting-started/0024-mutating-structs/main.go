package main

import "fmt"

type Recipe struct {
	Name                string
	IncludesFamilyStory bool
	Calories            int
}

func main() {
	grandmasStew := Recipe{
		Name:                "Grandma's Stew",
		IncludesFamilyStory: true,
		Calories:            1000,
	}
	fmt.Println(grandmasStew)

	grandmasStew.Calories = 2000
	fmt.Println(grandmasStew)
}
