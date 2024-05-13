package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./file.txt")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	// Do something with the file
	data := make([]byte, 100)
	_, err = file.Read(data)
	if err != nil {
		panic(err)
	}

	// Do something with the data
	fmt.Println(string(data))
}
