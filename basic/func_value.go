package main

import (
	"fmt"
	"math"
)

func main() {
	temp := func(xxx float64) float64 {
		return math.Sqrt(xxx)
	}
	fmt.Println(temp(9))
}
