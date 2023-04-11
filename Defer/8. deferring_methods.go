package main

import (
	"errors"
	"fmt"
)

type details struct {
	name string
}

func (d *details) setName(name string) error {
	defer d.setDefaultName()

	if len(name) < 3 {
		return errors.New("Name len can't be less than 3")
	}

	d.name = name
	fmt.Println("Returning from setName()")
	return nil
}

func (d *details) setDefaultName() {
	fmt.Println("Inside setDefaultName()")

	if d.name == "" {
		fmt.Println("Name can not be empty. Setting default name")
		d.name = "Default Name"
	}

}

func main() {

	person1 := &details{}
	person1.setName("Jeffery")
	fmt.Println("Person 1 name set to :", person1.name)

	fmt.Println()

	person2 := &details{}
	person2.setName("Po") //no name will be set to this variable as len(name)<3
	//hence name is empty and hence default name is set
	fmt.Println("Person 2 name set to :", person2.name)

}
