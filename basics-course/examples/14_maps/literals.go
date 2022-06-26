package main

type Car struct {
	Doors, HP int
}

//START_1 OMIT
var inv = map[string]Car{
	"Fiat Punto":    Car{3, 50},
	"Telsa Model 3": Car{4, 480},
}

//END_1 OMIT

//START_2 OMIT
var inv2 = map[string]Car{
	"Fiat Punto":    {3, 50},
	"Telsa Model 3": {4, 480},
}

//END_2 OMIT
