package main

import "fmt"

// this test proves the multiple functions of interface
// (as return type / as func's param)

type EventBroadcaster interface {

}

type eventBroadcasterImpl struct {
	var1 int
	var2 string
}

/**
	interface as return type
 */
func NewBroadcaster() EventBroadcaster {
	return &eventBroadcasterImpl{0, "xxx"}
}

/**
	interface as parameter
 */
func NewBroadcaster2(broadcaster EventBroadcaster) {
	fmt.Println(broadcaster)
}

func main() {

	// as return type
	broad := NewBroadcaster()
	fmt.Println(broad)

	// as func param
	NewBroadcaster2(&eventBroadcasterImpl{10, "yyy"})
}
