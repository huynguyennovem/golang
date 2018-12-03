package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 1
	fmt.Println("capacity is", cap(ch))
	fmt.Println("length is", len(ch))
}