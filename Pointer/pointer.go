package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 10
	b := new(int)

	fmt.Println("Type of a :", reflect.TypeOf(a))
	fmt.Println("Type of b :", reflect.TypeOf(b))
	fmt.Println("a=", a, "b=", b, "*b=", *b)

	b = &a
	fmt.Println("*b = ", *b)

	*b = 15
	fmt.Println("a =", a)
}
