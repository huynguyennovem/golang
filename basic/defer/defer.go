package _defer

import "fmt"

// duoc thuc thi truoc khi func return (exit) hoac khi panic() duoc call (dc execute truoc panic)
// usage: nen su dung khi cac case co default final action da biet truoc. eg: close db after access (by db always need to close)

type person struct {
	firstName string
	lastName  string
}

func (p person) fullName() {
	fmt.Printf("%s %s", p.firstName, p.lastName)
}

func base() {
	p := person{
		firstName: "John",
		lastName:  "Smith",
	}
	defer p.fullName()
	fmt.Printf("Welcome ")
}

///////////////////////////////

func printA(a int) {
	fmt.Println("value of a in deferred function", a)
}
func var_value() {
	a := 5
	defer printA(a)
	a = 10
	fmt.Println("value of a before deferred function call", a)

}

// Last In First Out (LIFO)
func multi_defer() {
	name := "Naveen"
	fmt.Printf("Orignal String: %s\n", string(name))
	fmt.Printf("Reversed String: ")
	for _, v := range []rune(name) {
		defer fmt.Printf("%c", v)
	}
}

func main() {
	base()

	//var_value()

	//multi_defer()
}
