####1. Definition<br/>
- are named collections of method signatures
- interface defines the behaviour of an object

####2. Usage
- Example 1:
```go
package main
import "fmt"

// define interface
type Person interface {
    cal_age() int
}

// init an object
type Student struct {
    age int
}
// implement interface
func (student Student) cal_age() int{
    return student.age * 2
}

func main() {
    var st = Student{33}

    // normally, we can do this
    fmt.Println(st.cal_age())

    // use interface
    var p Person
    p = st
    fmt.Println(p.cal_age())
}

```

- Example 2:
```go
package main

import (  
    "fmt"
)

type SalaryCalculator interface {  
    CalculateSalary() int
}

type Permanent struct {  
    empId    int
    basicpay int
    pf       int
}

type Contract struct {  
    empId  int
    basicpay int
}

class
func (p Permanent) CalculateSalary() int {  
    return p.basicpay + p.pf
}

class
func (c Contract) CalculateSalary() int {  
    return c.basicpay
}

/*
total expense is calculated by iterating though the SalaryCalculator slice and summing  
the salaries of the individual employees  
*/
func totalExpense(s []SalaryCalculator) {  
    expense := 0
    for _, v := range s {
        expense = expense + v.CalculateSalary()
    }
    fmt.Printf("Total Expense Per Month $%d", expense)
}

func main() {  
    pemp1 := Permanent{1, 5000, 20}
    pemp2 := Permanent{2, 6000, 30}
    cemp1 := Contract{3, 3000}
    employees := []SalaryCalculator{pemp1, pemp2, cemp1}
    totalExpense(employees)

}
```
- Example 3: Empty interface
```go
package main

import (  
    "fmt"
)

func describe(i interface{}) {  
    fmt.Printf("Type = %T, value = %v\n", i, i)
}

func main() {  
    s := "Hello World"
    describe(s)
    i := 55
    describe(i)
    strt := struct {
        name string
    }{
        name: "Naveen R",
    }
    describe(strt)
}
```

- Example 4: Assertion
```go
package main

import (  
    "fmt"
)

func assert(i interface{}) {  
    v, ok := i.(int)
    fmt.Println(v, ok)
}
func main() {  
    var s interface{} = 56
    assert(s)
    var i interface{} = "Steven Paul"
    assert(i)
}
```