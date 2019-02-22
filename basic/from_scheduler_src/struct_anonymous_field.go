package main

import (
	"fmt"
)


// anonymous field: string (has no variable name)
type eventBroadcasterImpl1 struct {
	string
}

func main() {
	obj := eventBroadcasterImpl1{"Hello World"}
	fmt.Println(obj.string)
}
