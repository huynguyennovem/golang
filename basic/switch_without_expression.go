package main

import (
	"fmt"
	"time"
)

/**
	replace for if-else
	expression will be in each cases
 */

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	x := 12
	switch {
	case x > 5:
		fmt.Println("x")
	case x > 6:
		fmt.Println("y")
	}
}
