//START_1 OMIT
{
	"Doors": 5,
	"Color": "Red"
}
//END_1 OMIT

//START_2 OMIT
type Car struct {
	Doors int
	Color string
}
//END_2 OMIT

//START_3 OMIT
jsonString := `{"Doors":2, "Color":"Red"}`
var car Car	
json.Unmarshal([]byte(jsonString), &car) // HL3
fmt.Printf("Doors: %d, Color: %s", car.Doors, car.Color)
//END_3 OMIT
