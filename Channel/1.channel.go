package main

import "fmt"

func main() {
	// empty channels

	// channel declaration
	ch1 := make(chan int)
	fmt.Println("Channel 1:", ch1)

	// channel definition with make()
	var ch2 chan int
	fmt.Println("Channel 2:", ch2)
}
