package class_access

import (
	"fmt"
	"golang/advance/class_access/class_a_nested"
)
var a int
var A int
func BBB() {
	fmt.Println("BBB")
	class_a_nested.DDD()
}
