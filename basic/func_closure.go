package main

import "fmt"

func getSequence() func() int{
	i:=0
	return func() int {
		i+=1
		return i
	}
}

func main() {
	xxx := getSequence()
	fmt.Println(xxx())
}
