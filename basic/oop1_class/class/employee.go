package class

import (
	"fmt"
)

type Employee struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}
func New(firstName string, lastName string, totalLeave int, leavesTaken int) Employee {
	e := Employee {firstName, lastName, totalLeave, leavesTaken}
	return e
}
func (e Employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}
