//START_1 OMIT
type Vehicle interface { // HL1
	Drive() string
}

//END_1 OMIT

//START_2 OMIT
func (v *Car) Drive() string { // HL2
	if c.EngineRunning {
		return "Wroooom!"
	} else {
		return "No wrooom!"
	}
}

//END_2 OMIT

//START_3 OMIT

func main() {
	var v Vehicle
	c := Car{Doors: 3, HP: 120}

	v = &c // a *Car implements Vehicle
	v.Drive()
}

//END_3 OMIT
