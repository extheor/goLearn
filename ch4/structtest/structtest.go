package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	Dob       time.Time
	Position  string
	Salary    int
	ManagerID int
}

var dilbert Employee

func main() {
	// fmt.Println(dilbert)

	fmt.Println(EmployeeByID(dilbert.ID).Position)

	id := dilbert.ID
	EmployeeByID(id).Salary = 0
}

func EmployeeByID(id int) *Employee {
	return &dilbert
}
