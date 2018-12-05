package main

import "fmt"

// Khi một goroutines write hoặc read ma khong co channel đích - receiver thi sẽ gay ra deadlock (lỗi runtime).

func deadlock(){
	chn := make(chan int)
	chn <- 33
	fmt.Println(<-chn)
}

func fix_deadlock(){
	chn := make(chan int, 3)
	chn <- 33
	chn <- 22
	fmt.Println(<-chn)
	fmt.Println(<-chn)
}

func main() {

	//deadlock()

	fix_deadlock()
}
