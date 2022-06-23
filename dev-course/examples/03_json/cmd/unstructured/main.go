package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//START_1 OMIT
	birdJson := `{"birds":{"pigeon":"likes to perch on rocks","eagle":"bird of prey"},"animals":"none"}`

	var result map[string]interface{}         // HL1
	json.Unmarshal([]byte(birdJson), &result) // HL1

	// The object stored in the "birds" key is also stored as
	// a map[string]interface{} type, and its type is asserted from
	// the interface{} type
	birds := result["birds"].(map[string]interface{}) // HL1

	for key, value := range birds {
		// Each value is an interface{} type, that is type asserted as a string
		fmt.Println(key, value.(string)) // HL1
	}
	//END_1 OMIT
}
