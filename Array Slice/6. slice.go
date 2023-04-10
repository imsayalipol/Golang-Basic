package main

import (
	"fmt"
)

func main() {
	// ways of declaring slice:

	var firstSlice []int
	fmt.Println("firstSlice==", firstSlice)

	var secondSlice = make([]float32, 5, 8) //float size with length 5 and capacity 8
	secondSlice = []float32{10.11, 20.22, 30.33, 10.10, 20.20}
	fmt.Println("secondSlice ==", secondSlice)

	var thirdSlice = make([]int, 5)
	fmt.Println("thirdSlice ==", thirdSlice)

	var copySlice = thirdSlice
	fmt.Println("Copy slice : === :", copySlice)

	thirdSlice[2] = 5
	fmt.Println("thirdSlice new value ===>", thirdSlice)
}
