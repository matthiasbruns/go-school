package main

type HasIngredients interface {
	GetIngredients() map[string]int
}

type Recipe struct {
	Name        string
	Ingredients map[string]int
}

func (r Recipe) GetIngredients() map[string]int {
	return r.Ingredients
}

func printIngredients(i HasIngredients) {
	for ingredient, amount := range i.GetIngredients() {
		println(ingredient, amount)
	}
}

func main() {
	cookBook := map[string]Recipe{
		"pizza": {
			Name:        "Pizza",
			Ingredients: map[string]int{"Flour": 500, "Water": 300, "Yeast": 10},
		},
		"steak": {
			Name:        "Steak",
			Ingredients: map[string]int{"Meat": 500, "Salt": 10, "Pepper": 10},
		},
		"pudding": {
			Name:        "Pudding",
			Ingredients: map[string]int{"Milk": 500, "Water": 100, "Sugar": 100},
		},
	}

	for s, recipe := range cookBook {
		println(s, recipe.Name)
		printIngredients(recipe)
	}
}
