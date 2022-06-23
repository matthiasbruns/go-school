package main

import (
	"fmt"
	"os"
)

//START_1 OMIT
func main() {
	if len(os.Args) > 1 {
		fmt.Println("Hello", os.Args[1])
	} else {
		fmt.Println("Hello world!")
	}
}

//END_1 OMIT
