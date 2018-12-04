package main

import "hello/oop/employee"

func main() {
	e := employee.Employee {
		FirstName: "Sam",
		LastName: "Adolf",
		TotalLeaves: 30,
		LeavesTaken: 20,
	}
	e.LeavesRemaining()
}