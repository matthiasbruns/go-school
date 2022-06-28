import "encoding/json"

//...

type Car struct {
	Doors int
	Color string
}

// ... 

jsonString := `{"Doors":2, "Color":"Red"}`

var car Car
// `&car` is the address of the variable we want to store our parsed data in

err := json.Unmarshal([]byte(jsonString), &car) // HL1

//...
