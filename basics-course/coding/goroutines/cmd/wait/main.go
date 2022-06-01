package main

import (
	"fmt"
	"sync"
	"time"
)

func doSomething(wg *sync.WaitGroup) {
	defer wg.Done() // tell WaitGroup that we are done at the end of this function

	time.Sleep(100 * time.Millisecond) // simulate work
	fmt.Println("doSomething done")
}

func main() {
	wg := sync.WaitGroup{}
	
	wg.Add(1)
	go doSomething(&wg)
	wg.Wait() // block main goroutine until doSomething is done

	fmt.Println("main done")
}
