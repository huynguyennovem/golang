package main

import (
	"encoding/json"
	"fmt"
)

// omitempty: muc dich la de cac field empty value se khong duoc output
// https://www.sohamkamani.com/blog/golang/2018-07-19-golang-omitempty/

type dimension struct {
	Height int
	Width int
}

type Dog struct {
	Breed    string
	// Field appears in JSON as empty key and
	// the field is omitted from the object if its value is empty
	WeightKg *int  `json:",omitempty"`
	Size *dimension `json:",omitempty"`
}

/*
	2 case dac biet khi su dung omitempty:
	case 1:  Size *dimension: 	ap dung TH dac biet field co type la 1 struct (can phai la 1 pointer)
	case 2:  WeightKg *int: 	TH khi WeightKg = 0 thi omitempty se khong loai bo field khoi output (can phai la 1 pointer)
*/
func main() {
	x:=0
	d := Dog{
		Breed:    "dalmation",
		WeightKg:	&x,
	}
	b, _ := json.Marshal(d)
	fmt.Println(string(b))
}