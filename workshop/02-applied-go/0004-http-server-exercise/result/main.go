package main

import (
	"fmt"
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
	http.HandleFunc("GET /recipes", func(w http.ResponseWriter, r *http.Request) {
		var response string
		for id, recipe := range recipes {
			response += fmt.Sprintf("%s: %s\n", id, recipe.Name)
		}
		_, err := w.Write([]byte(response))
		if err != nil {
			http.Error(w, "failed to write response", http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("GET /recipes/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if id == "" {
			http.Error(w, "missing id", http.StatusBadRequest)
			return
		}

		recipe, ok := recipes[id]
		if !ok {
			http.Error(w, "recipe not found", http.StatusNotFound)
			return
		}

		var response string
		response += "Name: " + recipe.Name + "\n"
		response += "Ingredients:\n"
		for _, ingredient := range recipe.Ingredients {
			response += "- " + ingredient + "\n"
		}
		response += "Steps:\n"
		for _, step := range recipe.Steps {
			response += "- " + step + "\n"
		}
		_, err := w.Write([]byte(response))
		if err != nil {
			http.Error(w, "failed to write response", http.StatusInternalServerError)
			return
		}
	})

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		panic(err)
	}

}
