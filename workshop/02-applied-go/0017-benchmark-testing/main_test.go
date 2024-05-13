package main

import "testing"

func BenchmarkRecipe_CalculateTotalCalories(b *testing.B) {
	recipe := Recipe{
		Name: "Test Recipe",
		Ingredients: map[string]Ingredient{
			"ingredient1": {Calories: 100, Amount: 1},
			"ingredient2": {Calories: 200, Amount: 2},
		},
	}

	for i := 0; i < b.N; i++ {
		recipe.CalculateTotalCalories()
	}
}
