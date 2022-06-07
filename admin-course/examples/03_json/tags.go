//START_1 OMIT
{
	"birdType": "pigeon",
	"what it does": "likes to perch on rocks" // HL1
}
//END_1 OMIT

//START_2 OMIT
type Bird struct {
	Species string `json:"birdType"`
	Description string `json:"what it does"` // HL2
}
//END_2 OMIT