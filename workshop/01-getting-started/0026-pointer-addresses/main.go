package main

import "fmt"

type Recipe struct {
	Name                string
	IncludesFamilyStory bool
	Calories            int
}

func renameRecipe(recipe Recipe, newName string) {
	recipe.Name = newName
}

func renameRecipeWithPointer(recipe *Recipe, newName string) {
	recipe.Name = newName
}

func main() {
	grandmasStew := Recipe{
		Name:                "Grandma's Stew",
		IncludesFamilyStory: true,
		Calories:            1000,
	}
	fmt.Println(grandmasStew)

	renameRecipe(grandmasStew, "Grandma's Stew with a Twist")
	fmt.Println(grandmasStew)

	renameRecipeWithPointer(&grandmasStew, "Grandma's Stew with a Twist")
	fmt.Println(grandmasStew)
}
