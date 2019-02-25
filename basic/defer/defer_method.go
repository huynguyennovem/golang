package main

import "fmt"

type Car struct {
	model string
}

// POINTER
func (c *Car) PrintModel1() {
	fmt.Println(c.model)
}

// VALUE
func (c Car) PrintModel2() {
	fmt.Println(c.model)
}

func main() {

	// test with pointer OR value

	// POINTER
	c := Car{model: "DeLorean DMC-12"}
	defer c.PrintModel1()
	c.model = "Chevrolet Impala"


	// VALUE
	//c2 := Car{model: "DeLorean DMC-12"}
	//defer c2.PrintModel2()					// c2 will be copied when executes fmt.Println(c.model)
	//c2.model = "Chevrolet Impala"
}
