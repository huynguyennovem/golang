package oop1_class

import "golang/basic/oop1_class/class"

func main() {
	//e := class.Employee {
	//	FirstName: "Sam",
	//	LastName: "Adolf",
	//	TotalLeaves: 30,
	//	LeavesTaken: 20,
	//}
	//e.LeavesRemaining()

	e := class.New("Sam", "Adolf", 30, 20)
	e.LeavesRemaining()
}
