package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.Unlock()
}

//Decrease the counter for the given key.
func (c *SafeCounter) Dec(key string){
	c.Lock()
	c.v[key]--
	c.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
		//go c.Dec("somekey")
	}
	
	for i:=0;i<500;i++{
		go c.Dec("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
