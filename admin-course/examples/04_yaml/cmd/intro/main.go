package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v3"
)

//START_1 OMIT
// Define data to be parsed
var data = `
a: Easy!
b:
  c: 2
  d: [3, 4]
`

// Define target structure
type Data struct {
	A string
	B struct {
		RenamedC int   `yaml:"c"`     // HL1
		D        []int `yaml:",flow"` // HL1
	}
}

//END_1 OMIT

//START_2 OMIT
func main() {
	dataStruct := Data{}

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
