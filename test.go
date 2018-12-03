package main

import (
	"fmt"
	_ "image/gif"
	_ "image/png"
)

func init() {
	fmt.Println("sssssss")
}

func main() {
	sum := 0
	array:= []int{1,2,3,4}
	for _, value := range array {
		sum += value
	}
	fmt.Println("aaaaaa")
}