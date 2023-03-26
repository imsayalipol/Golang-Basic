package main

import "fmt"

func main() {

	type Person struct {
		Name    string
		Surname string
	}

	type Employee struct {
		Person
		Salary float32
	}

	employee := Employee{Person{"Arvind", "Ghosh"}, 152031.25}
	// fields can be accessed directly with var or var.struct.fieldName
	fmt.Println(employee.Name, employee.Person.Surname, employee.Salary)
}
