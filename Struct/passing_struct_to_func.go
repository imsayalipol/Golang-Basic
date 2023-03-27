package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func test(data Person) {
	data.Name = "Sam"
	fmt.Println("inside test() :", data)
}

func main() {

	p := Person{
		Name: "Max",
		Age:  30,
	}
	ptr := &p
	fmt.Println("*ptr :", *ptr)
	test(*ptr)
	fmt.Println("inside main() :", *ptr)
}
