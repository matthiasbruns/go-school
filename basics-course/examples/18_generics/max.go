//START_1 OMIT
func Max[E comparable](input []E) (max E) {
	for _, v := range input { // HL1
		if v > max {
			max = v
		}
	}
	return max
}

//..

fmt.Println(Max([]int{1, 2, 3}))
// cannot compare v > max (operator > not defined for E)
//END_1 OMIT

//START_2 OMIT
type Ordered interface {
    ~int | ~int8 | ~int16 ... // and so on
}
//END_2 OMIT
