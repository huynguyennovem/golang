package main

import "fmt"

type example struct {
	X, Y string
}

var (
a = example{"Aus","Can"}
b = &example{"Jap","Kor"}
c = example{X:"US",Y:"UK"}
d = example{}
)

func temp1(){
	fmt.Println(a)
	fmt.Println(b)
	e := b
	b.X = "Rus"
	f := *b
	fmt.Println("a:\t",a)
	fmt.Println("b:\t",b)
	fmt.Println("c:\t",c)
	fmt.Println("d:\t",d)
	fmt.Println("e:\t",e)
	fmt.Println("f:\t",f)
}

func temp2() *example{
	return &example{"a","b"}
}

func main() {

	//temp1()


	fmt.Println(*temp2())
	fmt.Println(temp2().X)
	fmt.Println((*temp2()).X)
}