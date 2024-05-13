package main

import (
	"encoding/json"
	"fmt"
)

type Meal struct {
	ID           string `json:"idMeal"`
	Name         string `json:"strMeal"`
	Instructions string `json:"strInstructions"`
}

func main() {
	jsonString := `[{  
		"idMeal":"2",
		"strMeal":"Spaghetti Carbonara",	
		"strInstructions":"Cook Spaghetti"
	}, 
	{
		"idMeal":"3",
		"strMeal":"Spaghetti Bolognese",
		"strInstructions":"Cook Spaghetti Bolognese"
	}]`

	var meals []Meal
	err := json.Unmarshal([]byte(jsonString), &meals)
	if err != nil {
		panic(err)
	}

	fmt.Println(meals)
}
