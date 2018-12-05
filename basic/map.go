package main

import "fmt"

func main() {
	var _map map[string] int
	fmt.Println(_map)
	_map = make(map[string] int)

	// assign pair
	_map["sss"] = 5
	_map["vvv"] = 1
	fmt.Println(_map)

	// check pair exist on map
	a, b := _map["goroutines_channel"]
	fmt.Println(a, b)

	// delete
	delete(_map, "sss")
	delete(_map, "aaa")
	fmt.Println(_map)

}
