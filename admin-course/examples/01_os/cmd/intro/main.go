package main

import (
	"fmt"
	"log"
	"os"
)

//START_1 OMIT
func main() {
	file, err := os.Open("file.txt") // HL1
	if err != nil {
		log.Fatal(err)
	}

	// no err in this example

	data := make([]byte, 100)
	count, err := file.Read(data) // HL1
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, data[:count])
}

//END_1 OMIT
