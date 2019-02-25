package main

import "fmt"

func main() {
	var run func() = nil
	defer run()
	fmt.Println("runs")
}
