package main

import (
	"fmt"
	"strconv"
)

func main() {

	var s string
	fmt.Println("Enter the value :")
	fmt.Scanln(&s)

	val, err := strconv.Atoi(s)

	if err != nil {
		fmt.Println("!!! NOT a number !!!")
	} else {
		fmt.Println("Its a number", val)
	}
}
