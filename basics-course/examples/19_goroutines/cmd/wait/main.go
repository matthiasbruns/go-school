package main

import (
	"fmt"
	"sync"
	"time"
)

//START_1 OMIT
func doSomething(wg *sync.WaitGroup) {
	// tell WaitGroup that we are done at the end of this function
	defer wg.Done() // HL1

	time.Sleep(100 * time.Millisecond) // simulate work
	fmt.Println("doSomething done")
}

func main() {
	wg := sync.WaitGroup{} // HL1

	wg.Add(1) // HL1
	go doSomething(&wg)

	// block main goroutine until doSomething is done
	wg.Wait() // HL1

	fmt.Println("main done")
}

//END_1 OMIT
