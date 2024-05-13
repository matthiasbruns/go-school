package main

import (
	"encoding/json"
)

type Meal struct {
	ID           string `json:"idMeal"`
	Name         string `json:"strMeal"`
	Instructions string `json:"strInstructions"`
}

func main() {
	jsonString := `{  
		"idMeal":"1",
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
	}`

	var result map[string]interface{}
	if err := json.Unmarshal([]byte(jsonString), &result); err != nil {
		panic(err)
	}

	ingredients := result["arrIngredients"].([]interface{})

	for i, ingredient := range ingredients {
		ingredientMap := ingredient.(map[string]interface{})
		println(i, ingredientMap["idIngredient"].(string), ingredientMap["strIngredient"].(string), ingredientMap["strAmount"].(string))
	}

}
