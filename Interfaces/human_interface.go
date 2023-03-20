// This example is similar to interface_simplified_example.go
// only fdifference is here we are passing a struct variable to a function
// which takes interface variable as a parameter and invoke expected method of struct

package main

import "fmt"

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	kill bool
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, "says Hello James !!")
}

func (sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, "says I am here to kill")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	P := person{"Miss", "Moneypenny"}
	P.speak()

	SA := secretAgent{person{"James", "Bonad"}, true}
	SA.speak()
	SA.person.speak() //VIMP*** THIS wil print the person's msg
	// which simply means secretAgent can access method of person

	// to demo the interface function access :
	// speak() is implemented by both structs and we can create a struct variable and
	// pass it as a interface variable which in return will invoke speak() method of
	// the expected struct
	demoFunc := person{"Taylor", "Swift"}
	saySomething(demoFunc)
}
