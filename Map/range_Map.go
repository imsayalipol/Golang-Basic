package main

import "fmt"

func main() {

	m := map[int]string{
		1: "Golang",
		2: "Gin Framework",
		3: "DSA"}

	for k, v := range m {
		fmt.Println("Key: ", k, "Value: ", v)
	}
}
