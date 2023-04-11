package main

import "fmt"

func main() {
	defer test() //a defer statement calling the custom function named test()
	fmt.Println("In main func")
}

func test() {
	fmt.Println("In TEST func")
}
