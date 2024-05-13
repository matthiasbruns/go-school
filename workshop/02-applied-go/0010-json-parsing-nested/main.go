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
	Ingredients []Ingredient `json:"arrIngredients"`
}

func main() {
	jsonString := `[{  
		"idMeal":"2",
		"strMeal":"Spaghetti Carbonara",	
		"arrIngredients":[
			{
				"idIngredient":"1",
				"strIngredient":"Spaghetti",
				"strAmount":"200g"	
			},
			{
				"idIngredient":"2",
				"strIngredient":"Egg",	
				"strAmount":"2"
			}
		]
	}, 
	{
		"idMeal":"3",
		"strMeal":"Spaghetti Bolognese",
		"arrIngredients":[
			{
				"idIngredient":"1",
				"strIngredient":"Spaghetti",	
				"strAmount":"200g"
			},
			{
				"idIngredient":"3",
				"strIngredient":"Tomato Sauce",
				"strAmount":"100g"
			}
		]
	}]`

	var meals []Meal
	err := json.Unmarshal([]byte(jsonString), &meals)
	if err != nil {
		panic(err)
	}

	fmt.Println(meals)
}
