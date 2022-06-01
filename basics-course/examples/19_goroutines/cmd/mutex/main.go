package main

import (
	"fmt"
	"sync"
	"time"
)

//START_1 OMIT
type SafeCounter struct {
	mu sync.Mutex // HL1
	v  map[string]int
}

//END_1 OMIT

//START_2 OMIT
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock() // HL2
	c.v[key]++
	c.mu.Unlock() // HL2
}

func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()         // HL2
	defer c.mu.Unlock() // HL2
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}

//END_2 OMIT
