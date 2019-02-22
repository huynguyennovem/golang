package main

import "fmt"

func undirectional_chan(sendchan chan<- int) {
	sendchan <- 10
}

func main() {

	sendch := make(chan int)
	go undirectional_chan(sendch)
	fmt.Println(<-sendch)

	// or

	//a := <-sendch
	//fmt.Println(a)

}
