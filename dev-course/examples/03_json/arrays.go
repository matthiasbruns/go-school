//START_1 OMIT
[
  {
    "Doors": 2,
    "Color": "Red"
  },
  {
    "Doors": 5,
    "Color": "Blue"
  }
]
//END_1 OMIT

//START_2 OMIT
carsJson := `[{"doors":2,"color":"Red"},{"doors":5,"color":"Blue"}]`
var cars []Cars
json.Unmarshal([]byte(carsJson), &cars) // HL2
fmt.Printf("Cars : %+v", cars)
//END_2 OMIT
