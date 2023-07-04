package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {

	arrOfArray := [2][2]int{
		{548, 589},
		{589, 255},
	}
	fmt.Println("Array of Array :", arrOfArray)

	arrOfSlice := [2][]int{
		{1, 2, 3},
		{10, 30},
	}

	fmt.Println("Array of Slice :", arrOfSlice[1])

	arrOfStruct := [3]Person{
		{"Taylro Swift", 33},
		{"Ariana", 30},
		{"Dua Lipa", 28},
	}

	fmt.Println("Array of Struct :", arrOfStruct)

	arrOfInterface := [5]interface{}{
		20,
		100.5,
		"The Boat",
		true,
		[]int{1, 2, 3, 4},
	}

	fmt.Println("Array of Interfaces :", arrOfInterface)
}
