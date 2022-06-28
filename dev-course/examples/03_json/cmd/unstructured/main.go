package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//START_1 OMIT
	carJson := `{
		"cars":{"vm polo":"nice car to drive",
		"audi a8":"even nicer car to drive"},"bikes":"none"
	}`

	var result map[string]interface{}        // HL1
	json.Unmarshal([]byte(carJson), &result) // HL1

	// The object stored in the "cars" key is also stored as a map[string]interface{} type,
	// and its type is asserted from the interface{} type
	cars := result["cars"].(map[string]interface{}) // HL1

	for key, value := range cars {
		fmt.Println(key, ":", value.(string)) // HL1
	}
	//END_1 OMIT
}
