type List[T any] struct {
	next *List[T] // HL1
	val  T
}