package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v3"
)

//START_1 OMIT
// Define data to be parsed
var data = `
Model: VW Golf
Config:
  Doors: 2
  Features: [Engine, Headlights]
`

// Define target structure
type Car struct {
	Model  string
	Config struct {
		Doors    int      `yaml:"Doors"` // HL1
		Features []string `yaml:",flow"` // HL1
	}
}

//END_1 OMIT

//START_2 OMIT
func main() {
	dataStruct := Car{}

	// parse yaml to struct
	err := yaml.Unmarshal([]byte(data), &dataStruct) // HL2
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", dataStruct)

	// marschal struct to yaml
	marshaled, err := yaml.Marshal(&dataStruct) // HL2
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Printf("--- t dump:\n%s\n\n", string(marshaled))
}

//END_2 OMIT
