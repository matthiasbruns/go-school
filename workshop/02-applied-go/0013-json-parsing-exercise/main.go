package main

import (
	"fmt"

	"github.com/matthiasbruns/go-school/workshop/02-applied-go/0013-json-parsing-exercise/api"
)

func main() {
	meal, err := api.GetRandomMeal()
	if err != nil {
		panic(err)
	}

	fmt.Println(meal)
}
