package main

type Vertex struct {
	Lat, Long float64
}

//START_1 OMIT
var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

//END_1 OMIT

//START_2 OMIT
var m2 = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

//END_2 OMIT
