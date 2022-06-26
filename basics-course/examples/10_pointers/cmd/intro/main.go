package main

//START_1 OMIT
func main() {
	a := 5
	aPtr := &a
	aPtrPtr := &aPtr

	println(a)       // 5
	println(aPtr)    // 0x01
	println(aPtrPtr) // 0x02
}

//END_1 OMIT
