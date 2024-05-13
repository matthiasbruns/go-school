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
	jsonString := `{  
		"idMeal":"2",
		"strMeal":"Spaghetti Carbonara",	
		"strInstructions":"Cook Spaghetti"
	}`

	var meal Meal
	err := json.Unmarshal([]byte(jsonString), &meal)
	if err != nil {
		panic(err)
	}

	fmt.Println(meal)
}
