package main

import "fmt"

func main() {

	// declaring a struct
	type Person struct {
		Name   string
		Age    int
		Salary float32
	}

	// creating variable of a struct in different ways
	// way 1:
	var person1 = Person{Name: "Taylor Swift", Age: 34, Salary: 802458.365}
	fmt.Println("Person 1 info :", person1)

	// way 2:
	person2 := Person{"Sabrina Carpenter", 24, 12015.25}
	fmt.Println("Person 2 info :", person2)

	// way 3:
	var person3 Person
	person3.Name = "Ariana Grande"
	person3.Age = 30
	person3.Salary = 210250.20
	fmt.Println("Person 3 info :", person3)

	// not declaring all fields of a struct
	person4 := Person{"Sayali", 31, 40215.25}
	var person5 Person
	person5.Age = 50

	fmt.Println("person 4 info :", person4)
	fmt.Println("Person 5 info :", person5)
	fmt.Println("Another way of printing")
	fmt.Printf("Person 5: %+v\n", person5)

}
