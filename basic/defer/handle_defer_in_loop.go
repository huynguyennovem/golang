package main

import "fmt"

func main() {
	// comment out the func to see the results

	//wrong()
	//solution1()
	solution2()
}

/**
	func() calls inside the loop:
		- Don't get executed until the func ends, not when each step of for loop ends
		- All calls here will eat up the funcs stack and may cause unforeseen problems.
 */
 /**
 	Output:
 	for step #0 ends
	for step #1 ends
	for step #2 ends
	for step #3 ends
	for step #4 ends
	for step #5 ends
	for step #6 ends
	for step #7 ends
	for step #8 ends
	for step #9 ends
	defer runs #10
	defer runs #10
	defer runs #10
	defer runs #10
	defer runs #10
	defer runs #10
	defer runs #10
	defer runs #10
	defer runs #10
	defer runs #10
  */
func wrong() {
	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Printf("defer runs #%d\n", i)
		}()

		// deferred func doesn't run here

		fmt.Printf("for step #%d ends\n", i)
	}

	// it runs here
}

/**
	Output:
	runs #0
	for step #0 ends
	runs #1
	for step #1 ends
	runs #2
	for step #2 ends
	runs #3
	for step #3 ends
	runs #4
	for step #4 ends
	runs #5
	for step #5 ends
	runs #6
	for step #6 ends
	runs #7
	for step #7 ends
	runs #8
	for step #8 ends
	runs #9
	for step #9 ends
 */
func solution1() {
	for i := 0; i < 10; i++ {
		fmt.Printf("runs #%d\n", i)
		// no need for defer

		fmt.Printf("for step #%d ends\n", i)
	}
}

/**
processing #0
process(): defer runs #0
for step #0 ends

processing #1
process(): defer runs #1
for step #1 ends

processing #2
process(): defer runs #2
for step #2 ends

processing #3
process(): defer runs #3
for step #3 ends

processing #4
process(): defer runs #4
for step #4 ends

processing #5
process(): defer runs #5
for step #5 ends

processing #6
process(): defer runs #6
for step #6 ends

processing #7
process(): defer runs #7
for step #7 ends

processing #8
process(): defer runs #8
for step #8 ends

processing #9
process(): defer runs #9
for step #9 ends
 */
// for situations which need defer use this solution
func solution2() {
	for i := 0; i < 10; i++ {
		func() {
			defer func() {
				fmt.Printf("process(): defer runs #%d\n", i)
			}()

			fmt.Printf("processing #%d\n", i)
		}()
		fmt.Printf("for step #%d ends\n", i)
		fmt.Println()
	}
}
