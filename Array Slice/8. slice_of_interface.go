package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	var interfaceSlice1 []interface{}
	// setting values of slice
	interfaceSlice1 = append(interfaceSlice1, 20, "Taylor", 10.10, true, nil)
	fmt.Println("interfaceSlice1", interfaceSlice1)
	fmt.Println()

	interfaceSlice2 := make([]interface{}, 3)
	interfaceSlice2[0] = 20
	interfaceSlice2[1] = "Taylor"
	interfaceSlice2[2] = true
	fmt.Println("interfaceSlice2", interfaceSlice2)
	fmt.Println()

	//in case of fmt.Println, the ... notation is used to print the elements of a slice 
	//as separate values.
	fmt.Println(interfaceSlice1...)

	// slice of struct
	var structSlice []Person
	structSlice = append(structSlice, Person{"James", 25})
	structSlice = append(structSlice, Person{"MoneyPenny", 30})

	fmt.Println(structSlice)
}
