package main

import (
	"errors"
	"fmt"
)

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	// you can return multiple values from a function
	divResult, err := div(10, 2)
	if err != nil {
		fmt.Println(err)
		return
	}

	// you can use _ to ignore the error
	divResult, _ = div(10, 0)
	fmt.Println(divResult)
}
