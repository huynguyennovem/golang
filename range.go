package main

import "fmt"

func main() {
	numbers := []int{0,1,2,3,4,5,6,7,8}
	for i:= range numbers{
		fmt.Println(numbers[i])
	}

	fmt.Println()

	_map := map[string] int{"a":1, "b":2}
	for i:= range _map{
		fmt.Println(i, _map[i])
	}

	fmt.Println()

	for i, j := range _map{
		fmt.Println(i, j)
	}
}
