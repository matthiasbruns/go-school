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

	stewPtr := &grandmasStew
	fmt.Println(stewPtr)
	fmt.Println(*stewPtr)

	*stewPtr = Recipe{
		Name:                "Grandma's Stew with a Twist",
		IncludesFamilyStory: true,
		Calories:            1337,
	}

	fmt.Println(grandmasStew)
}
