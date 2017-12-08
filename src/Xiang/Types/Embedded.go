package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	EmpNum int64
	Dept   string
}

type FairHREmployee struct {
	Person // Is-a relationship. A FairHR employee is a person and is an employee.
	Employee
	// Person Person // Has-a relationship.
}

func (p Person) PrintName() {
	if p != (Person{}) {
		fmt.Printf("The FairHR's employee's name is %v\n", p.Name)
	}
}

func (e Employee) GetEmployeeNumber() {
	if e != (Employee{}) {
		fmt.Printf("The employee's number is %d\n", e.EmpNum)
	}
}

func main() {
	Ross := FairHREmployee{Person{Name: "Ross", Age: 32}, Employee{Dept: "Dev", EmpNum: 3088869}}
	Ross.PrintName()
	Ross.GetEmployeeNumber()
}
