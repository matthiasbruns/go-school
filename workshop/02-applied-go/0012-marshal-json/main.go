package main

import (
	"encoding/json"
	"fmt"
)

type Ingredient struct {
	ID     string `json:"idIngredient"`
	Name   string `json:"strIngredient"`
	Amount string `json:"strAmount"`
}

type Meal struct {
	ID          string       `json:"idMeal"`
	Name        string       `json:"strMeal"`
	Ingredients []Ingredient `json:"arrIngredients,omitempty"`
}

func main() {
	meals := []Meal{
		{
			ID:   "2",
			Name: "Spaghetti Carbonara",
			Ingredients: []Ingredient{
				{
					ID:     "1",
					Name:   "Spaghetti",
					Amount: "200g",
				},
			},
		},
		{
			ID:   "3",
			Name: "Spaghetti Bolognese",
		},
	}

	if j, err := json.Marshal(meals); err != nil {
		panic(err)
	} else {
		fmt.Println(string(j))
	}
}
