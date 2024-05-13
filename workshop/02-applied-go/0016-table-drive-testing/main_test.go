package main

import "testing"

func TestRecipe_CalculateTotalCalories(t *testing.T) {
	type fields struct {
		Name        string
		Ingredients map[string]Ingredient
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Spaghetti Carbonara should have 800 calories",
			fields: fields{
				Name: "Spaghetti Carbonara",
				Ingredients: map[string]Ingredient{
					"spaghetti": {Calories: 200, Amount: 1},
					"egg":       {Calories: 100, Amount: 1},
					"bacon":     {Calories: 300, Amount: 1},
					"cheese":    {Calories: 200, Amount: 1},
				},
			},
			want: 800,
		},
		{
			name: "Spaghetti Bolognese should have 1000 calories",
			fields: fields{
				Name: "Spaghetti Bolognese",
				Ingredients: map[string]Ingredient{
					"spaghetti": {Calories: 200, Amount: 1},
					"beef":      {Calories: 300, Amount: 1},
					"tomato":    {Calories: 200, Amount: 1},
					"cheese":    {Calories: 300, Amount: 1},
				},
			},
			want: 1000,
		},
		{
			name: "Spaghetti Aglio e Olio should have 600 calories",
			fields: fields{
				Name: "Spaghetti Aglio e Olio",
				Ingredients: map[string]Ingredient{
					"spaghetti": {Calories: 200, Amount: 2},
					"garlic":    {Calories: 100, Amount: 1},
					"olive oil": {Calories: 100, Amount: 1},
				},
			},
			want: 600,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Recipe{
				Name:        tt.fields.Name,
				Ingredients: tt.fields.Ingredients,
			}
			if got := r.CalculateTotalCalories(); got != tt.want {
				t.Errorf("CalculateTotalCalories() = %v, want %v", got, tt.want)
			}
		})
	}
}
