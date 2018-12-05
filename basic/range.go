package main

import "fmt"

type post struct {
	title   string
	content string
}

func (p post) details() {
	fmt.Println("Title: ", p.title)
	fmt.Println("Content: ", p.content)
}

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	for i := range numbers {
		fmt.Println(numbers[i])
	}

	fmt.Println()

	_map := map[string]int{"a": 1, "b": 2}
	for i := range _map {
		fmt.Println(i, _map[i])
	}

	fmt.Println()

	for i, j := range _map {
		fmt.Println(i, j)
	}

	fmt.Println()

	// range with _
	post1 := post{
		"Inheritance in Go",
		"Go supports composition instead of inheritance",
	}
	post2 := post{
		"Struct instead of Classes in Go",
		"Go does not support classes but methods can be added to structs",
	}
	post3 := post{
		"Concurrency",
		"Go is a concurrent language and not a parallel one",
	}
	posts := []post{post1, post2, post3}
	for _, p := range posts { // in this case _ is index, p is value
		p.details()
		fmt.Println()
	}
}
