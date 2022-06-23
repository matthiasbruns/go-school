package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Welcome to my awesome 'ls' program written in Golang")

	var dirPath string

	if len(os.Args) > 1 {
		dirPath = os.Args[1]
		fmt.Println("Scanning provided directory", dirPath)
	} else {
		// Path to current work dir
		if d, err := os.Getwd(); err != nil {
			log.Fatalln(err)
		} else {
			dirPath = d
		}
		fmt.Println("Scanning current directory", dirPath)
	}

	// File pointer to work dir
	dir, err := os.Open(dirPath)
	if err != nil {
		log.Fatalln(err)
	}

	// Get filed from dir
	files, err := dir.Readdir(0)
	if err != nil {
		log.Fatalln(err)
	}

	// Print file contents
	fmt.Printf("%d file(s) in directory %s\n", len(files), dirPath)
	fmt.Printf("  Mode\t\t| Size\t| Name\n")
	for _, v := range files {
		fmt.Println("  ", v.Mode(), "\t", v.Size(), "\t", v.Name())
	}
}
