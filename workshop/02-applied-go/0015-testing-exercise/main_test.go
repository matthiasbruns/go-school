package main

import "testing"

func TestRecipe_CalculateTotalCalories(t *testing.T) {
	recipe := Recipe{
		Name: "Test Recipe",
		Ingredients: map[string]Ingredient{
			"ingredient1": {Calories: 100, Amount: 1},
			"ingredient2": {Calories: 200, Amount: 2},
		},
	}

	totalCalories := recipe.CalculateTotalCalories()

	if totalCalories != 500 {
		t.Errorf("Expected total calories to be 500, got %d", totalCalories)
	}
}
