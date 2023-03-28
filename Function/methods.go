package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Salary float32
}

// value receiver method
func (p Person) NameChange(name string) {
	p.Name = name
	fmt.Println("Inside NameChnage() name is:", p.Name)
}

// pointer receiver method
func (p *Person) AgeChange(age int) {
	p.Age = age
	fmt.Println("Inside AgeChange() age is:", p.Age)
}

func main() {

	person1 := Person{
		Name:   "James Bond",
		Age:    55,
		Salary: 5541258.52,
	}
	person1.NameChange("Taylor Swift")
	fmt.Println("After calling NameChange func :", person1.Name) // Name = James Bond as its a value receiver method

	person1.AgeChange(30)
	fmt.Println("After calling AgeChange func :", person1.Age) // Age = 30 as its a pointer receiver method

}
