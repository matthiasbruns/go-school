//START_1 OMIT
{
  "doors": 2,
  "color": "red"
  "dimensions": {
    "height": 24,
    "width": 10
  }
}
//END_1 OMIT

//START_2 OMIT
type Dimensions struct { // HL2
  Height int
  Width int
}

type Car struct {
  Doors int
  Color string
  Dimensions Dimensions // HL2
}
//END_2 OMIT

//START_3 OMIT
carJson := `{
                "Doors": 2,
                "Color": "red",
                "dimensions":{
                  "height":24,
                  "width":10
                }
             }`

var car Car
json.Unmarshal([]byte(carJson), &car) // HL3

fmt.Printf(car)
//END_3 OMIT
