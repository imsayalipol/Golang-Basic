package main

import "fmt"

type person struct {
	Name string
}

type kitchen interface {
	cook()
}

type bedroom interface {
	sleep()
}

type home interface {
	kitchen
	bedroom
	sing()
}

func (p *person) cook() {
	fmt.Println(p.Name, "cooks delicious food")
}

func (p *person) sleep() {
	fmt.Println(p.Name, "sleeps like a panda")
}

func check(h home) {
	h.cook()
}
func main() {

	sayali := person{"Taylor Swift"}
	sayali.cook()
	sayali.sleep()
	check(sayali)

	// sayali is not of type home even if its type kitchen and bedroom
	// as it can be seen that code dosen't through an error for not implementing sing()
}

// Explanation ::::
//  sayali is not of type home. The home interface embeds the kitchen and bedroom interfaces,
//  but it does not declare any additional methods. Therefore, a type that implements kitchen
//  and bedroom satisfies the home interface, but it is not required to have any additional
//  methods beyond those already declared by the embedded interfaces.

// In the code you provided, sayali is of type person, which implements both the kitchen and
//  bedroom interfaces via its cook() and sleep() methods, respectively.
// Therefore, a value of type person can be used as a value of type kitchen or bedroom,
// but it does not implement any additional methods required by the home interface.

// To create a value of type home, you would need to define a new type that embeds person
// and implements any additional methods required by the home interface, if any.
