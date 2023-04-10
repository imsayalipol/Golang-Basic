package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	var interfaceSlice []interface{}

	interfaceSlice = append(interfaceSlice, 20, "Taylor", 10.10, true, nil)
	fmt.Println(interfaceSlice)
	fmt.Println()
	//in case of fmt.Println, the ... notation is used to print the elements of a slice as separate values.
	fmt.Println(interfaceSlice...)

	// slice of struct
	var structSlice []Person
	structSlice = append(structSlice, Person{"James", 25})
	structSlice = append(structSlice, Person{"MoneyPenny", 30})

	fmt.Println(structSlice)
}
