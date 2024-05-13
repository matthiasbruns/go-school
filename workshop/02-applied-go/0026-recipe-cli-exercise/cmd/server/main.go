package main

import (
	"fmt"
	"goschool/0026-recipe-cli-exercise/api"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// create random, id and name search endpoints
	mux.HandleFunc("GET /random", func(w http.ResponseWriter, r *http.Request) {
		meal, err := api.GetRandomMeal()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, _ = fmt.Fprintf(w, "Random meal: %s", meal)
	})

	mux.HandleFunc("GET /recipe/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		meal, err := api.GetMealById(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, _ = fmt.Fprintf(w, "Meal with id %s: %s", id, meal)
	})

	mux.HandleFunc("GET /search", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")

		meals, err := api.SearchMealsByName(name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, _ = fmt.Fprintf(w, "Meals with name %s: %s", name, meals)
	})

	if err := http.ListenAndServe(":8090", mux); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
