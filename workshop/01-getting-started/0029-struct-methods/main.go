package main

import "fmt"

type Recipe struct {
	Name                string
	IncludesFamilyStory bool
	Calories            int
}

type Oven struct {
	// lower-case variables are not exported (private)
	temperature int
}

func (o *Oven) Bake(recipe Recipe) {
	// bake the recipe
	fmt.Println("Baking", recipe.Name, "at", o.temperature, "degrees")
}

func (o *Oven) SetTemperature(temperature int) {
	o.temperature = temperature
}

func main() {
	grandmasStew := Recipe{
		Name:                "Grandma's Stew",
		IncludesFamilyStory: true,
		Calories:            1000,
	}

	oven := Oven{temperature: 350}
	fmt.Println(oven.temperature)

	oven.SetTemperature(375)
	fmt.Println(oven.temperature)

	oven.Bake(grandmasStew)
}
