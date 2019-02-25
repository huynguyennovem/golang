package main

import "fmt"

type database struct{}

func (db *database) connect() (disconnect func()) {
	fmt.Println("connect")
	return func() {
		fmt.Println("disconnect")
	}
}

func fail() {
	db := &database{}
	defer db.connect()

	fmt.Println("query db...")
}

func solution_1() {
	db := &database{}
	close := db.connect()
	defer close()
	fmt.Println("query db...")
}

func solution_2() {
	db := &database{}
	defer db.connect()()
	fmt.Println("query db...")
}

func main() {
	//fail()
	solution_1()
	//solution_2()
}
