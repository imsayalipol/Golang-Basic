package main

import "fmt"

func main() {
// empty channels

	ch1 := make(chan int)
	fmt.Println("Channel 1:", ch1)

	var ch2 chan int
	fmt.Println("Channel 2:", ch2)
}
