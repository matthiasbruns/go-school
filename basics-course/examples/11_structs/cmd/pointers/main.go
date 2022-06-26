package main

import "fmt"

//START_1 OMIT
type Car struct { // HL1
	Doors int
	HP    int
}

func main() {
	car := Car{5, 200}
	p := &car // HL1
	p.HP = 350
	fmt.Println(car)
}

//END_1 OMIT
