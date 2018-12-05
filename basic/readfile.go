package main

import (
	"fmt"
	"github.com/gobuffalo/packr"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))

	box := packr.NewBox("../packbox")


}
