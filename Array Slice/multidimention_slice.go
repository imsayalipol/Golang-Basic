package main

import "fmt"

func main() {

	// 2D Slice
	var twoDSlice = make([][]int, 3) // 3 means a slice of 3 slices
	fmt.Println("twoDSlice :", twoDSlice)

	twoDSlice[0] = []int{1, 2, 3}
	twoDSlice[1] = []int{11, 22, 33}
	twoDSlice[2] = []int{111, 222, 333}

	fmt.Println("twoDSlice is :", twoDSlice)
	fmt.Println("Length is :", len(twoDSlice))

	fmt.Println()

	fmt.Println("Iterating the slice")
	for _, FirstSlice := range twoDSlice {
		for _, SecondSlice := range FirstSlice {
			fmt.Println(SecondSlice)
		}
	}
	fmt.Println()
	// 2D Slice of variable length

	var sample = make([][]int, 3)
	sample[0] = []int{1, 2, 3}
	sample[1] = []int{11, 22}
	sample[2] = []int{100, 200, 300, 400}

	fmt.Println("2D slice of variable length is :", sample)
	fmt.Println("Length is:", len(sample))

	fmt.Println()

	// 3D Slice
	var threeDSlice = make([][][]int, 2)

	for i := range threeDSlice {
		threeDSlice[i] = make([][]int, 3)
		for j := range sample[i] {
			threeDSlice[i][j] = make([]int, 3)
		}
	}
	fmt.Println(threeDSlice)

	fmt.Println("Iterating 3D slice:")

	for _, first := range threeDSlice {
		for _, second := range first {
			for _, third := range second {
				fmt.Println(third)
			}
		}
	}
}
