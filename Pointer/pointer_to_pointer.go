package main

import "fmt"

func main() {
	a := 30
	b := &a
	c := &b

	// printing the values stored
	fmt.Println("a =", a)
	fmt.Println("*&a =", *&a)
	fmt.Println("*b =", *b)
	fmt.Println("*c =", **c)

	//printing addresses
	fmt.Println("add of a: &a :", &a)
	fmt.Println("add of a: b :", b)
	fmt.Println("add of a: *c", *c)
	fmt.Println("add of b: &b", &b)
	fmt.Println("add of b: c", c)
	fmt.Println("add of c: &c", &c)

	// pointer arithmetic
	*b++
	fmt.Println("*b++ :", *b)
	*b--
	fmt.Print("*b-- :", *b)
}
