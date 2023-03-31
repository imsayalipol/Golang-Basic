package main

import "fmt"

type animal interface {
	walks()
}

type pet struct {
	age int
}

func (p pet) walks() {
	fmt.Println("This animal walks")
}

type dog struct {
	a    animal // named interface
	name string
}

type cat struct {
	animal // embedded anonymously
	name   string
}

func main() {

	p := pet{age: 5}

	// if interface is named we call it with that name, here 'a'
	d := dog{a: p, name: "Marcos"}
	d.a.walks()
	// d.walks()  this will throw error

	// for anonymous interafce we call it with interface name
	c := cat{animal: p, name: "Snowy"}
	c.walks()

}

// important :
// If we don’t initialise the embedded interface animal, then it will be intialised
//with the zero value of the interface which is nil.
//Calling  walks() method  on such an instance of dog or cat struct will
//create a panic.
