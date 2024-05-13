package main

import (
	"net/http"
)

type Recipe struct {
	Name        string
	Ingredients []string
	Steps       []string
}

var recipes = map[string]Recipe{
	"1": {
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
	"2": {
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
	},
}

func main() {
	// recipes list handler
	http.HandleFunc("TODO", func(w http.ResponseWriter, r *http.Request) {
		// format all recipes and write them to the response writer
	})

	// recipes id handler
	http.HandleFunc("TODO", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		// get the recipe with the given id and write it to the response writer
	})

	// start server on port 8090
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		panic(err)
	}

}
