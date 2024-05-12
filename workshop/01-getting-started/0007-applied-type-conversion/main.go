package main

func main() {
	i := 0

	// any = interface{} - a type that can hold any value
	var j any = i

	// typecast j to int
	k, ok := j.(int)
	if ok {
		println(k)
	}

	// typecast j to string
	l, ok := j.(string)
	if ok {
		println(l)
	}
}
