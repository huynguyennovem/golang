package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	wordPtr := flag.String("a", "a value", "a string")

	var svar string
	flag.StringVar(&svar, "b", "b value", "a string var")

	flag.Parse()

	fmt.Println("---a:", *wordPtr)
	fmt.Println("---b:", svar)
	fmt.Println("---flag args: ", flag.Args()) // list tham so flag
	fmt.Println()
	fmt.Println(os.Args) // list tham so khi truyen lenh, bao gom ./flag duoc coi nhu 1 tham so
	fmt.Println("os args len: ", len(os.Args))
	for i, arg := range os.Args {
		fmt.Println(i, ":\t", arg)
	}
}
