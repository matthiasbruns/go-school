package main

import (
	"fmt"
	"runtime"
)

//START_1 OMIT
func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os { // HL1
	case "darwin": // HL1
		fmt.Println("OS X.")
	case "linux": // HL1
		fmt.Println("Linux.")
	default: // HL1
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

//END_1 OMIT
