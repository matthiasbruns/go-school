package main

import (
	"encoding/json"
	"fmt"
)

//START_1 OMIT
type Car struct {
	Doors int `json:"doorsAmount"`
	// we can set the "omitempty" property as part of the JSON tag
	Color string `json:"color,omitempty"` // HL1
}

func main() {
	car := &Car{ // HL1
		Doors: 3,
		Color: "black",
	}

	data, _ := json.Marshal(car) // HL1

	fmt.Println(string(data))
}

//END_1 OMIT
