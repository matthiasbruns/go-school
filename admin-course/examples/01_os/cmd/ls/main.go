package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

//START_1 OMIT
func execute() {
	// Add arguments to the command call - in this case "-ltr"
	out, err := exec.Command("ls", "-ltr").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println("Command Successfully Executed")
	output := string(out[:])
	fmt.Println(output)
}

//END_1 OMIT

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
	} else {
		execute()
	}
}
