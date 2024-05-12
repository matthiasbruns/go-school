package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Go runs on:")

	// without os as an inline variable
	switch os := runtime.GOOS; {
	case os == "darwin":
		fmt.Println("OSX.")
	case os == "linux":
		fmt.Println("Linux.")
	case os == "windows":
		fmt.Println("Windows.")
	default:
		fmt.Printf("%s.\n", os)
	}

	// with os as an inline variable
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OSX.")
	case "linux":
		fmt.Println("Linux.")
	case "windows":
		fmt.Println("Windows.")
	default:
		fmt.Printf("%s.\n", os)
	}
}
