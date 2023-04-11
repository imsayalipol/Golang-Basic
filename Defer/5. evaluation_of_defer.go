package main

import "fmt"

func main() {

	sample := "abc"
	defer fmt.Println("Inside defer:", sample) //defer arguments are evaluated at the time defer statement is evaluated

	sample = "xyz"
	// defer fmt.Println("Inside defer 2nd time:", sample)  //this will print xyz

}
