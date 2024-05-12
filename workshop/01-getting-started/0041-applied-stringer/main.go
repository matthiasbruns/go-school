package main

import "fmt"

type Recipe struct {
	Name                string
	IncludesFamilyStory bool
	Calories            int
}

func (r Recipe) String() string {
	return fmt.Sprintf("==============\n%s\n\tIncludesFamilyStory: %t, Calories: %d\n==============\n", r.Name, r.IncludesFamilyStory, r.Calories)
}

func main() {
	grandmasStew := Recipe{
		Name:                "Grandma's Stew",
		IncludesFamilyStory: true,
		Calories:            1000,
	}

	fmt.Println(grandmasStew)
}
