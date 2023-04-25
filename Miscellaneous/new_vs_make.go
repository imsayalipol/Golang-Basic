package main

import "fmt"

func main() {

	// make() creates and initializes built-in types like slices, maps etc while
	//new() doesn't initialize the object but allocates a memory location
	//and returns a poinyter to that memory location

	var s *[]int
	s = new([]int)
	*s = make([]int, 5)

	(*s)[0] = 10
	(*s)[1] = 20
	(*s)[2] = 30
	(*s)[3] = 40
	(*s)[4] = 50

	fmt.Println("Slice is :", s)
	fmt.Println("Slice is :", *s)
}
