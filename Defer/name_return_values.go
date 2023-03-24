// package main

// import "fmt"

//	func main() {
//		s := test()
//		fmt.Println("Value is :", s)
//	}
//
//	func test() (int size) {
//		// var size int
//		defer func() { size = 20 }()
//		size = 30
//		return
//	}
package main

import "fmt"

func main() {
	s := test()
	fmt.Println(s)
}
func test() int {
	var size int
	defer func() { size = 20 }()
	size = 30
	return size
}
