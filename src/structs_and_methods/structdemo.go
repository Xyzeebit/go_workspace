package main

import (
	"fmt"
	"time"
)

type Employee struct {
	id int
	name string
	country string
	created time.Time
}

func main() {
	var emp Employee
	newEmp := new(Employee);
	emp.id = 1;
	emp.name = "Employee 1"
	emp.country = "DE";
	emp.created = time.Now()

	newEmp.id = 5;
	newEmp.name = "Employee 5"
	newEmp.country = "UK"
	newEmp.created = time.Now();

	// display
	fmt.Println("=====================")
	fmt.Println(emp.id)
	fmt.Println(emp.name)
	fmt.Println(emp.country)
	fmt.Println(emp.created)
	fmt.Println("=====================")
	fmt.Println(newEmp.id)
	fmt.Println(newEmp.name)
	fmt.Println(newEmp.country)
	fmt.Println(newEmp.created)
}