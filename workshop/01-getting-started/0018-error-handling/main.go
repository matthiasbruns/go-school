package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := "_e42"

	i, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("Couldn't convert number: %v\n", err)
		return
	}

	fmt.Println("Converted integer:", i)
}
