package main

import "fmt"

//START_1 OMIT
func main() {
	var models []string // HL1

	fmt.Println("models:", models, "\nlen(models):", len(models), "\ncap(models):", cap(models))

	if models == nil {
		fmt.Println("models is nil!")
	}
}

//END_1 OMIT
