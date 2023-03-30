package main

import "fmt"

type animal interface {
	walk()
	breath()
}

// human embed animal interface so any type implementing human interface
// has to implement animal interface as well
type human interface {
	speak()
	animal
}

func (w woman) walk() {
	fmt.Println("Can walk")
}

func (w woman) breath() {
	fmt.Println("Can breath")
}

func (w woman) speak() {
	fmt.Println("Can speak")
}

type woman struct {
	name string
}

func main() {

	w := woman{name: "Taylor"}

	var h human
	h = w
	h.breath()
	h.speak()
	h.walk()

}
