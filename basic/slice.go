package main

import (
	"fmt"
)

func init1() {
	var number []int
	number = make([]int, 3, 5)
	fmt.Println(len(number))
	fmt.Println(cap(number))
	fmt.Println(number)
	fmt.Println()
	for i := 0; i < len(number); i++ {
		fmt.Println(number[i])
	}
	for _, i := range number {
		fmt.Println(i)
	}

	var number2 []int
	fmt.Println(number2)
	fmt.Printf("%v", number2)
	if number2 == nil {
		fmt.Println("num2 is nil")
	}
}

func _append() {
	var xxx []int
	xxx = append(xxx, 1, 2, 3)
	fmt.Println(xxx)
	fmt.Println(len(xxx))
	fmt.Println(cap(xxx))

	// copy slice
	xxx1 := make([]int, len(xxx), cap(xxx))
	copy(xxx1, xxx)
	fmt.Println(xxx1)

	// edit slice, test cap vs len
	xxx = append(xxx, 4, 5, 6, 7)
	fmt.Println(xxx)
	fmt.Println(len(xxx))
	fmt.Println(cap(xxx))
}

func main() {
	//init1()
	_append()
}
