package main

import (
	"fmt"
)

type Car struct {
	Doors, HP     int
	EngineRunning bool
}

//START_1 OMIT
func (c Car) Drive() string { // HL1
	if c.EngineRunning {
		return "Wroooom!"
	} else {
		return "No wrooom!"
	}
}

func (c *Car) Start() { // HL1
	c.EngineRunning = true
}

func main() {
	v := Car{Doors: 3, HP: 120}
	fmt.Println(v.Drive())
	v.Start()
	fmt.Println(v.Drive())
}

//END_1 OMIT
