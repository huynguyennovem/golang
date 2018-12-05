package main

import (
	"fmt"
)

// this aim to check status of channel is closed or opened

func goroutines_channel(chn chan int) {
	for i := 0; i < 10; i++ {
		chn <- i
	}
	fmt.Println("closed 1")

	close(chn)

	// sau khi close channel, vẫn còn 1 khoản time channel run tiếp
	// time này ko cố định và cho dù là infi loop hay range loop
	for i := 0; i < 10000; i++ {
		fmt.Println("closed 2", i)
	}
}

func for_infi_loop() {
	chn := make(chan int)
	go goroutines_channel(chn)
	for {
		output, ok := <-chn
		if ok == false {
			fmt.Println("close channel")
			break
		}
		fmt.Println("ok ok ", output, ok)
	}
}

func for_range_loop() {
	chn := make(chan int)
	go goroutines_channel(chn)
	for output := range chn {
		fmt.Println("ok ok ", output)
	}
}

////////////////// multiple channels //////////////////////

func alphabets(chn chan rune) {
	for i := 'a'; i <= 'e'; i++ {
		chn <- i
	}
	close(chn)
}

func numbers(chn chan int) {
	for i := 1; i <= 5; i++ {
		chn <- i
	}
	close(chn)
}

func merge_alphabets(out_chn chan string) {
	var out string
	chn1 := make(chan rune)
	go alphabets(chn1)
	for i := range chn1 {
		out += string(i)
	}
	out_chn <- out
	close(out_chn)
}

func add_numbers(out_chn chan int) {
	var out int
	chn1 := make(chan int)
	go numbers(chn1)
	for i := range chn1 {
		out += i
	}
	out_chn <- out
	close(out_chn)
}

func intergrate_channel() {
	chn1 := make(chan string)
	chn2 := make(chan int)
	go merge_alphabets(chn1)
	go add_numbers(chn2)

	a, b := <-chn1, <-chn2
	fmt.Println(a, b)
}

func main() {

	//for_infi_loop()

	//for_range_loop()

	intergrate_channel()

}
