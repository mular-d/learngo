package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DOB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

func main() {

	var dilbert Employee
	dilbert.Salary -= 5000

	fmt.Println(EmployeeByID(dilbert.ManagerID).Position)

	id := dilbert.ID
	EmployeeByID(id).Salary = 0
}

func EmployeeByID(id int) *Employee {
	return &Employee{1, "M", "A.A", time.Now(), "Mr.", 12, 23}
}
