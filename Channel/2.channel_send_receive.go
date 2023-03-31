package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)

	fmt.Println("Sending value to channel")
	go send(ch) //goroutine

	fmt.Println("Receiving value from channel")
	go receive(ch) //goroutine

	time.Sleep(1 * time.Second)
}

// sending value TO channel
func send(ch chan int) {
	ch <- 1
}

func receive(ch chan int) {
	val := <-ch
	fmt.Println("Received Value is :", val)
}
