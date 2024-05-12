package main

import "fmt"

func main() {
	var number int
	var name string
	var isTrue bool
	var floatNumber float32
	var complexNumber complex64
	var array []int
	var slice []int
	var mapType map[string]int
	var channel chan int
	var function func() int
	var pointer *int
	var structType struct {
		name string
		age  int
	}
	var interfaceType interface{}

	fmt.Println("number: ", number)
	fmt.Printf("name: %q\n", name)
	fmt.Println("isTrue: ", isTrue)
	fmt.Println("floatNumber: ", floatNumber)
	fmt.Println("complexNumber: ", complexNumber)
	fmt.Println("array: ", array)
	fmt.Println("slice: ", slice)
	fmt.Println("mapType: ", mapType)
	fmt.Println("channel: ", channel)
	fmt.Println("function: ", function)
	fmt.Println("pointer: ", pointer)
	fmt.Println("structType: ", structType)
	fmt.Println("interfaceType: ", interfaceType)
}
