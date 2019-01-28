package main

import (
	"fmt"
	"math"
)

func anonymous_func() {
	func(n string) {
		fmt.Println("Welcome", n)
	}("Gophers")
}

///////////////////////

type add func(a int, b int) int

func func_type() {
	var a add = func(x int, y int) int {
		return x + y
	}
	b := a(1, 2)
	fmt.Println(b)

	// another way
	var x = add(func(x int, y int) int {
		return x + y
	})
	fmt.Println(x(3, 4))
}

///////////////////////

func as_argument(a func(a, b int) int) {
	fmt.Println(a(1, 2))
}
func func_as_argument() {
	a := func(a, b int) int {
		return a + b
	}
	as_argument(a)
}

/////////////////////////

func simple() func(a, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	return f
}

func return_func() {
	s := simple()
	fmt.Println(s(60, 7))
}

func closure_func() {
	a := 5
	func() {
		fmt.Println("a =", a)
	}()
}

func func_value() {
	temp := func(xxx float64) float64 {
		return math.Sqrt(xxx)
	}
	fmt.Println(temp(9))
}

//////////////////////////

func main() {
	//anonymous_func()

	//func_type()

	//func_as_argument()

	//return_func()

	//closure_func()

	//func_value()

}
