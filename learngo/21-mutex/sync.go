package main

import ( ."fmt"; "sync"; ."time" )

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v map[string]int
	mux sync.Mutex
}

func main(){
	mutualexclusion()
}

func mutualexclusion() {
	Println("-- Mutual Exclusion")
	
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ { go c.Inc("somekey") }

	Sleep(Second)
	Println(c.Value("somekey"))
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mux.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.Unlock()
	return c.v[key]
}
