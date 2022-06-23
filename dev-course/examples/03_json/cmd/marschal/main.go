package main

import (
	"encoding/json"
	"fmt"
)

//START_1 OMIT
type Bird struct {
	Species string `json:"birdType"`
	// we can set the "omitempty" property as part of the JSON tag
	Description string `json:"what it does,omitempty"` // HL1
}

func main() {
	pigeon := &Bird{
		Species: "Pigeon",
	}

	data, _ := json.Marshal(pigeon) // HL1

	fmt.Println(string(data))
}

//END_1 OMIT
