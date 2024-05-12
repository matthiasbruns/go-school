package main

import "fmt"

type List[T any] struct {
	value T
	next  *List[T]
}

func (l List[T]) String() string {
	if l.next == nil {
		return fmt.Sprintf("%v", l.value)
	}

	return fmt.Sprintf("%v", l.value) + " -> " + l.next.String()
}

func main() {
	l := List[int]{value: 420,
		next: &List[int]{value: 1337, next: &List[int]{value: 69}},
	}
	fmt.Println(l)

	l2 := List[string]{value: "hello",
		next: &List[string]{value: "world", next: &List[string]{value: "!"}},
	}
	fmt.Println(l2)
}
