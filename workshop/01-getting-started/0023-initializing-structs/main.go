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

	homeMadePizza := Recipe{
		"Home Made Pizza",
		false,
		2000,
	}
	fmt.Println(homeMadePizza)
}
