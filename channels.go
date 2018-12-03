package main

import (
	"fmt"
)

// một cách để các Goroutines communicate với nhau

///////// INIT ///////

func init_chan() {
	var _chan chan int
	if _chan == nil {
		fmt.Println("channel is nil, init it!")
		_chan = make(chan int)
		fmt.Printf("%T", _chan)
	}

	// second way
	_chan2 := make(chan int)
	fmt.Printf("%T", _chan2)
}

///////// READ & WRITE ///////
var _global int

func test_goroutines(test chan int) {
	fmt.Println("test channel")

	// write to channel
	test <- 1

	_global = 22
	fmt.Println("_global", _global)
}

func main() {

	//init_chan()

	test := make(chan int)
	go test_goroutines(test)

	// read from channel
	_receiver_chan := <- test

	fmt.Println("_receiver_chan",_receiver_chan)
	fmt.Println("_global", _global)
	fmt.Println("continue to main thread")

}
