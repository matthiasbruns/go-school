package main

import "fmt"

func main() {
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)
	fmt.Println(i, f, u)

	i = -42
	f = float64(i)
	u = uint(f)
	fmt.Println(i, f, u)
}
