select {
case i := <-c:
	// use i
default: // HL1
	// receiving from c would block
}