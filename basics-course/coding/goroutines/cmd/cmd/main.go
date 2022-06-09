package main

import (
	"fmt"
	"time"
)

// Create the worker method first that handles the workload
func worker(id int, jobs <-chan int, results chan<- int) { // HL1
	// Keeps waiting for jobs until close(jobs) is called somewhere
	for j := range jobs { // HL1
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)    // HL2
	results := make(chan int, numJobs) // HL2

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results) // HL2
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j // HL3
	}
	close(jobs) // HL3

	for a := 1; a <= numJobs; a++ {
		res := <-results // HL3
		fmt.Println(res)
	}
}
