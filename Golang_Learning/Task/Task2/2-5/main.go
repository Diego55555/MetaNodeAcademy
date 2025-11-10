package main

import "fmt"

func main() {
	employee := Employee{Person: Person{Name: "Wang", Age: 21}, EmployeeID: 1001}
	employee.PrintInfo()
}

type Person struct {
	Name string
	Age  byte
}

type Employee struct {
	Person
	EmployeeID uint32
}

func (employee Employee) PrintInfo() {
	fmt.Println("Name: ", employee.Name, ", Age: ", employee.Age, ", EmployeeID", employee.EmployeeID)
}
