package main

import "fmt"

type Car struct {
	Doors, HP     int
	EngineRunning bool
}

//START_1 OMIT
func (p Car) String() string { // HL1
	running := "not running"
	if p.EngineRunning {
		running = "running"
	}
	return fmt.Sprintf("%d Doors (%d HP) is %s", p.Doors, p.HP, running) // HL1
}

func main() {
	bmw := Car{Doors: 5, HP: 120, EngineRunning: true}
	fiat := Car{Doors: 3, HP: 80, EngineRunning: false}
	fmt.Println("bmw", bmw, "\nfiat:", fiat) // HL1
}

//END_1 OMIT
