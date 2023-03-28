package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Salary float32
}

func (p Person) ShowName() Person {
	fmt.Println("Name is :", p.Name)
	return p
}

func (p Person) ShowAge() Person {
	fmt.Println("Age is :", p.Age)
	return p
}

func (p Person) ShowSalary() {
	fmt.Println("Salary is :", p.Salary)
}

func main() {

	person1 := Person{
		Name:   "Dua Lipa",
		Age:    28,
		Salary: 4581258.15,
	}

	person1.ShowName().ShowAge().ShowSalary()

	// In method chaining methods should return the receiver, except for the last method
	// where returning the receiver is not neccessary
}
