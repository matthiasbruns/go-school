package main

import "fmt"

type Vehicle interface {
	Drive() string
}

type Car struct {
	Doors, HP     int
	EngineRunning bool
}

//START_2 OMIT
//START_1 OMIT
func (c *Car) Drive() string {
	if c == nil { // HL1
		fmt.Println("Drive(): <nil>")
		return ""
	} else if c.EngineRunning {
		return "Wroooom!"
	} else {
		return "No wrooom!"
	}
}

//END_1 OMIT
func main() {
	var v Vehicle // interface
	var c *Car    // implementing struct
	v = c
	fmt.Printf("Vehicle: (%v)\n", v)
	v.Drive() // HL1

	v = &Car{Doors: 2, HP: 100} // implementing struct
	fmt.Printf("Vehicle: (%v)\n", v)
	v.Drive()
}

//END_2 OMIT
