//START_1 OMIT
{
	"species": "pigeon",
	"description": "likes to perch on rocks"
}
//END_1 OMIT

//START_2 OMIT
type Bird struct {
	Species string
	Description string
}
//END_2 OMIT


//START_3 OMIT
birdJson := `{"species": "pigeon","description": "likes to perch on rocks"}`
var bird Bird	
json.Unmarshal([]byte(birdJson), &bird) // HL3
fmt.Printf("Species: %s, Description: %s", bird.Species, bird.Description)
//END_3 OMIT