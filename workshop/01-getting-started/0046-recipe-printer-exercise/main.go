package main

type Recipe struct {
	Name        string
	Ingredients map[string]int
}

// Implement stringer here

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
			Ingredients: map[string]int{"Milk": 500, "Water": 100, "Sugar": 100, "Vanilla": 10},
		},
	}

	for _, recipe := range cookBook {
		// TODO
	}
}
