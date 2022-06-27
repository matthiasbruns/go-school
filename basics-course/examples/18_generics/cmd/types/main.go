package main

import "fmt"

type List[T any] struct {
	next *List[T] // HL1
	val  T
}

func main() {
	l := List[int]{next: nil, val: 1337}
	fmt.Println(l.val)
}
