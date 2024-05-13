package main

import (
	"flag"
	"fmt"
)

func main() {
	wordFlag := flag.String("word", "world", "word to print")

	flag.Parse()

	fmt.Println("Hello", *wordFlag)
}
