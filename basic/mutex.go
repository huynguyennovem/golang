package main

import (
	"fmt"
	"sync"
	_ "sync"
	"time"
)

// if we do not use Mutex to make sure only one goroutines at a time can access the map, we will got error (fatal error: concurrent map writes)
// Because many goroutines are trying to access the map in same time

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
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
	//c.mux.Lock()
	//Lock so only one goroutine at a time can access the map c.v.
	//defer c.mux.Unlock()
	return c.v[key]
}

//////////////////////////

var temp sync.Mutex
var mapTemp = make(map[string]int)

func AnotherMutex(key string) {
	temp.Lock()
	defer temp.Unlock()
	mapTemp[key]++
}

//////////////////////////

func main() {

	////// test 1
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))

	////// test 2
	for i := 0; i < 1000; i++ {
		go AnotherMutex("aaa")
	}
}
