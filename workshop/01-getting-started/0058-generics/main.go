package main

import (
	"fmt"
)

func JoinToString[T any](s []T, delimiter string) string {
	var result string
	for i, v := range s {
		if i > 0 {
			result += delimiter
		}
		result += fmt.Sprint(v)
	}
	return result
}

func main() {
	fmt.Println(JoinToString([]int{1, 2, 3, 4, 5}, ", "))
	fmt.Println(JoinToString([]string{"a", "b", "c", "d", "e"}, "🥳 "))
	fmt.Println(JoinToString([]float64{1.1, 2.2, 3.3, 4.4, 5.5}, "# "))
}
