package main

import "fmt"


// interface declare
type Person interface {
	cal_age() int
}



// init Student and it's method
type Student struct {
	age int
}
// implement interface's method cal_age()
func (student Student) cal_age() int{
	return student.age * 2
}

// function for main call
func get_age(person Person) int{
	return person.cal_age()
}
func main() {
	var st = Student{33}

	fmt.Println(st.cal_age())

	fmt.Println(get_age(st))

	var p Person
	p = st
	fmt.Println(p.cal_age())

}
