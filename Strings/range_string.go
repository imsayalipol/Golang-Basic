package main

import "fmt"

func main() {

	var name = "Taylor Swift"

	for i, v := range name {
		fmt.Printf("Index :%d Value :%s\n", i, string(v))
		// fmt.Println("Index :", i, " Value", v)

	}

	var animal string
	fmt.Println("animal :", animal)
	// iterating empty string won't through error
	for m, n := range animal {
		fmt.Printf("Index :%d Value :%s\n", m, string(n))

	}
}
