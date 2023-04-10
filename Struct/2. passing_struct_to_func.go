package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func test(data Person) {
	data.Name = "Sam"
}

func test2(data *Person) {
	data.Name = "Shivani"
}

func main() {

	// In this case value of field Name won't change in original struct
	p := Person{
		Name: "Max",
		Age:  30,
	}
	ptr := &p
	fmt.Println("*ptr :", *ptr)
	test(*ptr) // even after sending pointer, the value of field will remain 30
	fmt.Println("inside main() p:", p)

	// In this case value of field Name will change in original struct
	s := &Person{
		Name: "Sayali",
		Age:  30}
	fmt.Println("inside main() s:", s)

	test2(s)
	fmt.Println("inside main() s after test 2:", s)

}
