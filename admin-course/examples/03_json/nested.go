//START_1 OMIT
{
  "species": "pigeon",
  "decription": "likes to perch on rocks"
  "dimensions": {
    "height": 24,
    "width": 10
  }
}
//END_1 OMIT

//START_2 OMIT
type Dimensions struct {
  Height int
  Width int
}

type Bird struct {
  Species string
  Description string
  Dimensions Dimensions // HL2
}
//END_2 OMIT

//START_3 OMIT
birdJson := `{
                "species":"pigeon",
                "description":"likes to perch on rocks",
                "dimensions":{
                  "height":24,
                  "width":10
                }
             }`

var birds Bird
json.Unmarshal([]byte(birdJson), &birds) // HL3

fmt.Printf(bird)
// {pigeon likes to perch on rocks {24 10}}
//END_3 OMIT