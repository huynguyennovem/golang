package main

import (
	"fmt"
	"math"
)

// this test to prove variable is loaded before init() func.

var WhatIsThe = AnswerToLife()

func AnswerToLife() int {
	fmt.Println("AnswerToLife()")
	return 42
}

func init() {
	fmt.Println("init")
	WhatIsThe = 0
}

////////////////////////////


func fun1() float64{
	temp := func(xxx float64) float64 {
		return math.Sqrt(xxx)
	}
	fmt.Println("load before init")
	return temp(9)
}

var hihi = fun1()

func main() {
	if WhatIsThe == 0 {
		fmt.Println("It's all a lie.")
	}

}


