package main

import "fmt"

func main() {

	ch := make(chan int, 1) //buffured channel with capacity 1

	fmt.Println("Sending value to channel")
	ch <- 1
	// ch <- 1    // uncommenting this line will cause error :- fatal error: all goroutines are asleep - deadlock!

	fmt.Println("Receiving channel")
	val := <-ch
	fmt.Println("Received value is :", val)

}
