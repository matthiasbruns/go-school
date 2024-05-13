package main

type Recipe struct {
	Name        string
	Ingredients []string
	Steps       []string
}

var recipes = []Recipe{
	{
		Name: "Pasta",
		Ingredients: []string{
			"pasta",
			"salt",
			"water",
		},
		Steps: []string{
			"Boil water",
			"Add salt",
			"Add pasta",

		},
	},
	{
		Name: "Salad",
		Ingredients: []string{
			"lettuce",
			"tomato",
			"cucumber",
		},
		Steps: []string{
			"Chop lettuce",
			"Chop tomato",


		},

		}
	}
}

func main() {
	// TODO
}
