package main

import "fmt"

func main() {
	defer test()
	fmt.Println("In main func")
}

func test() {
	fmt.Println("In TEST func")
}
