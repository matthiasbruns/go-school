package main

func main() {
	defer func() {
		// Recover from the panic
		if r := recover(); r != nil {
			println("recovered from", r)
		}
	}()

	panicVal := panicInFunction()
	println("panicVal", panicVal)

	// Simulate a panic
	panic("oh no")
}

func panicInFunction() int {
	// if a function returns a value, recover will return the zero value of the return type
	defer func() {
		// Recover from the panic
		if r := recover(); r != nil {
			println("recovered from", r)
		}
	}()

	// Simulate a panic
	panic("oh no")
}
