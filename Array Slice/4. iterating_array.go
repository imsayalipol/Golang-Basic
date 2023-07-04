package main

import "fmt"

func main() {

	var arr = [5]int{10, 20}
	// 1D array using for loop
	fmt.Printf("Iterating with for loop :")
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	fmt.Println("Iterating with for-range :")
	for index, value := range arr {
		fmt.Println("index :", index, "value :", value)
	}

	fmt.Println()

	// 2D array
	fmt.Println("Iterating 2D array:")
	var TwoD_array = [3][4]int{{1, 2, 3, 4},
		{55, 66, 77, 88},
		{500, 300, 400, 800}}

	for _, row := range TwoD_array {
		// fmt.Println(row)   //this prints only rows
		for i, v := range row {
			fmt.Println("index :", i, "value :", v)
		}
	}
	fmt.Println("Using for loop on TwoD_array :")
	for r := 0; r < 3; r++ {
		for c := 0; c < 4; c++ {
			fmt.Println(TwoD_array[r][c])
		}
	}
}
