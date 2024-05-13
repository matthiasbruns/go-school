package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func ls() {
	out, err := exec.Command("ls", "-l").Output()
	if err != nil {
		panic(err)
	}

	fmt.Println("command output:", string(out))
}

func main() {
	switch runtime.GOOS {
	case "windows":
		fmt.Println("Can't run ls command on Windows")
	default:
		ls()
	}
}
