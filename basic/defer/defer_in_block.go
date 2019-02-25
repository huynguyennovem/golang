package main

import "fmt"

// DEFER only executes after the containing func ends.
// DEFER belongs to a func not to a block.

/**
	Output:
	block: ends
	main: ends
	block: defer runs
 */

func example() {
	{
		defer func() {
			fmt.Println("block: defer runs")
		}()
		fmt.Println("block: ends")
	}
	fmt.Println("main: ends")
}

/**
	Output:
	block: ends
	block: defer runs
	main: ends
 */
func solution() {
	func() {
		defer func() {
			fmt.Println("block: defer runs")
		}()
		fmt.Println("block: ends")
	}()
	fmt.Println("main: ends")
}

func main() {
	// example()
	solution()
}
