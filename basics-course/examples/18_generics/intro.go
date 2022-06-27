func PrintAnything[T any](thing T) {
	fmt.Println(thing)
}

//...

PrintAnything("Hello!")
PrintAnything(42)
PrintAnything(true)
