package main

import (
	"fmt"
	"strconv"
)

func main() {
	var asInt = 2
	var asFloat = float64(asInt)
	fmt.Println("var asFloat = float64(asInt)\t\t", asFloat)
	var asString = strconv.Itoa(asInt)
	fmt.Println("var asString = strconv.Itoa(asInt)\t", asString)

	var j interface{} = asInt
	s, ok := j.(string)
	fmt.Println("s, ok := j.(string)\t\t\t", "s", s, "ok", ok)
}
