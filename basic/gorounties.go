package main

import (
	"fmt"
	"time"
)

func test1(){
	fmt.Println("111")
}

func test2(){
	fmt.Println("222")
}

func main() {
	go test1()
	go test1()
	go test1()
	go test1()
	time.Sleep(time.Millisecond * 1)
	fmt.Println("main")
	go test2()
	go test2()
	go test2()
	go test2()
	time.Sleep(time.Millisecond * 1)
}
