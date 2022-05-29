package main

import "fmt"

//START_1 OMIT
type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v // HL1
	p.X = 1e9
	fmt.Println(v)
}

//END_1 OMIT
