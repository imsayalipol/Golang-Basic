package main

import "fmt"

type animals interface {
	roaring()
}

type lion struct {
	age int
}

func (l lion) roaring() {
	fmt.Println("Lion roars")
}

func main() {
	var a animals
	a = lion{age: 10}
	fmt.Printf("Type is %T\n", a)
	fmt.Printf("Value is %v\n", a)
}
