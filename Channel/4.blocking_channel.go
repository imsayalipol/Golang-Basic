package main

import "fmt"

func main() {

	ch := make(chan int, 1)

	fmt.Println("Sending to channel")
	ch <- 1
	//ch <- 1 // this will block channel as its full and will create a deadlock

	fmt.Println("Receiving from channel")
	val := <-ch
	// val = <-ch // this will create block as channel is empty
	fmt.Println("Received value:", val)
}
