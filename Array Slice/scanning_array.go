package main

import (
	"fmt"
)

func main() {
var arr []int

	fmt.Println("Enter:")
	for{
		var num int
		_, err := fmt.Scan(&num)
        if err != nil {
            break
        }
        arr = append(arr, num)
	}
	 fmt.Println("Array elements:", arr)
}
