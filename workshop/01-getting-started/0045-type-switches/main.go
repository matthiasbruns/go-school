package main

func main() {
	do(21)
	do("hello")
	do(true)
	do(21.0)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		println("int", v)
	case string:
		println("string", v)
	case bool:
		println("bool", v)
	default:
		println("unknown")
	}
}
