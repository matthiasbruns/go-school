//START_1 OMIT
[
  {
    "species": "pigeon",
    "decription": "likes to perch on rocks"
  },
  {
    "species":"eagle",
    "description":"bird of prey"
  }
]
//END_1 OMIT

//START_2 OMIT
birdJson := `[{"species":"pigeon","decription":"likes to perch on rocks"},{"species":"eagle","description":"bird of prey"}]`
var birds []Bird
json.Unmarshal([]byte(birdJson), &birds) // HL2
fmt.Printf("Birds : %+v", birds)
//Birds : [{Species:pigeon Description:} {Species:eagle Description:bird of prey}]
//END_2 OMIT