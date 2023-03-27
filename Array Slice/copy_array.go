package main

import "fmt"

// copy() accepts only slices as an argument so to copy array pass the slice of array
func main() {
	var mySlice = []int{1, 2, 3}
	fmt.Println("mySlice : ", mySlice)

	refCopySlice := mySlice
	fmt.Println("refCopySlice : ", refCopySlice)

	var newCopySlice = make([]int, 3)

	copy(newCopySlice, mySlice)
	fmt.Println("newCopySlice : ", newCopySlice)

	//modyifiying new copied slice using copy() won't make changes in original slice

	newCopySlice[1] = 100
	fmt.Println("Modified new copyied slice :", newCopySlice)
	fmt.Println("Original slice :", mySlice)
}
