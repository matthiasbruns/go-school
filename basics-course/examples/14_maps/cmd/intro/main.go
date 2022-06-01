package main

import "fmt"

//START_1 OMIT
type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex // HL1

func main() {
	m = make(map[string]Vertex) // HL1
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

//END_1 OMIT
