//START_1 OMIT
{
	"carType": "SUV",
	"doorsCount": 3 // HL1
}
//END_1 OMIT

//START_2 OMIT
type Car struct {
	Type string `json:"carType"`
	Doors int `json:"doorsCount"` // HL2
}
//END_2 OMIT
