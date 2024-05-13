package main

type Ingredient struct {
	Calories int
	Amount   int
}

type Recipe struct {
	Name        string
	Ingredients map[string]Ingredient
}

func (r Recipe) CalculateTotalCalories() int {
	totalCalories := 0
	for _, ingredient := range r.Ingredients {
		totalCalories += ingredient.Calories * ingredient.Amount
	}
	return totalCalories
}
