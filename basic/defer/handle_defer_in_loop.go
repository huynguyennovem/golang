package _defer

import "fmt"

func main() {
	// comment out the func to see the results

	// wrong()
	// solution1()
	solution2()
}

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

func solution1() {
	for i := 0; i < 10; i++ {
		fmt.Printf("runs #%d\n", i)
		// no need for defer

		fmt.Printf("for step #%d ends\n", i)
	}
}

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
	}
}
