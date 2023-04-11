package main

import "fmt"

func main() {
	s := test()
	fmt.Println(s)

	p := add(10, 20)
	fmt.Println(p)

}

// unnamed return will retun 30 here
func test() int {
	var size int
	defer func() { size = 20 }()
	size = 30
	return size
}

// named return will return value modified by defer()
func add(x, y int) (result int) {
	defer func() { result = 100 }()
	result = x + y
	return
}
